package service

import "sync"

var wg sync.WaitGroup

// Run ...
func Run() {
	wg.Add(1)
	go func() {
		for {
		}
	}()
	wg.Wait()
}
