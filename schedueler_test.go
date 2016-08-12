package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSchedulerRun(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	times := 0
	scheduler := NewScheduler(func() error {
		defer wg.Done()
		times++

		return nil
	})
	go func() { scheduler.Run() }()
	wg.Wait()
	scheduler.Stop()
	assert.Equal(t, 1, times)
}
