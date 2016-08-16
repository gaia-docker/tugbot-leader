package util

import "os"

func SetEnv() {
	os.Setenv("DOCKER_CERT_PATH", "/home/effi/.docker/machine/certs/")
	os.Setenv("DOCKER_HOST", "tcp://192.168.99.100:2376")
}
