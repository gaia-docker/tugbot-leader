package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/opencontainers/runc/Godeps/_workspace/src/github.com/urfave/cli"
)

var (
	scheduler *Scheduler
	wg        sync.WaitGroup
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "tugbot-leader"
	app.Usage = "Continuous Testing Framework for Docker Swarm"
	app.Before = before
	app.Action = start
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug mode with verbose logging",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func before(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	controller := NewServiceUpdater(dockerClient)
	scheduler = NewScheduler(func() error { return controller.Update() })

	return nil
}

func start() {
	log.Info("Starting Tugbot Leader...")
	go func() {
		wg.Add(1)
		scheduler.Run()
		wg.Done()
	}()
	waitForInterrupt()
}

func waitForInterrupt() {
	// Graceful shut-down on SIGINT/SIGTERM/SIGQUIT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
	scheduler.Stop()
	wg.Wait()
	os.Exit(1)
}
