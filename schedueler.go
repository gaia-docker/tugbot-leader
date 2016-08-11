package main

import (
	"time"
)

func NewScheduler(task func()) *Scheduler {
	return &Scheduler{task: task, stop: false, timer: time.NewTimer(time.Minute)}
}

type Scheduler struct {
	task  func()
	stop  bool
	timer *time.Timer
}

func (s Scheduler) Run() {
	for !s.stop {
		s.task()
		s.timer = time.NewTimer(time.Minute)
		<-s.timer.C
	}
}

func (s Scheduler) Stop() bool {
	s.stop = true

	return s.timer.Stop()
}
