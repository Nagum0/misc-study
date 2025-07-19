# DOCKER NOTES

## DOCKER COMMANDS

- **docker build \<options\> \<dockerfile_path\>** = Builds an image from a dockerfile
  - **-t \<name\>:\<tag\>** = sets the name and the tag of the image **REQUIRED**
- **docker pull \<image\>** = downloads a new image
- **docker push \<image\>** = pushes a docker image to a remote repository
  - needs a domain url before the image name and tag; you can use the tag command from below for setting this up
- **docker tag \<image\>:\<tag\> \<new_name\>:\<tag\>** = sets a new name and tag for the image (technically it copies the image keeping the old one too)
- **docker run \<image\>** = starts a container from an image
  - **-d \<image\>** = container will run in the background in detached mode (prints the id)
  - **-p \<host_port\>:\<container_port\> \<image\>** = starts a container and binds a port
  - **--name \<given_name\> \<image\>** = starts a container with the given name
  - **--net \<network\>** = connect to the specified network
  - **--rm** = remove container after it exits
- **docker rm \<id\>** = remove a container
- **docker rmi \<image_id\>** = remove an image
- **docker stop \<id\>** = stops the container
- **docker start \<id\>** = start a stopped container
- **docker images** = lists all the downloaded images
- **docker ps** = lists all of the running docker containers
  - **-a** = lists all of the containers (even the stopped ones)
- **docker logs \<id\>** = shows you the logs of the container
- **docker exec \<id\>** = run a new process inside of a container
  - **-it \<id\> /bin/\<shell\>** = run an interactive bash terminal inside of the container
- **docker network** = command for working with docker networks
  - **ls** = lists all of the available docker networks
  - **create \<name\>** = creates a new docker network

## DOCKER COMPOSE

- Used to orchestrate multiple docker containers locally or on a server
- Docker compose will create a common network for the listed containers (it takes care of container network related issues)

### DOCKER COMPOSE YAML CONFIG FILE

- Always start with docker compose version
``` yaml
version: <version number>
```
- In the services section you list your containers and their specifications
``` yaml
services:
  container_name:
    # Image to run container from
    image: <docker image> 
    # -p flag
    ports: 
      - <host>:<container>
    # Environment variables
    environment:
      - <VAR=value>
```

### DOCKER COMPOSE COMMANDS

- **docker-compose**
  - **-f \<filename\>** = add the config yaml file
  - **up** = run the listed docker containers
  - **down** = stops all of the containers of the specified config yaml file and closes the created network

## DOCKER FILES

### DOCKERFILE COMMANDS

- **FROM \<image\>** = Specifies which existing image our image will be built upon
``` Dockerfile
FROM ubuntu:22.04
``` 
- **ENV <key>=<value>** = Sets the images environment variables
  - It's better to set these in a docker-compose file for convenience
``` Dockerfile
ENV APP_VERSION=1.0.0 DEBUG_MODE=false
```
- **RUN \<linux_command\>** = Run any linux command inside of a container
``` Dockerfile
RUN mkdir /home/app
```
- **WORKDIR \<path\>** = Sets the working directory of the container for commands like: RUN, ENTRYPOINT, CMD, etc
- **COPY \<host_dir\> \<container_dir\>** = Copy files from host to container
``` Dockerfile
COPY . /home/app
```
- **EXPOSE \<port\>** = informs Docker that the container listens on the specified network ports at runtime
``` Dockerfile
EXPOSE 8080
```
- **CMD [\<command_list\>]** = The entrypoint command of the container
``` Dockerfile
CMD ["node", "server.js"]
```

## DOCKER VOLUMES

- Folder of physical host filt system is mounted into the virtual file system of a container.

### DOCKER VOLUME COMMANDS

- **docker volume** = allows the inspection of docker volumes
  - **ls** = lists all of the existing docker volumes

### DOCKER VOLUME TYPES

#### HOST VOLUMES

- **docker run -v \<host_path\>:\<container_path\>** = mounts the container_path to the host_path and all future changes will be stored and presisted at the home_path

#### ANONYMOUS VOLUMES

- **docker run -v \<container_path\>** = we omit the host_path and the mounting point will automatically created and handled by docker in the background
- For each container that uses this volume type docker will create a mount point at */var/lib/docker/volumes*; this mount point is a folder created for the container

#### NAMED VOLUMES

- This is a better version of the type mentioned above where docker creates the mountpoint but you specify the name of it (the name of the folder)
- **docker run -v \<name\>:\<container_path\>** = container_path will be mounted at */var/lib/docker/volumes/\<name\>*
- This one should be used in production
``` yaml
# This how to mount in docker-compose yaml file
services:
  database:
    volumes:
      - db_data:<container_path>
```