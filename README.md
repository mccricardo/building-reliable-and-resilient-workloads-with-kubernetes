# Building Reliable and Resilient Workloads in Kubernetes

## Requirements
  - [Github account](https://github.com/)
  - [git](https://git-scm.com/)
  - [Docker](https://www.docker.com/get-started)
  - [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
  - [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
  - [Helm](https://helm.sh/docs/intro/install/)
  - [curl](https://curl.se/) (optional)


## Create Kuberentes cluster
```sh
~  kind create cluster --config kind/config.yaml
```

## Delete Kubernetes cluster
```sh
~  kind delete cluster -n local-cluster
```