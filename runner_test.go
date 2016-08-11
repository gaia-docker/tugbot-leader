package main

import (
	"errors"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gaia-docker/tugbot-leader/mockclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOnServiceUpdate_ErrorServiceList(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("ServiceList", mock.Anything, mock.Anything).Return([]swarm.Service{}, errors.New("Expected :)")).Once()
	err := OnServiceUpdate(client, []string{"1"})
	assert.Error(t, err)
	client.AssertExpectations(t)
}

func TestOnServiceUpdate_EmptyUpdatedServices(t *testing.T) {
	client := mockclient.NewMockClient()
	err := OnServiceUpdate(client, []string{})
	assert.NoError(t, err)
	client.AssertExpectations(t)
}

func TestOnServiceUpdate(t *testing.T) {
	const testServiceId = "test-service-id"
	client := mockclient.NewMockClient()
	testServices := []swarm.Service{{ID: testServiceId}}
	client.On("ServiceList", mock.Anything, mock.Anything).Return(testServices, nil).Once()
	client.On("ServiceUpdate",
		mock.Anything,
		mock.MatchedBy(func(serviceId string) bool {
			return testServiceId == serviceId
		}), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	err := OnServiceUpdate(client, []string{"updated-service-id"})
	assert.NoError(t, err)
	client.AssertExpectations(t)
}
