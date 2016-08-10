#!/bin/bash
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  go get
  go build -v -o ${distdir}/tugbot-swarm
}

go_build
