package util

import (
	"os"
)

const DockerCertPath string = "DOCKER_CERT_PATH"
const DockerHost string = "DOCKER_HOST"
const TugbotInterval string = "TUGBOT_LEADER_INTERVAL"
const TugbotLogLevel string = "TUGBOT_LEADER_LOG_LEVEL"

func SetEnv() {
	os.Setenv(DockerCertPath, "/home/effi/.docker/machine/certs/")
	os.Setenv(DockerHost, "tcp://192.168.99.100:2376")
	os.Setenv(TugbotInterval, "7s")
	os.Setenv(TugbotLogLevel, "debug")
}
