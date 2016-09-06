package swarm

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

const LabelTugbotEvents = "tugbot.swarm.events"

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
	services, err := s.comparator.GetUpdatedServices(s.client)
	if err != nil {
		return err
	}

	return s.doUpdate(services)
}

func (s ServiceUpdater) doUpdate(services []string) error {
	if len(services) == 0 {
		log.Debugf("No service were update")
		return nil
	}
	log.Debugf("Updated services (%v): %+v", len(services), services)

	testServices, err := s.getTestServices()
	if err != nil {
		return err
	}

	for _, currService := range testServices {
		log.Info("Rerunnig Service: ", currService)
		version, spec, err := s.getServiceInspect(currService.ID)
		if err != nil {
			log.Errorf("Failed to get service ID: %s (%+v)", currService.ID, err)
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
	filters.Add("label", LabelTugbotEvents+"=update")

	return s.client.ServiceList(context.Background(), types.ServiceListOptions{filters})
}

func (s ServiceUpdater) getServiceInspect(serviceId string) (swarm.Version, swarm.ServiceSpec, error) {
	service, _, err := s.client.ServiceInspectWithRaw(context.Background(), serviceId)
	if err != nil {
		return swarm.Version{}, swarm.ServiceSpec{}, err
	}
	version := service.Meta.Version

	return version, service.Spec, nil
}
