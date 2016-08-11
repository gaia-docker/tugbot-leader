package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

func OnServiceUpdate(client client.ServiceAPIClient, updatedServices []string) error {
	if len(updatedServices) == 0 {
		return nil
	}

	testServices, err := getTestServices(client)
	if err != nil {
		return err
	}

	for _, currService := range testServices {
		log.Debug("Rerunnig Service: ", currService)
		client.ServiceUpdate(context.Background(), currService.ID,
			swarm.Version{}, swarm.ServiceSpec{}, types.ServiceUpdateOptions{})
	}

	return nil
}

func getTestServices(client client.ServiceAPIClient) ([]swarm.Service, error) {
	filters := filters.NewArgs()
	filters.Add("label", "tugbot.docker.events=update")

	return client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}
