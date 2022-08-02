## Introduction

This repo was used as a sanbox to test and play with K8s configurations.

## Setup with minikube

Note I have only been using this repo with `minikube` other settings may require a different setup.

To install minikube follow the [guide here](https://minikube.sigs.k8s.io/docs/start/).

### minikube with local images

Note that the `deployment` config will try to pull the image `aelnahas/echoserver` which I built locally with docke. To make K8S / minikube successfully pull the image you will need to run the following command:

```sh
eval $(minikube docker-env)
```

This will setup the necessary docker env variables.

After this step is complete you can build the `Dockerfile` included with this repo:

```sh
ls /path/to/k8s-lab
docker build -t aelnahas/echoserver:latest .
```

At this point you should be able to use the k8s yaml files.

### apply config files and verify setup

```sh
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

# to check deployment
kubectl get deployments
# to check pods
kubectl get pods
# to check service
kubectl get services
```

### Accessing the node with minikube

The network is limited if using the Docker driver on Darwin and the Node IP is not reachable directly.

We can expose our service by running:

```sh
minikube service echoserver-service --url
```

The command above will:

- open a tunnel
- return a url that you can access the pod with

you can ping the server to check that the setup is successful for e.g.:

```sh
curl http://127.0.0.1:53887/ping
```
