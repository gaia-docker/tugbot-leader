package swarm

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"time"
)

func NewComparator() *Comparator {
	return &Comparator{serviceIdToLastUpdate: make(map[string]time.Time)}
}

type Comparator struct {
	serviceIdToLastUpdate map[string]time.Time
}

func (c Comparator) GetUpdatedServices(client client.ServiceAPIClient) ([]string, error) {
	services, err := client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, currService := range services {
		if currService.UpdatedAt.After(c.serviceIdToLastUpdate[currService.ID]) {
			c.serviceIdToLastUpdate[currService.ID] = currService.UpdatedAt
			ret = append(ret, currService.ID)
		}
	}

	return ret, nil
}
