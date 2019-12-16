package asyncut

import (
	"log"
	"runtime/debug"
	"sync/atomic"
	"time"
	"math/rand"
)

const (
	defaultWaitSeconds = 1 * time.Second
	defaultPoolSize = 10
	)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

type Dispatcher interface {
	Go(...func())
	Close()
	//getRunningJob() int32
}

type dispatcher struct {
	poolSize     int32
	runningJob   int32
	jobCGroup    []chan func()
	exitC        chan struct{}
	resizePeriod time.Duration
	timeoutTimer *time.Timer
}

func NewDispatcher(poolSize ...int) Dispatcher {
	size := defaultPoolSize
	if len(poolSize) > 0 && poolSize[0] >0{
		size = poolSize[0]
	}

	d := &dispatcher{
		poolSize:     int32(size),
		jobCGroup:    make([]chan func(),0, size),
		resizePeriod: time.Second * 3,
		timeoutTimer: time.NewTimer(defaultWaitSeconds),
	}

	for i:=0;i<size;i++{
		d.jobCGroup = append(d.jobCGroup, make(chan func(),10))
	}

	go d.resizeLoop()
	go d.dispatch()

	return d
}

func (d *dispatcher) dispatch() {

	poolSize := int(d.poolSize)
	for i := 0; i < poolSize; i++ {
		i := i
		go func() {
			defer catch()
			d.listenJob(d.jobCGroup[i])
		}()
	}
}

func (d *dispatcher) Go(jobs ...func()) {
	d.timeoutTimer.Reset(defaultWaitSeconds)
	defer d.timeoutTimer.Stop()

	for _, job := range jobs {
		if job == nil {
			continue
		}

		i := rand.Intn(len(d.jobCGroup))

		select {
		case <-d.exitC:
			return
		case <-d.timeoutTimer.C:
			d.logWarning()
			d.resize(job)
		case d.jobCGroup[i] <- job:
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

func (d *dispatcher) resizeLoop() {
	for {
		time.Sleep(d.resizePeriod)

		select {
		case <-d.exitC:
			return
		default:
		}

		d.resize()
	}
}

func (d *dispatcher) resize(job ...func())  {
	if d.getRunningJob() == d.poolSize {
		go func() {
			defer catch()
			jobC := make(chan func())
			d.jobCGroup = append(d.jobCGroup, jobC)
			d.listenJob(jobC)
			if len(job) > 0{
				jobC <- job[0]
			}
		}()
		d.incPoolSize()
	}
}

func (d *dispatcher) listenJob(jobC chan func()) {
	for job := range jobC {
		select {
		case <-d.exitC:
			return
		default:
		}

		d.incRunningJob()
		job()
		d.decRunningJob()
	}
}

func (d *dispatcher) Close() {
	d.exitC <- struct{}{}
}

func (d *dispatcher) logWarning() {
	log.Printf("job pool is full -- runningJobs: %d; poolSiez: %d", d.runningJob, d.poolSize)
}

func catch() {
	if err := recover(); err != nil {
		debug.PrintStack()
		log.Println(" panic happen -- error : ", err)
	}
}
