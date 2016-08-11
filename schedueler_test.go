package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

func TestRun(t *testing.T) {
	sem := sync.WaitGroup{}
	sem.Add(1)
	times := 0
	scheduler := NewScheduler(func() {
		defer sem.Done()
		times++
	})
	go func() { scheduler.Run() }()
	log.Print(scheduler)
	sem.Wait()
	log.Print(scheduler)
	scheduler.Stop()
	log.Print(scheduler)
	assert.Equal(t, 1, times)
}
