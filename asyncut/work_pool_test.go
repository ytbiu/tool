package asyncut

import (
	"testing"
	"time"
)

func Test_dispatcher_Add(t *testing.T) {
	job := func() {
		t.Log("job do")
		time.Sleep(time.Second * 100)
	}

	jobFail := func() {
		t.Log("job fail do")
		time.Sleep(time.Second * 100)
	}

	limitPass := int32(3)
	limitFail := int32(2)
	toAdd := 3
	jobs := make([]func(), 0, toAdd)
	for i := 0; i < toAdd; i++ {
		jobs = append(jobs, job)
	}

	jobsFail := make([]func(), 0, toAdd)
	for i := 0; i < toAdd; i++ {
		jobsFail = append(jobsFail, jobFail)
	}

	type fields struct {
		Cap        int32
		RunningJob int32
		JobC       chan func()
		ExitC      chan struct{}
	}
	type args struct {
		jobs []func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				Cap:        limitPass,
				RunningJob: 0,
				JobC:       make(chan func(), limitPass),
				ExitC:      make(chan struct{}),
			},
			args: args{
				jobs: jobs,
			},
		},
		{
			fields: fields{
				Cap:        limitFail,
				RunningJob: 0,
				JobC:       make(chan func(), limitFail),
				ExitC:      make(chan struct{}),
			},
			args: args{
				jobs: jobsFail,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(time.Millisecond * 10)
			NewDispatcher(tt.fields.Cap).Go(tt.args.jobs...)
		})
		time.Sleep(3 * time.Second)
	}
}
