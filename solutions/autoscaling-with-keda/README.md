# Solution for Challenge: Autoscaling with KEDA

Here's how you can solve the challenge.

## 1. Install KEDA

As mentioned in the challenge, you need to install KEDA in your cluster.

```bash
kubectl apply -f https://github.com/kedacore/keda/releases/download/v2.12.0/keda-2.12.0.yaml
```

## 2. Create a Deployment

First, let's create a simple deployment for our application. Create a file named `resilient-workload-app.yaml`:

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
```
Deploy it: `kubectl apply -f resilient-workload-app.yaml`

## 3. Create a ScaledObject

Now, let's create the `ScaledObject`. Create a file named `resilient-workload-app-scaledobject.yaml`:

```yaml
apiVersion: keda.sh/v1alpha1
kind: ScaledObject
metadata:
  name: cron-scaledobject
spec:
  scaleTargetRef:
    name: resilient-workload-app
  triggers:
  - type: cron
    metadata:
      timezone: "Etc/UTC"
      start: "* 8 * * *"
      end: "* 19 * * *" 
      desiredReplicas: "3"
  cooldownPeriod: 30
```

A few things to note:
- `scaleTargetRef` points to our deployment.
- The `cron` trigger is configured to fire every minute (`* * * * *`).
- `desiredReplicas` is the number of replicas to scale to when the cron expression is met.
- `cooldownPeriod` is the period to wait after the last trigger reported active before scaling back to 0.

Apply the `ScaledObject`: `kubectl apply -f resilient-workload-app-scaledobject.yaml`

## 4. Watch KEDA in Action

Watch the pods of the deployment:

```bash
~ kubectl get pods -l app=resilient-workload-app -w
```

You will see KEDA scaling the deployment up to 3 replicas every minute. After the `cooldownPeriod`, it will scale it back down.

You can also check the HPA that KEDA creates and manages for you:
```bash
~ kubectl get hpa
```
