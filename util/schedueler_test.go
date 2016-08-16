package util

import (
	"sync"
	"testing"
	"time"
)

func TestSchedulerRun(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	scheduler := NewScheduler(func() error {
		defer wg.Done()
		return nil
	}, time.Minute)
	go func() {
		defer wg.Done()
		scheduler.Run()
	}()
	scheduler.Stop()
	// Wait validates that scheduler runs the provided method and finish running go routine,
	// meaning stop scheduler is working (otherwise, test will run forever)
	wg.Wait()
}
