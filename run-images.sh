#!/usr/bin/env sh

docker run -d -p 3000:3000 -v testVolume:/app/images wasaphotobe

docker run -d -p 3000:3000 wasaphotobe
docker run -d -p 80:80 wasaphotofe