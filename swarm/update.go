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
		version, spec, err := s.getServiceInspect(currService.ID)
		if err != nil {
			log.Errorf("Failed to get service ID: %s (%v)", currService.ID, err)
		} else {
			err = s.client.ServiceUpdate(context.Background(),
				currService.ID,
				version,
				spec,
				types.ServiceUpdateOptions{})
			if err != nil {
				log.Errorf("Failed to update service (Service ID: %s, %v)", currService.ID, err)
			}
		}
	}

	return nil
}

func (s ServiceUpdater) getTestServices() ([]swarm.Service, error) {
	filters := filters.NewArgs()
	filters.Add("label", "tugbot.docker.events=update")

	return s.client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}

func (s ServiceUpdater) getServiceInspect(serviceId string) (swarm.Version, swarm.ServiceSpec, error) {
	service, _, err := s.client.ServiceInspectWithRaw(context.Background(), serviceId)
	if err != nil {
		return swarm.Version{}, swarm.ServiceSpec{}, err
	}
	version := service.Meta.Version

	return version, service.Spec
}
