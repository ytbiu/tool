package gout

import "sync"

func JobPipeline(jobs ...func() interface{}) chan interface{} {

	resC := make(chan interface{}, 100)
	wg := sync.WaitGroup{}
	for _, job := range jobs {

		job := job
		wg.Add(1)
		go func() {
			defer wg.Done()
			if res := job(); res != nil {
				resC <- res
			}

		}()
	}

	go func() {
		wg.Wait()
		close(resC)
	}()

	return resC
}

