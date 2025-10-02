# Solution for Challenge 3: Exploring Replicas, Deployments, and Scaling

Here's how you can solve the challenge.

## 1. Launch a single Pod

First, you need to have the `resilient-workload-app` image built and available in your local Docker registry.

Create a file named `resilient-workload-app-pod.yaml`:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: resilient-workload-app
  labels:
    app: resilient-workload-app
spec:
  containers:
  - name: resilient-workload-app
    image: resilient-workload-app:latest
    imagePullPolicy: IfNotPresent
    ports:
    - containerPort: 3000
```

Launch the Pod:

```bash
~ kubectl apply -f resilient-workload-app-pod.yaml
```

Verify it's running:

```bash
~ kubectl get pods
```

Delete the Pod:

```bash
~ kubectl delete pod resilient-workload-app
```

You'll see that the Pod is terminated and not restarted.

## 2. Use a ReplicaSet

Create a file named `resilient-workload-app-replicaset.yaml`:

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: resilient-workload-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: resilient-workload-app
  template:
    metadata:
      labels:
        app: resilient-workload-app
    spec:
      containers:
      - name: resilient-workload-app
        image: resilient-workload-app:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3000
```

Launch the ReplicaSet:

```bash
~ kubectl apply -f resilient-workload-app-replicaset.yaml
```

Get the Pods:

```bash
~ kubectl get pods
```

Delete one of the Pods and you'll see a new one being created automatically.

## 3. Use a Deployment

Create a file named `resilient-workload-app-deployment.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resilient-workload-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: resilient-workload-app
  template:
    metadata:
      labels:
        app: resilient-workload-app
    spec:
      containers:
      - name: resilient-workload-app
        image: resilient-workload-app:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3000
```

Launch the Deployment:

```bash
~ kubectl apply -f deployment.yaml
```

Again, you can delete a Pod and see it being replaced.

## 4. Update and Rollback a Deployment

To update the deployment, you can change the image in the `resilient-workload-app-deployment.yaml` file to a new version (e.g., `resilient-workload-app:v2`) and re-apply it.

```bash
~ kubectl apply -f resilient-workload-app-deployment.yaml
```

You can check the rollout status:

```bash
~ kubectl rollout status deployment/resilient-workload-app
```

You can check the created ReplicaSets:

```bash
~ kubectl rollout status deployment/resilient-workload-app
```

To see the history of rollouts:

```bash
~ kubectl rollout history deployment/resilient-workload-app
```

To rollback to the previous version:

```bash
~ kubectl rollout undo deployment/resilient-workload-app
```

## 5. Manual Scaling

To manually scale a Deployment, you can use the `kubectl scale` command.

Scale up to 5 replicas:

```bash
~ kubectl scale deployment resilient-workload-app --replicas=5
```

Verify the number of Pods:

```bash
~ kubectl get pods
```

Scale down to 2 replicas:

```bash
~ kubectl scale deployment resilient-workload-app --replicas=2
```

And verify again:

```bash
~ kubectl get pods
```