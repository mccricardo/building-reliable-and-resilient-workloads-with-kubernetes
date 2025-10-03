# Challenge: Autoscaling with KEDA

In this challenge, we'll explore how to use KEDA (Kubernetes-based Event-driven Autoscaling) to scale our application based on a cron schedule.

## 1. Prerequisites: Install KEDA

First, you need to install KEDA in your cluster. You can do this using Helm or by applying the YAML files directly. We will use the YAML files.

If you get an error about annotations being too long, try adding the `--server-side` flag to the command.

```bash
kubectl apply -f https://github.com/kedacore/keda/releases/download/v2.12.0/keda-2.12.0.yaml --server-side
```
This will install all the necessary KEDA components in the `keda` namespace.

Verify the installation:
```bash
kubectl get pods -n keda
```
You should see the KEDA operator and metrics server pods running.

## 2. The Scenario: Scheduled Scaling

Imagine you have a workload that needs to be scaled up during business hours and scaled down during off-hours. KEDA's `cron` scaler is perfect for this.

## 3. The Challenge

Create a `ScaledObject` that targets our `resilient-workload-app` deployment. Configure it to scale the deployment to 3 replicas every minute, and then scale it back down to 1 replica after a cooldown period.

## 4. Watch KEDA in Action

Deploy the `ScaledObject` and observe how KEDA scales the deployment up and down based on the cron schedule.