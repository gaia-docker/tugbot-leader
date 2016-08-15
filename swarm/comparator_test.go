package swarm

import (
	"errors"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gaia-docker/tugbot-leader/mockclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNotify_Error(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("ServiceList", mock.Anything, mock.Anything).Return([]swarm.Service{}, errors.New("Expected :)")).Once()
	_, err := NewComparator().GetUpdatedServices(client)
	assert.Error(t, err)
	client.AssertExpectations(t)
}

func TestNotify(t *testing.T) {
	swarmServices := []swarm.Service{{ID: "service-id", Meta: swarm.Meta{UpdatedAt: time.Now()}}}
	client := mockclient.NewMockClient()
	client.On("ServiceList", mock.Anything, mock.Anything).Return(swarmServices, nil).Once()
	services, err := NewComparator().GetUpdatedServices(client)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(services))
	assert.Equal(t, "service-id", services[0])
	client.AssertExpectations(t)
}
