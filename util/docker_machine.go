package util

import "os"

const DockerCertPath string = "DOCKER_CERT_PATH"
const DockerHost string = "DOCKER_HOST"

func SetEnv() {
	os.Setenv(DockerCertPath, "/home/effi/.docker/machine/certs/")
	os.Setenv(DockerHost, "tcp://192.168.99.100:2376")
}
