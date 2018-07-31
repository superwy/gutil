package gutil

import "sync"

func RunAsyncWorker(fnWorkers ...func() error) (err error) {
	if len(fnWorkers) <= 0 {
		return
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(fnWorkers))
	for _, itemWork := range fnWorkers {
		go func() {
			defer wg.Done()
			if err != nil {
				return
			}
			e := itemWork()
			if e != nil {
				func() {
					mu.Lock()
					defer mu.Unlock()
					err = e
					return
				}()
			}
		}()
	}
	wg.Wait()
	return
}
