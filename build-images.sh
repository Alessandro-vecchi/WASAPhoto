#!/usr/bin/env sh

docker build -f Dockerfile.backend -t wasaphotobe .
docker build -f Dockerfile.frontend -t wasaphotofe .