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
~ kubectl apply -f resilient-workload-app-hpa.yaml
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