# Solution for Challenge 5: Autoscaling with Horizontal Pod Autoscaler (HPA)

Here's how you can solve the challenge.

## 1. Prerequisites: Metrics Server

As mentioned in the challenge, you need the Metrics Server running in your cluster. Follow the instructions in the challenge to install it if you haven't already. If you are using Kind, make sure to apply the patch to the metrics-server deployment.

## 2. Generate CPU Load

With the deployment and service running, port-forward to the service:

```bash
~ kubectl port-forward svc/resilient-workload-app 3000:3000
```

And in another terminal, generate load:

```bash
~ while true; do curl http://localhost:3000/heavy; done
```

## 3. Create an HPA

You can create an HPA resource using a YAML file or with an imperative command.

**YAML file (`hpa.yaml`):**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: resilient-workload-app
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: resilient-workload-app
  minReplicas: 1
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 50
```

Apply the HPA:
```bash
~ kubectl apply -f hpa.yaml
```

**Imperative command:**
```bash
~ kubectl autoscale deployment resilient-workload-app --cpu-percent=50 --min=1 --max=10
```

## 4. Watch the HPA in Action

You can watch the HPA with the following command:

```bash
~ kubectl get hpa -w
```

You will see the `TARGETS` column showing the current CPU utilization vs the target (e.g., `150%/50%`). The `REPLICAS` column will show the number of pods, which will increase as the HPA scales up the deployment.

You can also watch the pods being created:
```bash
~ kubectl get pods -w
```

When you stop the load generation (by stopping the `curl` loop), the CPU utilization will drop. After a few minutes, the HPA will scale the deployment back down to 1 replica.

## 5. How it Works: Metrics Server and HPA

The Kubernetes Metrics Server collects resource usage data (CPU and memory) from each node's kubelet. It then exposes this data through the Kubernetes Metrics API.

The Horizontal Pod Autoscaler (HPA) controller periodically queries this API to get the metrics for the pods it's configured to scale. Based on the current metric values and the target values you define, the HPA decides whether to increase or decrease the number of replicas.

You can query the Metrics API yourself to see the raw data. The following command will show you the CPU and memory usage for all pods:

```bash
~ kubectl get --raw "/apis/metrics.k8s.io/v1beta1/pods" | jq .
```
(The `jq` command is optional, but it makes the output much more readable).
