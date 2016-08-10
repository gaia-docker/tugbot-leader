package mockclient

import (
	"errors"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

// MockClient is mock implementation of container.Client which is a wrapper for Docker API.
type MockClient struct {
	mock.Mock
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (m *MockClient) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
	return types.ServiceCreateResponse{}, errors.New("Not Implemented")
}

func (m *MockClient) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error) {
	args := m.Called(ctx, options)
	return args.Get(0).([]swarm.Service), args.Error(1)
}

func (m *MockClient) ServiceInspectWithRaw(ctx context.Context, serviceID string) (swarm.Service, []byte, error) {
	return swarm.Service{}, nil, errors.New("Not Implemented")
}
func (m *MockClient) ServiceRemove(ctx context.Context, serviceID string) error {
	return errors.New("Not Implemented")
}
func (m *MockClient) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) error {
	return errors.New("Not Implemented")
}
func (m *MockClient) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error) {
	return swarm.Task{}, nil, errors.New("Not Implemented")
}
func (m *MockClient) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error) {
	return []swarm.Task{}, errors.New("Not Implemented")
}
