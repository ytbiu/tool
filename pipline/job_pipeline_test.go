package pipline

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJobPipeline(t *testing.T) {

	job1 := func() interface{} {
		return 1
	}

	job2 := func() interface{} {
		return 2
	}

	var total int
	JobPipeline(func(c chan interface{}) {
		for res := range c {
			total += res.(int)
		}
	},
		job1,
		job2,
	)

	time.Sleep(time.Second)
	assert.Equal(t, 3, total)
}
