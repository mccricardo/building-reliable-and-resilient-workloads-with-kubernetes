# Solution: Exploring Kubernetes Deployment Strategies

This solution provides the Kubernetes manifests for the deployment strategies challenge.

## Recreate Deployment

This manifest creates a Deployment that uses the `Recreate` strategy. When you update the application version, Kubernetes will terminate all existing Pods before creating new ones.

**File:** `resilient-workload-app-recreate-deployment.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resilient-workload-app-recreate
spec:
  replicas: 4
  selector:
    matchLabels:
      app: resilient-workload-app-recreate
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: resilient-workload-app-recreate
    spec:
      containers:
      - name: resilient-workload-app
        image: resilient-workload-app:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3000
```

## RollingUpdate Deployment

This manifest creates a Deployment that uses the `RollingUpdate` strategy. This is the default strategy in Kubernetes. We've customized it with `maxSurge` and `maxUnavailable` to control the update process.

*   `maxSurge: 50%`: Allows the number of Pods to exceed the desired replica count by 50% during an update.
*   `maxUnavailable: 25%`: Ensures that at least 75% of the desired replica count is running at all times.

**File:** `resilient-workload-app-rollingupdate-deployment.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resilient-workload-app-rollingupdate
spec:
  replicas: 4
  selector:
    matchLabels:
      app: resilient-workload-app-rollingupdate
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 50%
      maxUnavailable: 25%
  template:
    metadata:
      labels:
        app: resilient-workload-app-rollingupdate
    spec:
      containers:
      - name: resilient-workload-app
        image: resilient-workload-app:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3000
```

## Observing the Deployments

Once you have created both deployments, open two new terminal windows.

**In the first new terminal**, watch the deployments:

```bash
watch kubectl get deployments
```

**In the second new terminal**, watch the pods:

```bash
~ watch kubectl get pods
```

Now, let's trigger an update for both deployments. You can do this by changing the image tag from `v1.0.0` to `v1.1.0`.

For the `Recreate` deployment:

```bash
~ kubectl set image deployment/resilient-workload-app-recreate resilient-workload-app=resilient-workload-app:nonexistent
```

For the `RollingUpdate` deployment:

```bash
~ kubectl set image deployment/resilient-workload-app-rollingupdate resilient-workload-app=resilient-workload-app:nonexistent
```

### What to Expect

**Recreate Strategy:** You will see all of the `resilient-workload-app-recreate` pods terminate at the same time. Then, new pods will be created with the new version of the application. This will cause a brief period of downtime.

**RollingUpdate Strategy:** You will see the `resilient-workload-app-rollingupdate` pods being updated one by one. Kubernetes will first create a new pod, wait for it to be ready, and then terminate an old pod. This ensures that the application remains available during the update. The `maxSurge` and `maxUnavailable` settings control how many pods can be created or be unavailable at the same time.
