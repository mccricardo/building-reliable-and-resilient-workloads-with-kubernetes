# Solution for Challenge: Liveness, Readiness and Startup Probes

Here's how you can solve the challenge.

First, create a new deployment file, for example `resilient-workload-app-probes.yaml`.

## 1. Startup, Liveness and Readiness Probes

Here is a complete deployment manifest with all three probes configured:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resilient-workload-app
spec:
  replicas: 1
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
        startupProbe:
          httpGet:
            path: /start
            port: 3000
          failureThreshold: 30
          periodSeconds: 1
        livenessProbe:
          httpGet:
            path: /live
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 5
        readinessProbe:
          httpGet:
            path: /ready
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 5
```

A few things to note:
- The `startupProbe` has a `failureThreshold` of 30 and a `periodSeconds` of 1, giving the application 30 seconds to start.
- The `livenessProbe` and `readinessProbe` have an `initialDelaySeconds` of 5 to give the application some time to start before they begin probing.

Deploy this with `kubectl apply -f esilient-workload-app-probes.yaml`.

You can check the events of the pod with `kubectl describe pod resilient-workload-app`.

## 2. Testing Failure Scenarios

### Liveness Probe Failure

Update the `livenessProbe` in your `resilient-workload-app-probes.yaml` to use the `/live_fail` endpoint:

```yaml
        livenessProbe:
          httpGet:
            path: /live_fail
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 5
```

Apply the changes and watch the pod's restarts with `kubectl get pods -w`. You will see the `RESTARTS` count increase.

### Readiness Probe Failure

Update the `readinessProbe` to use the `/ready_fail` endpoint:

```yaml
        readinessProbe:
          httpGet:
            path: /ready_fail
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 5
```

Apply the changes. You will see that the pod never reaches the `READY` state (e.g., `0/1`). If you have a service selecting this pod, it will not receive traffic.
