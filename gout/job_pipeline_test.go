package gout

import (
	"testing"
)

func TestJobPipeline(t *testing.T) {

	job1 := func() interface{}{
		return 1
	}

	job2 := func() interface{}{
		return 2
	}

	type args struct {
		jobs []func() interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				jobs: []func()interface{}{job1,job2},
			},
		},
	}

	var i int
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC := JobPipeline(tt.args.jobs...)
			for res := range gotC{

				if res.(int) == 1 || res.(int) == 2{
					i++
					continue
				}else{
					t.Fatal(i,res)
				}
			}
		})
	}
}
