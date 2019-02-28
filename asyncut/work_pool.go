package asyncut

import (
	"log"
	"runtime/debug"
	"sync/atomic"
	"time"
)

const waitSeconds = 3 * time.Second

type Dispatcher interface {
	Add(...func())
}

type dispatcher struct {
	Cap        int32
	RunningJob int32
	JobC       chan func()
	ExitC      chan struct{}
}

func NewDispatcher(maxLimit int32) Dispatcher {
	d := &dispatcher{
		Cap:  maxLimit,
		JobC: make(chan func(), maxLimit),
	}

	go d.dispatch()

	return d
}

func (d *dispatcher) dispatch() {
	for {
		select {
		case <-d.ExitC:
			break
		default:
		}

		for job := range d.JobC {
			if d.GetRunningJob() >= d.Cap {
				<-time.After(waitSeconds)
				if d.GetRunningJob() >= d.Cap {
					d.logWarning()
					return
				}
			}

			d.incRunningJob()
			go func() {
				defer catch()
				defer d.decRunningJob()
				job()
			}()
		}
	}

}

func (d *dispatcher) Add(jobs ...func()) {
	for _, job := range jobs {
		if job == nil {
			continue
		}

		select {
		case <-d.ExitC:
			return
		case <-time.After(waitSeconds):
			d.logWarning()
			return
		case d.JobC <- job:
		}
	}
}

func (d *dispatcher) GetRunningJob() int32 {
	return atomic.LoadInt32(&d.RunningJob)
}

func (d *dispatcher) incRunningJob() {
	atomic.AddInt32(&d.RunningJob, 1)
}

func (d *dispatcher) decRunningJob() {
	atomic.AddInt32(&d.RunningJob, -1)
}

func (d *dispatcher) logWarning() {
	log.Printf("jobs is full -- runningJobs: %d; capJobs: %d", d.RunningJob, d.Cap)
}

func (d *dispatcher) Exit() {
	d.ExitC <- struct{}{}
}

func catch() {
	if err := recover(); err != nil {
		debug.PrintStack()
		log.Println(" panic happen -- error : ", err)
	}
}
