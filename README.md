# goScat
Text-based, Go implementation of the [(in)famous card game](https://en.wikipedia.org/wiki/Scat_(card_game)).

# Deployment

## Minikube

[Minikube]() is a great tool to spin up a local VM & host a local Kubernetes cluster, regardless of your base platform

> I highly recommend looking into the Hyper-V driver for Minikube if you're on Windows and you're working on projects outside of Docker Desktop's terms for free-use. 

```
> minikube start --mount-string="D:\projects\goScat\:/mount_dir/" --mount 
```

If you're having troubles mounting on start, try doing them separate operations:

```
> minikube start
> minikube mount "D:\projects\goScat\:/mount_dir/"
```