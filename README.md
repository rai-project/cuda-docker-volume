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

### Build your own

Get the dependencies

    glide install

Build the binary

    go build

## 2. Set up the rai-docker-volume systemd service

Move the binary into the `/usr/local/bin` directory. How you do this depends on how you got the binary.

Move the service and socket unit files into the appropriate places. On Ubuntu 16.04, for example:

    cp build/rai-docker-volume.service /lib/systemd/system/.
    cp build/rai-docker-volume.socket /lib/systemd/system/.


# Check
Reload the systemd daemon since you changed some files

    systemctl daemon-reload

Try starting the socket. It should succeed.

    systemctl start rai-docker-volume.socket
    systemctl status rai-docker-volume.socket

Try startring the plugin. It should succeed.

    systemctl start rai-docker-volume.service
    systemctl status rai-docker-volume.service

Identify your nvidia driver version with `nvidia-smi` or through some other means.

Try creating a volume using that diver version. For example, for driver version `361.119`:

    docker volume create --driver=rai-cuda rai-cuda_361.119


