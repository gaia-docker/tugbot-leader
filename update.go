package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

func NewServiceUpdater(client client.ServiceAPIClient) *ServiceUpdater {
	return &ServiceUpdater{client: client, comparator: NewComparator()}
}

type ServiceUpdater struct {
	client     client.ServiceAPIClient
	comparator *Comparator
}

func (s ServiceUpdater) Update() error {
	services, err := s.comparator.GetUpdatedServices(s.client)
	if err != nil {
		return err
	}

	return s.onServiceUpdate(services)
}

func (s ServiceUpdater) onServiceUpdate(updatedServices []string) error {
	if len(updatedServices) == 0 {
		return nil
	}

	testServices, err := s.getTestServices()
	if err != nil {
		return err
	}

	for _, currService := range testServices {
		log.Debug("Rerunnig Service: ", currService)
		s.client.ServiceUpdate(context.Background(), currService.ID,
			swarm.Version{}, swarm.ServiceSpec{}, types.ServiceUpdateOptions{})
	}

	return nil
}

func (s ServiceUpdater) getTestServices() ([]swarm.Service, error) {
	filters := filters.NewArgs()
	filters.Add("label", "tugbot.docker.events=update")

	return s.client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}
