package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/gaia-docker/tugbot-leader/swarm"
	"github.com/gaia-docker/tugbot-leader/util"
)

var (
	scheduler *util.Scheduler
	wg        sync.WaitGroup
)

func main() {
	initialize()
	start()
	waitForInterrupt()
}

func initialize() {
	setLogLevel()

	// Uncomment when debug using docker machine
	// util.SetEnv()

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Failed to create docker client (%+v)", err)
	}

	if err = swarm.IsValidNode(dockerClient); err != nil {
		log.Fatal(err)
	}

	updater := swarm.NewServiceUpdater(dockerClient)
	scheduler = util.NewScheduler(func() error { return updater.Run() }, getInterval())
}

func setLogLevel() {
	if os.Getenv(util.TugbotLogLevel) == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
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
