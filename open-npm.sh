#!/usr/bin/env sh

docker run -it --rm -v "$(pwd):/src" -u "$(id -u):$(id -g)" --network host --workdir /src/webui node:lts /bin/bash

# This command will start a Docker container with the node:lts image and run a bash shell inside the container.

# Here's what each of the flags does:

# -it: This flag specifies that the container should be run in interactive mode and allocate a pseudo-TTY. This allows you to interact with the container through the command line.
# --rm: This flag tells Docker to automatically remove the container when it exits.
# -v "$(pwd):/src": This flag mounts the current working directory ($(pwd)) on the host machine as /src in the container. This allows you to access the files in the current working directory from within the container.
# -u "$(id -u):$(id -g)": This flag sets the user and group ID of the process running inside the container to the current user and group ID on the host machine. This allows you to access files on the host machine with the correct permissions.
# --network host: This flag tells Docker to use the host machine's network stack for the container. This allows the container to access the host machine's network resources as if it were running directly on the host.
# --workdir /src/webui: This flag sets the working directory for the bash shell inside the container to /src/webui.
# Finally, node:lts /bin/bash specifies that the bash shell should be run inside the container. node:lts is the name of the Docker image that the container is based on. It's an image with Node.js LTS (long-term support) version installed.
