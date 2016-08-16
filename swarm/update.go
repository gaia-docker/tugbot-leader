package swarm

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

func (s ServiceUpdater) Run() error {
	log.Debug("Polling for updated services...")
	services, err := s.comparator.GetUpdatedServices(s.client)
	if err != nil {
		return err
	}
	log.Debug("Updated services: ", services)

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
		log.Info("Rerunnig Service: ", currService)
		// Sending an empty swarm.ServiceSpec{} result:
		// Error response from daemon: ContainerSpec: image reference must be provided
		// Setup ServiceSpec (includes image) using service inspect, result:
		// Error response from daemon: update out of sequence
		s.client.ServiceUpdate(context.Background(), currService.ID,
			swarm.Version{}, s.getServiceSpec(currService.ID), types.ServiceUpdateOptions{})
	}

	return nil
}

func (s ServiceUpdater) getTestServices() ([]swarm.Service, error) {
	filters := filters.NewArgs()
	filters.Add("label", "tugbot.docker.events=update")

	return s.client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}

func (s ServiceUpdater) getServiceSpec(serviceId string) swarm.ServiceSpec {
	service, _, _ := s.client.ServiceInspectWithRaw(context.Background(), serviceId)

	return service.Spec
}
