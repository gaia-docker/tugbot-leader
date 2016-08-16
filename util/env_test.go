package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	SetEnv()
	assert.True(t, os.Getenv(DockerCertPath) != "")
	assert.True(t, os.Getenv(DockerHost) != "")
}
