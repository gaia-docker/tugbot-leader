#!/bin/bash
distdir=.dist

go_build() {
  rm -rf "${distdir}"
  mkdir "${distdir}"
  go get -v
  go build -v -o ${distdir}/tugbot-leader
}

go_build
