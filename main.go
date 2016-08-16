package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/gaia-docker/tugbot-leader/swarm"
	"github.com/gaia-docker/tugbot-leader/util"
	"time"
)

var (
	scheduler *util.Scheduler
	wg        sync.WaitGroup
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	start()
	waitForInterrupt()
}

func init() {
	setLogLevel()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	updater := swarm.NewServiceUpdater(dockerClient)
	scheduler = util.NewScheduler(func() error { return updater.Run() }, getInterval())
}

func setLogLevel() {
	if os.Getenv(util.TugbotLogLevel) == "debug" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
}

func getInterval() time.Duration {
	ret := time.Minute
	interval := os.Getenv(util.TugbotInterval)
	if interval != "" {
		duration, err := time.ParseDuration(interval)
		if err != nil {
			log.Errorf("Failed to parse %s (%v)", util.TugbotInterval, err)
		} else {
			ret = duration
		}
	}

	return ret
}

func start() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Starting Tugbot Leader...")
		scheduler.Run()
	}()
}

func waitForInterrupt() {
	// Graceful shut-down on SIGINT/SIGTERM/SIGQUIT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
	log.Info("Stopping Tugbot Leader...")
	scheduler.Stop()
	wg.Wait()
	os.Exit(1)
}
