package swarm

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"time"
)

type Comparator struct {
	serviceIdToLastUpdate map[string]time.Time
}

func NewComparator() *Comparator {
	return &Comparator{serviceIdToLastUpdate: make(map[string]time.Time)}
}

func (c Comparator) Initialize(client client.ServiceAPIClient) error {
	services, err := client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Errorf("Failed initialize comparator (%v)", err)
		return err
	}

	for _, currService := range services {
		c.serviceIdToLastUpdate[currService.ID] = currService.UpdatedAt
	}

	return nil
}

func (c Comparator) GetUpdatedServices(client client.ServiceAPIClient) ([]string, error) {
	services, err := client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, currService := range services {
		if currService.Spec.Labels[LabelTugbotEvents] == "" &&
			currService.UpdatedAt.After(c.serviceIdToLastUpdate[currService.ID]) {
			c.serviceIdToLastUpdate[currService.ID] = currService.UpdatedAt
			ret = append(ret, currService.ID)
		}
	}

	return ret, nil
}
