package main

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
	"log"
)

func OnServiceUpdate(client client.ServiceAPIClient, updatedServices []string) error {
	services, err := getTestServices(client)
	if err != nil {
		return err
	}

	for _, currService := range services {
		log.Print(currService)
	}

	return nil
}

func getTestServices(client client.ServiceAPIClient) ([]swarm.Service, error) {
	filters := filters.NewArgs()
	filters.Add("label", "tugbot.docker.events=update")

	return client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}
