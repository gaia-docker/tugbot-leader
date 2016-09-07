# tugbot-leader
[![CircleCI](https://circleci.com/gh/gaia-docker/tugbot-leader.svg?style=shield)](https://circleci.com/gh/gaia-docker/tugbot-leader)
[![codecov](https://codecov.io/gh/gaia-docker/tugbot-leader/branch/master/graph/badge.svg)](https://codecov.io/gh/gaia-docker/tugbot-leader)
[![Go Report Card](https://goreportcard.com/badge/github.com/gaia-docker/tugbot-leader)](https://goreportcard.com/report/github.com/gaia-docker/tugbot-leader)
[![Docker](https://img.shields.io/docker/pulls/gaiadocker/tugbot-leader.svg)](https://hub.docker.com/r/gaiadocker/tugbot-leader/)
[![Docker Image Layers](https://imagelayers.io/badge/gaiadocker/tugbot-leader:latest.svg)](https://imagelayers.io/?images=gaiadocker/tugbot-leader:latest 'Get your own badge on imagelayers.io')

**Tugbot Leader** is a Continuous Testing Framework for Docker Swarm based production/staging/testing environment. **Tugbot Leader** executes *Docker Swarm Test Services* upon some *event*.

## Test Service

*Test Service* is a regular Docker swarm service. Docker `LABEL` is used to discover *test service* and **Tugbot** related test metadata. These labels can be specified at runtime, using `--label` `docker service create` option.
**Tugbot Leader** will trigger a sequential *test service* execution upon *event* (see `tugbot.event.swarm` label).

### Tugbot labels

All **Tugbot** labels must be prefixed with `tugbot.` to avoid potential conflict with other labels.
**Tugbot** labels are divided into:

1) Container labels:
- `tugbot.results.dir` - directory, where *test container* reports test results; default to `/var/tests/results`

2) Swarm Service labels:

- `tugbot.swarm.events` - list of comma separated Docker Swarm events (*currently only service update is supported*)

#####Example Docker Swarm Test Service creation:
```bash
docker service create --network my_net --replicas 1 --restart-condition none --label tugbot.swarm.event=update --name mytest my-test-img
```
> It is highly recommended to determain `--restart-condition none` when creating a test service. Otherwise, Swarm will restart test service all the time (default “any”).

> Use `--label tugbot.swarm.event=update` to tell tugbot framework that this is a test service that should be updated each time an application service is updated.

## Running Tugbot Leader inside a Docker container
```bash
docker run -d -e DOCKER_HOST=<IP:Port> -e DOCKER_CERT_PATH=<Docker Certificate Path> --log-driver=json-file --name tugbot-leader gaiadocker/tugbot-leader
```
- `DOCKER_HOST` - IP:Port Docker Swarm **Master** host (**Tugbot Leader** should run as part of Docker Swarm *Master* in order to update Docker Swarm Services).
- `DOCKER_CERT_PATH` - dirctory should contain: ca.pem, cert.pem and key.pem. While using Docker secure communication.
- `DOCKER_TLS_VERIFY` - Use TLS when connecting to the Docker socket and verify the server's certificate.
- `TUGBOT_LEADER_INTERVAL` - Interval time between polling Docker Swarm for list of updated services (currently docker is not exposing Swarm events, this is why Tugbot Leader is polling the Swarm master node). An interval string is a possibly signed sequence of decimal numbers, each with optional fraction and a time unit suffix, such as "300s", "1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
- `TUGBOT_LEADER_LOG_LEVEL` - Enable debug mode. When this option set to `debug` you'll see more verbose logging in the Tugbot Leader log file.
