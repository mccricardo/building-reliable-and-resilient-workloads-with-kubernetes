# Solution for Challenge: PriorityClass

Here's how you can solve the challenge.

## 1. Create a "Filler" Deployment

First, let's create a deployment that will consume most of the resources on our worker nodes. Create a file named `filler-deployment.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: filler-deployment
spec:
  replicas: 8 # Adjust as needed for your cluster size
  selector:
    matchLabels:
      app: filler
  template:
    metadata:
      labels:
        app: filler
    spec:
      containers:
      - name: filler-container
        image: nginx
        resources:
          requests:
            cpu: "1000m"
```
The number of replicas and the cpu request might need to be adjusted depending on the size of your worker nodes. The goal is to make the nodes appear full.

Deploy it: `kubectl apply -f filler-deployment.yaml`

Check the status of the nodes to see the resource allocation: `kubectl describe nodes`

## 2. Try to Deploy a New Pod

Now, try to deploy a new pod without a priority class. It will get stuck in `Pending` state.

## 3. Create PriorityClasses

Create a file named `priority-classes.yaml`:

```yaml
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: high-priority
value: 1000000
globalDefault: false
description: "This priority class should be used for high priority service pods."

---
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: low-priority
value: 1000
globalDefault: false
description: "This priority class should be used for low priority pods."
```

Apply them: `kubectl apply -f priority-classes.yaml`

## 4. Test Preemption

First, update your `filler-deployment.yaml` to use the `low-priority` class by adding `priorityClassName: low-priority` to the pod spec:

```yaml
//...
    spec:
      priorityClassName: low-priority
      containers:
//...
```
Apply the change.

Now, create a new deployment for our application with the `high-priority` class. Create a file named `high-priority-app.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: high-priority-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: high-priority-app
  template:
    metadata:
      labels:
        app: high-priority-app
    spec:
      priorityClassName: high-priority
      containers:
      - name: resilient-workload-app
        image: resilient-workload-app:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3000
```

Deploy it: `kubectl apply -f high-priority-app.yaml`

Now, watch the pods: `kubectl get pods -w`. You will see that one of the `filler-deployment` pods is terminated and the `high-priority-app` pod is scheduled and starts running. This is preemption in action.
