#!/usr/bin/env sh
# Con docker desktop attivo
docker build -f Dockerfile.backend -t wasaphotobe .
docker build -f Dockerfile.frontend -t wasaphotofe .

# -f option specifies the Dockerfile
# -t option specifies the name and tag to apply to the image
# . current directory