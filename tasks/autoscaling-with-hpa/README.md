# Challenge 5: Autoscaling with Horizontal Pod Autoscaler (HPA)

In this challenge, we'll explore how to use the Horizontal Pod Autoscaler (HPA) to automatically scale our application based on CPU utilization.

## 1. Prerequisites: Metrics Server

The HPA needs a source of metrics to make scaling decisions. The most common one is the Kubernetes Metrics Server.

First, check if you have the Metrics Server installed:

```bash
~ kubectl get pods -n kube-system | grep metrics-server
```

If you don't see a running pod, you need to install it.

For most clusters, you can install it using the following command:
```bash
~ kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

**For Kind users:** If you are using Kind, the Metrics Server might not work out of the box. After installing it, you will need to patch the deployment to allow insecure TLS connections to the kubelet:
```bash
~ kubectl patch deployment metrics-server -n kube-system --type='json' -p='[{"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-insecure-tls"}]'
```

It might take a minute or two for the Metrics Server to start collecting metrics. You can check if it's working by running:
```bash
~ kubectl top pods
```
If you see CPU and memory usage for your pods, you are good to go.

## 2. Generate CPU Load

Our application has a `/heavy` endpoint that simulates CPU load. We need to send traffic to this endpoint to trigger the HPA.

First, make sure you have the deployment and service from the previous challenges running.

In a terminal, port-forward to the service:
```bash
~ kubectl port-forward svc/resilient-workload-app 3000:3000
```

In another terminal, we will use a simple loop to generate load. You can use `curl` or any other tool.
```bash
~ while true; do curl http://localhost:3000/heavy; done
```

## 3. Create an HPA

Create a HorizontalPodAutoscaler that targets our `resilient-workload-app` deployment. Configure it to maintain an average CPU utilization of 50%. Set the minimum number of replicas to 1 and the maximum to 10.

## 4. Watch the HPA in Action

Observe the HPA as it detects the increased CPU load and starts scaling up the number of pods. You can use `kubectl get hpa` to see the status of the HPA.

Once you stop the load generation, watch the HPA scale the deployment back down to the minimum number of replicas.