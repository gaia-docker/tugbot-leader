package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

func NewScheduler(task func() error) *Scheduler {
	return &Scheduler{task: task, stop: false, timer: time.NewTimer(time.Minute)}
}

type Scheduler struct {
	task  func() error
	stop  bool
	timer *time.Timer
}

func (s Scheduler) Run() {
	for !s.stop {
		if err := s.task(); err != nil {
			log.Error(err)
		}
		s.timer = time.NewTimer(time.Minute)
		<-s.timer.C
	}
}

func (s Scheduler) Stop() bool {
	s.stop = true

	return s.timer.Stop()
}
