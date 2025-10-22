# Building Reliable and Resilient Workloads in Kubernetes

## Requirements
  - [Github account](https://github.com/)
  - [git](https://git-scm.com/)
  - [Docker](https://www.docker.com/get-started)
  - [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
  - [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
  - [curl](https://curl.se/) (optional)


## Create Kuberentes cluster
```sh
~  kind create cluster --config kind/config.yaml
```

## Delete Kubernetes cluster
```sh
~  kind delete cluster -n local-cluster
```

## Tasks and Challenges

- [Deploying an Application as a Pod](tasks/deploy-application-as-pod/README.md) [(Solution)](solutions/deploy-application-as-pod/README.md)
- [Pod Quality of Service](tasks/pod-quality-of-service/README.md) [(Solution)](solutions/pod-quality-of-service/README.md)
- [Replicas, Deployments, and Scaling](tasks/replicas-deployments-and-scaling/README.md) [(Solution)](solutions/replicas-deployments-and-scaling/README.md)
- [Exposing a Deployment with a Service](tasks/exposing-deployment-with-service/README.md) [(Solution)](solutions/exposing-deployment-with-service/README.md)
- [Liveness, Readiness, and Startup Probes](tasks/liveness-readiness-startup-probes/README.md) [(Solution)](solutions/liveness-readiness-startup-probes/README.md)
- [Deployment Strategies](tasks/deployment-strategies/README.md) [(Solution)](solutions/deployment-strategies/README.md)
- [Autoscaling with HPA](tasks/autoscaling-with-hpa/README.md) [(Solution)](solutions/autoscaling-with-hpa/README.md)
- [Pod Disruption Budgets](tasks/pod-disruption-budgets/README.md) [(Solution)](solutions/pod-disruption-budgets/README.md)