package swarm

import (
	"testing"

	"github.com/docker/engine-api/client"
	"github.com/gaia-docker/tugbot-leader/util"
	"github.com/stretchr/testify/assert"
)

func ITestSwarmUpdateServices(t *testing.T) {
	util.SetEnv()
	client, err := client.NewEnvClient()
	assert.NoError(t, err)

	err = NewServiceUpdater(client).Run()
	assert.NoError(t, err)
}
