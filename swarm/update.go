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
	c := NewComparator()
	c.Initialize(client)

	return &ServiceUpdater{client: client, comparator: c}
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

	return s.doUpdate(services)
}

func (s ServiceUpdater) doUpdate(updatedServices []string) error {
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
		err = s.client.ServiceUpdate(context.Background(), currService.ID,
			swarm.Version{}, s.getServiceSpec(currService.ID), types.ServiceUpdateOptions{})
		if err != nil {
			log.Debugf("Swarm service update return an error, but ervice probebly will run anyway. ", err)
		}
	}

	return nil
}

func (s ServiceUpdater) getTestServices() ([]swarm.Service, error) {
	filters := filters.NewArgs()
	filters.Add("label", "tugbot.docker.events=update")

	return s.client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}

func (s ServiceUpdater) getServiceSpec(serviceId string) swarm.ServiceSpec {
	service, _, err := s.client.ServiceInspectWithRaw(context.Background(), serviceId)
	if err != nil {
		log.Errorf("Failed to get service ID: %s (%v)", serviceId, err)
		return swarm.ServiceSpec{}
	}

	return service.Spec
}
