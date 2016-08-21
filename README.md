# tugbot-leader
[![CircleCI](https://circleci.com/gh/gaia-docker/tugbot-leader.svg?style=shield)](https://circleci.com/gh/gaia-docker/tugbot-leader)
[![codecov](https://codecov.io/gh/gaia-docker/tugbot-leader/branch/master/graph/badge.svg)](https://codecov.io/gh/gaia-docker/tugbot-leader)
[![Go Report Card](https://goreportcard.com/badge/github.com/gaia-docker/tugbot-leader)](https://goreportcard.com/report/github.com/gaia-docker/tugbot-leader)
[![Docker](https://img.shields.io/docker/pulls/gaiadocker/tugbot-leader.svg)](https://hub.docker.com/r/gaiadocker/tugbot-leader/)
[![Docker Image Layers](https://imagelayers.io/badge/gaiadocker/tugbot-leader:latest.svg)](https://imagelayers.io/?images=gaiadocker/tugbot-leader:latest 'Get your own badge on imagelayers.io')

**Tugbot Leader** is an Continuous Testing Framework for Docker Swarm based production/staging/testing environment. **Tugbot Leader** executes *Docker Swarm Test Services* on some *event*.

## Test Service

*Test Service* is a regular Docker swarm service. We use Docker `LABEL` to discover *test service* and **Tugbot** related test metadata. These labels can be part specified at runtime, using `--label` `docker service create` option.
**Tugbot Leader** will trigger a sequential *test service* execution on *event* (see `tugbot.event.swarm` label).

### Tugbot labels

All **Tugbot** labels must be prefixed with `tugbot.` to avoid potential conflict with other labels.
**Tugbot** labels divided into:

1) Container labels:
- `tugbot.results.dir` - directory, where *test container* reports test results; default to `/var/tests/results`

2) Swarm Service labels:

- `tugbot.event.swarm` - list of comma separated Docker Swarm events (*currently only service update supported*)

#####Example Docker Swarm Test Service creation:
```bash
docker service create --network my_net --replicas 1 --restart-condition none --label tugbot.swarm.event=update --name mytest my-test-img
```
> It is higly recomanded to determain `--restart-condition none` when creating a test service. Otherwise, Swarm will restart test service all the time.

> Use `--label tugbot.swarm.event=update` to tell tugbot framework that this is a test service that should be updated each time that an application service been updated.

## Running Tugbot Leader inside a Docker container
```bash
docker run -d -e DOCKER_CERT_PATH=<Docker Certificate Path> -e DOCKER_HOST=<Swarm Master IP:Port> --name tugbot-leader --log-driver=json-file gaiadocker/tugbot-leader
```
