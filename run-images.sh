#!/usr/bin/env sh

# The purpose of creating two volumes instead of one is to 
# separate the storage needs of the application into two separate volumes with different paths. 
# The first volume, named "testVolume," is mounted at "/app/images" and is intended to store image data. 
# The second volume, named "tempVolume," is mounted at "/tmp" and is intended to store temporary data.
# This separation of storage provides better organization and isolation for the application's data, 
# and can help to ensure that critical data is not lost in case of issues with the tempVolume.
docker volume create databaseVolume
docker volume create imageVolume
docker run --rm -d -p 3000:3000 -v databaseVolume:/app/images -v imageVolume:/tmp wasaphotobe

docker run --rm -d -p 80:80 wasaphotofe

# -d option starts the container in detached mode (in the background), 
# -p option publishes port 3000 of the container to port 3000 of the host machine.
# -v option is used to mount the named volumes databaseVolume and imageVolume to the container's filesystem. 