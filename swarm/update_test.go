package swarm

import (
	"errors"
	"testing"
	"time"

	"github.com/docker/engine-api/types/swarm"
	"github.com/gaia-docker/tugbot-leader/mockclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateServices_ErrorServiceList(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("ServiceList", mock.Anything, mock.Anything).Return([]swarm.Service{}, errors.New("Expected :)")).Once()
	err := NewServiceUpdater(client).Run()
	assert.Error(t, err)
	client.AssertExpectations(t)
}

func TestUpdateServices_EmptyUpdatedServices(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("ServiceList", mock.Anything, mock.Anything).Return([]swarm.Service{}, nil).Once()
	err := NewServiceUpdater(client).Run()
	assert.NoError(t, err)
	client.AssertExpectations(t)
}

func TestUpdateServices(t *testing.T) {
	const testServiceId = "test-service-id"
	client := mockclient.NewMockClient()

	// get updated services
	updatedServices := []swarm.Service{{ID: "service-1", Meta: swarm.Meta{UpdatedAt: time.Now()}}}
	client.On("ServiceList", mock.Anything, mock.Anything).Return(updatedServices, nil).Once()

	// get test services
	testServices := []swarm.Service{{ID: testServiceId}}
	client.On("ServiceList", mock.Anything, mock.Anything).Return(testServices, nil).Once()

	// update test services
	client.On("ServiceUpdate",
		mock.Anything,
		mock.MatchedBy(func(serviceId string) bool {
			return testServiceId == serviceId
		}), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	err := NewServiceUpdater(client).Run()
	assert.NoError(t, err)
	client.AssertExpectations(t)
}
