package asyncut

import (
	"testing"
	"time"
)




func Test_dispatcher_Go(t *testing.T) {

	job := func() {
		//log.Println("job doing")
	}

	dispatcher := NewDispatcher(5)
	for i := 0; i < 10; i++ {
		dispatcher.Go(job)
	}

	time.Sleep(time.Second*2)
}

func Benchmark_dispatcher_Go(b *testing.B)  {
	b.ResetTimer()

	job := func() {
		_ = 1+1+1+1+1
	}

	dispatcher := NewDispatcher(50000)

	b.StartTimer()
	for i := 0; i<1000000; i++ {
		dispatcher.Go(job)
	}
}

func Benchmark_common_Go(b *testing.B)  {
	b.ResetTimer()

	job := func() {
		_ = 1+1+1+1+1
		time.Sleep(time.Second)
	}

	b.StartTimer()
	for i := 0; i<1000000; i++ {
		go job()
	}
}