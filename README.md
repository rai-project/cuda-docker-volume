# RAI Docker Volume [![Build Status](https://travis-ci.org/rai-project/rai-docker-volume.svg?branch=master)](https://travis-ci.org/rai-project/rai-docker-volume)

Provides a Docker volume for mapping the NVIDIA driver into containers run by the RAI Daemon.

This is a Docker plugin. There are systemd scripts for starting it before the Docker service, and stopping it after the Docker service.

Some more info on plugin activation [here](https://docs.docker.com/engine/extend/plugin_api/#plugin-activation).

# Installation

## 1. Acquire a Binary

### Download a Binary

( not supported yet )

### Installation via "go install"

    go install github.com/rai-project/rai-docker-volume

### Build your own

Get the dependencies 

    go get -u -v ./...

or via glide:

    glide install

Then build the binary

    go build

## 2. Run the plugin

    nohup ./rai-docker-volume &

## (Optional) Set up the rai-docker-volume systemd service

Move the binary into the `/usr/lib/docker` directory. How you do this depends on how you got the binary.

Move the service and socket unit files into the appropriate places. On Ubuntu 16.04, for example:

    cp build/rai-docker-volume.service /lib/systemd/system/.
    cp build/rai-docker-volume.socket /lib/systemd/system/.

# Check
Reload the systemd daemon since you changed some files

    systemctl daemon-reload

Try startring the plugin. It should succeed, and automatically start the socket also.

    systemctl start rai-docker-volume.service
    systemctl status rai-docker-volume.socket
    systemctl status rai-docker-volume.service

Check that the rai-cuda volume is present:

    docker volume ls

## Troubleshooting steps


You can run the rai-docker-volume binary as root yourself, and it should work the same way.

If the plugin does not create a volume, you can try making one yourself and seing what happens.
Identify your nvidia driver version with `nvidia-smi` or through some other means.
Try creating a volume using that diver version. For example, for driver version `361.119`:

    docker volume create --driver=rai-cuda rai-cuda_361.119

