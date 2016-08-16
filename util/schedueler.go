package util

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type Scheduler struct {
	task     Task
	interval time.Duration
	quit     chan int
}

type Task func() error

func NewScheduler(t Task, interval time.Duration) *Scheduler {
	return &Scheduler{task: t, interval: interval, quit: make(chan int)}
}

func (s Scheduler) Run() {
	for {
		if err := s.task(); err != nil {
			log.Error(err)
		}
		select {
		case <-s.quit:
			return
		case <-time.After(time.Second * 7):
		}
	}
}

func (s Scheduler) Stop() {
	s.quit <- 1
}
