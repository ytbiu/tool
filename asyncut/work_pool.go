package asyncut

import (
	"log"
	"runtime/debug"
	"sync/atomic"
	"time"
)

const waitSeconds = 3 * time.Second

type Dispatcher interface {
	Go(...func())
	//getRunningJob() int32
}

type dispatcher struct {
	poolSize     int32
	runningJob   int32
	jobC         chan func()
	exitC        chan struct{}
	workCancelCs []chan struct{}
	resizePeriod time.Duration
	timer        *time.Timer
}

func NewDispatcher(maxLimit int32) Dispatcher {
	d := &dispatcher{
		poolSize:     maxLimit,
		jobC:         make(chan func(), maxLimit),
		resizePeriod: time.Second * 3,
		timer:        time.NewTimer(waitSeconds),
	}

	d.workCancelCs = make([]chan struct{}, maxLimit)
	for i := 0; i < int(maxLimit); i++ {
		d.workCancelCs[i] = make(chan struct{})
	}

	go d.tryResize()
	go d.dispatch()

	return d
}

func (d *dispatcher) dispatch() {

	poolSize := int(d.poolSize)
	for i := 0; i < poolSize; i++ {
		i := i
		go func() {
			defer catch()
			d.listenJob(d.workCancelCs[i])
		}()
	}
}

func (d *dispatcher) Go(jobs ...func()) {
	d.timer.Reset(waitSeconds)
	for _, job := range jobs {
		if job == nil {
			continue
		}

		select {
		case <-d.exitC:
			return
		case <-d.timer.C:
			d.logWarning()
			return
		case d.jobC <- job:
		}
	}
}

func (d *dispatcher) getRunningJob() int32 {
	return atomic.LoadInt32(&d.runningJob)
}

func (d *dispatcher) incRunningJob() {
	atomic.AddInt32(&d.runningJob, 1)
}

func (d *dispatcher) decRunningJob() {
	atomic.AddInt32(&d.runningJob, -1)
}

func (d *dispatcher) incPoolSize() {
	atomic.AddInt32(&d.poolSize, 1)
}

func (d *dispatcher) decPoolSize() {
	atomic.AddInt32(&d.poolSize, 1)
}

func (d *dispatcher) tryResize() {
	for {
		time.Sleep(d.resizePeriod)

		select {
		case <-d.exitC:
			return
		default:
		}

		if d.getRunningJob() == d.poolSize {
			go func() {
				defer catch()
				cancelC := make(chan struct{})
				d.workCancelCs = append(d.workCancelCs, cancelC)
				d.listenJob(cancelC)
			}()
			d.incPoolSize()
		}

		minLimit := (1 / 2) * d.poolSize
		if d.getRunningJob() < minLimit {
			if minLimit <= 5 {
				return
			}

			for i := 0; i < int(minLimit); i++ {
				d.workCancelCs[i] <- struct{}{}
			}
		}
	}

}

func (d *dispatcher) listenJob(workerCancelC chan struct{}) {
	for job := range d.jobC {

		select {
		case <-d.exitC:
			return
		case <-workerCancelC:
			return
		default:
		}

		d.incRunningJob()
		job()
		d.decRunningJob()
	}
}

func (d *dispatcher) logWarning() {
	log.Printf("jobs is full -- runningJobs: %d; capJobs: %d", d.runningJob, d.poolSize)
}

func (d *dispatcher) Exit() {
	d.exitC <- struct{}{}
}

func catch() {
	if err := recover(); err != nil {
		debug.PrintStack()
		log.Println(" panic happen -- error : ", err)
	}
}
