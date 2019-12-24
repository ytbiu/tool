package pipline

import "sync"

// 异步执行jobs 将结果汇总至consume 中的入参chan中待消费
func JobPipeline(consume func(chan interface{}), jobs ...func() interface{}) {
	resC := make(chan interface{}, 100)
	go consume(resC)

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
}
