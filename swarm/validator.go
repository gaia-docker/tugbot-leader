package swarm

import (
	"errors"
	"fmt"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

func IsValidNode(client client.SystemAPIClient) error {
	info, err := client.Info(context.Background())
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to get docker node info (%+v)", err))
	} else if !info.Swarm.ControlAvailable {
		return errors.New(fmt.Sprintf("Node is not a swarm master (%s)", info.Swarm.NodeAddr))
	} else if info.Swarm.LocalNodeState != swarm.LocalNodeStateActive {
		return errors.New(fmt.Sprintf("Inactive Node (State: %s)", info.Swarm.LocalNodeState))
	}

	return nil
}
