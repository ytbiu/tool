package asyncut

import (
	"testing"
	"time"
	"github.com/panjf2000/ants"
)

func job() {
	f := make(map[string]int)
	f["l"]=1
	time.Sleep(time.Second)
}


func Test_dispatcher_Go(t *testing.T) {

	job := func() {

	}

	dispatcher := NewDispatcher(5)
	for i := 0; i < 100; i++ {
		dispatcher.Go(job)
	}

	time.Sleep(time.Second * 2)
}

func Benchmark_ants(b *testing.B)  {
	b.ResetTimer()

	p,_ := ants.NewPool(2000)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i:=0;i<10000;i++{
			p.Submit(job)
		}
	}
}

func Benchmark_dispatcher_Go(b *testing.B) {
	b.ResetTimer()

	dispatcher := NewDispatcher(2000)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i:=0;i<10000;i++{
			dispatcher.Go(job)
		}
	}
}

func Benchmark_common_Go(b *testing.B) {
	b.ResetTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i:=0;i<10000;i++{
			go job()
		}
	}
}
