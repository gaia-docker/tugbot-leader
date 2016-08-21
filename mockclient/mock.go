package mockclient

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func NewMockClient() *MockClient {
	return &MockClient{}
}
