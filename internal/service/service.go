package service

import (
	"sync"

	"github.com/jordan-lumley/a1pos/internal/logger"
)

var wg sync.WaitGroup

// Run ...
func Run() {
	wg.Add(1)
	go func() {
		for {
		}
	}()

	logger.Instance().Info("Service Running....")
	wg.Wait()
}
