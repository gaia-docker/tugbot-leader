package util

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchedulerRun(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	times := 0
	scheduler := NewScheduler(func() error {
		defer wg.Done()
		times++

		return nil
	}, time.Minute)
	go func() { scheduler.Run() }()
	wg.Wait()
	assert.Equal(t, 1, times)
}
