package mockclient

import (
	"errors"
	"io"

	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

// Mock implementation for client.SystemAPIClient
// see: github.com/docker/engine-api/client/interface.go

func (m *MockClient) Events(ctx context.Context, options types.EventsOptions) (io.ReadCloser, error) {
	return nil, errors.New("Not Implemented")
}

func (m *MockClient) Info(ctx context.Context) (types.Info, error) {
	args := m.Called(ctx)
	return args.Get(0).(types.Info), args.Error(1)
}

func (m *MockClient) RegistryLogin(ctx context.Context, auth types.AuthConfig) (types.AuthResponse, error) {
	return types.AuthResponse{}, errors.New("Not Implemented")
}
