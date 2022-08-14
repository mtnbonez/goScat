# goScat
Text-based, Go implementation of the [(in)famous card game](https://en.wikipedia.org/wiki/Scat_(card_game)).

# Deployment

## Minikube

[Minikube](https://minikube.sigs.k8s.io/docs/start/) is a great tool to spin up a local VM & host a local Kubernetes cluster, regardless of your base platform OS. From here, you can mount the project in the Minikube VM filesystem to do manual container building.

> I highly recommend looking into the Hyper-V driver for Minikube if you're on Windows and you're working on projects outside of Docker Desktop's terms for free-use. 

```
> minikube start --mount-string="D:\projects\goScat\:/mount_dir/" --mount 
```

If you're having troubles mounting on start, try doing them separate operations:

```
> minikube start
> minikube mount "D:\projects\goScat\:/mount_dir/"
```

# Building

## Building images in Minikube

One of my favorite features of Minikube is having a place (regardless of underlying OS) to build container images. Using the following command, you can SSH into the machine & do docker builds

```
> minikube ssh
```

With the mount point specified earlier, you can access the `build` directory & do Dockerfile build shenanigans here.

```
> cd /mount_dir/
```

From here, you can build the current Dockerfile with the following options:

```
> docker build --network=host -t goscat:latest -f ./build/Dockerfile .
```

Afterwards, you can view your newly built image in the Docker image list:

```
> docker image ls
```

> If you get complaints about file space, you can do a quick prune of the unused images doing `docker images prune`