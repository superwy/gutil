package gutil

import (
	"errors"
	"sync"
)

func RunAsyncWorker(fnWorkers ...func() error) (err error) {
	if len(fnWorkers) <= 0 {
		return
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(fnWorkers))
	for _, fnItemWork := range fnWorkers {
		go func(itemWork func() error) {
			defer func() {
				if r := recover(); r != nil {
					switch x := r.(type) {
					case string:
						err = errors.New(x)
					case error:
						err = x
					default:
						err = errors.New("Unknown panic")
					}
				}
			}()
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
		}(fnItemWork)
	}
	wg.Wait()
	return
}
