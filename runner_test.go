package main

import (
	"errors"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gaia-docker/tugbot-leader/mockclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOnServiceUpdate(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("ServiceList", mock.Anything, mock.Anything).Return([]swarm.Service{}, errors.New("Expected :)")).Once()
	err := OnServiceUpdate(client, []string{"1"})
	assert.Error(t, err)
	client.AssertExpectations(t)
}
