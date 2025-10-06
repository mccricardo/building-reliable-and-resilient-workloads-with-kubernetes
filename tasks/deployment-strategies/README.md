# Challenge: Exploring Kubernetes Deployment Strategies

In this challenge, you'll explore two fundamental Kubernetes deployment strategies: `Recreate` and `RollingUpdate`. You'll create two separate Deployments to observe how each strategy behaves when updating an application.

## Your Task

1.  **Create a `Recreate` Deployment:**
    *   Define a new Deployment named `resilient-workload-app-recreate`.
    *   Configure the deployment strategy to be `Recreate`.
    *   Save the manifest as `resilient-workload-app-recreate-deployment.yaml`.

2.  **Create a `RollingUpdate` Deployment:**
    *   Define a second Deployment named `resilient-workload-app-rollingupdate`.
    *   Configure the deployment strategy to be `RollingUpdate`.
    *   Set the `maxSurge` to `50%` and `maxUnavailable` to `25%`.
    *   Save the manifest as `resilient-workload-app-rollingupdate-deployment.yaml`.

## Observing the Deployments

Once you have created both deployments, open two new terminal windows.

**In the first new terminal**, watch the deployments:

```bash
watch kubectl get deployments
```

**In the second new terminal**, watch the pods:

```bash
watch kubectl get pods
```

Now, let's trigger an update for both deployments. You can do this by changing the image tag from `v1.0.0` to `v1.1.0`.

For the `Recreate` deployment:

```bash
kubectl set image deployment/resilient-workload-app-recreate resilient-workload-app=alexcfy/resilient-workload-app:v1.1.0
```

For the `RollingUpdate` deployment:

```bash
kubectl set image deployment/resilient-workload-app-rollingupdate resilient-workload-app=alexcfy/resilient-workload-app:v1.1.0
```

**Observe the output in your terminal windows. What do you notice?**

*   How does the `Recreate` strategy update the pods?
*   How does the `RollingUpdate` strategy update the pods?
*   What is the impact on the availability of the application during each update?

Good luck!
