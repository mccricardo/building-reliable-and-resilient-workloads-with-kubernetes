# Solution: Pod Quality of Service

This file contains the solutions for the Pod Quality of Service challenge.

## 1. Guaranteed QoS

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: resilient-workload-app-guaranteed
spec:
  containers:
  - name: resilient-workload-app
    image: resilient-workload-app:latest
    imagePullPolicy: IfNotPresent
    ports:
    - containerPort: 3000
    resources:
      requests:
        memory: "100Mi"
        cpu: "100m"
      limits:
        memory: "100Mi"
        cpu: "100m"
```

## 2. Burstable QoS

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: resilient-workload-app-burstable
spec:
  containers:
  - name: resilient-workload-app
    image: resilient-workload-app:latest
    imagePullPolicy: IfNotPresent
    ports:
    - containerPort: 3000
    resources:
      requests:
        memory: "100Mi"
        cpu: "100m"
```

## 3. BestEffort QoS

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: resilient-workload-app-bestffort
spec:
  containers:
  - name: resilient-workload-app
    image: resilient-workload-app:latest
    imagePullPolicy: IfNotPresent
    ports:
    - containerPort: 3000
```
