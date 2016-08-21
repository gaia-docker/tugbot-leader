package swarm

import (
	"errors"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gaia-docker/tugbot-leader/mockclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestIsValidNode_Error(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("Info", mock.Anything).Return(types.Info{}, errors.New("Expected :)")).Once()
	assert.Error(t, IsValidNode(client))
	client.AssertExpectations(t)
}

func TestIsValidNode_NotSwarmMaster(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("Info", mock.Anything).Return(types.Info{
		Swarm: swarm.Info{
			LocalNodeState:   swarm.LocalNodeStateActive,
			ControlAvailable: false}},
		nil).Once()
	assert.Error(t, IsValidNode(client))
	client.AssertExpectations(t)
}

func TestIsValidNode_NotActive(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("Info", mock.Anything).Return(types.Info{
		Swarm: swarm.Info{
			LocalNodeState:   swarm.LocalNodeStateInactive,
			ControlAvailable: true}},
		nil).Once()
	assert.Error(t, IsValidNode(client))
	client.AssertExpectations(t)
}

func TestIsValidNode(t *testing.T) {
	client := mockclient.NewMockClient()
	client.On("Info", mock.Anything).Return(types.Info{
		Swarm: swarm.Info{
			LocalNodeState:   swarm.LocalNodeStateActive,
			ControlAvailable: true}},
		nil).Once()
	assert.NoError(t, IsValidNode(client))
	client.AssertExpectations(t)
}
