package util

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type Scheduler struct {
	task     Task
	stop     bool
	timer    *time.Timer
	interval time.Duration
}

type Task func() error

func NewScheduler(t Task, interval time.Duration) *Scheduler {
	return &Scheduler{task: t, stop: false, timer: time.NewTimer(-1), interval: interval}
}

func (s Scheduler) Run() {
	for !s.stop {
		if err := s.task(); err != nil {
			log.Error(err)
		}
		s.timer = time.NewTimer(s.interval)
		<-s.timer.C
	}
}

func (s Scheduler) Stop() bool {
	s.stop = true

	return s.timer.Stop()
}
