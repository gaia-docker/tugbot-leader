package swarm

import (
	"os"
	"testing"

	"github.com/docker/engine-api/client"
	"github.com/stretchr/testify/assert"
)

func ITestSwarmUpdateServices(t *testing.T) {
	os.Setenv("DOCKER_CERT_PATH", "/home/effi/.docker/machine/certs/")
	os.Setenv("DOCKER_HOST", "tcp://192.168.99.100:2376")
	client, err := client.NewEnvClient()
	assert.NoError(t, err)

	err = NewServiceUpdater(client).Run()
	assert.NoError(t, err)
}
