#!/usr/bin/env sh
# Con docker desktop attivo
docker build -f Dockerfile.backend -t wasaphotobe .
docker build -f Dockerfile.frontend -t wasaphotofe .