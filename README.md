# RAI Docker Volume

Provides a Docker volume for mapping the NVIDIA driver into containers run by the RAI Daemon.

This is a Docker plugin. There are systemd scripts for starting it before the Docker service, and stopping it after the Docker service.

Some more info on plugin activation [here](https://docs.docker.com/engine/extend/plugin_api/#plugin-activation).

# Installation

## 1. Acquire a Binary

### Download a Binary

( not supported yet )

### Installation via "go install"

    go install github.com/rai-project/rai-docker-volume

## 2. Set up the rai-docker-volume systemd service

Move the binary into the `/usr/local/bin` directory.

    cp build/rai-docker-volume.service /lib/systemd/system/.
    cp build/rai-docker-volume.socket /lib/systemd/system/.


# Check

    docker volume create --driver=rai-cuda rai-cuda_361.119


# For developers 

Get the dependencies

    glide install


