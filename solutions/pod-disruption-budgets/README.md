# Solution for Challenge: Pod Disruption Budgets

Here's how you can solve the challenge.

## 1. Create a Deployment

First, create a deployment with 3 replicas. You can use a file like this (`resilient-workload-app.yaml`):

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: resilient-workload-app
spec:
  replicas: 6
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

Deploy it: `kubectl apply -f resilient-workload-app-resilient-workload-app-pdb.yaml`

You can see which nodes the pods are running on with `kubectl get pods -o wide`. You should see pods on both worker nodes.

## 2. Drain a Worker Node (without PDB)

First, get the names of your worker nodes. They will be something like `local-cluster-worker` and `local-cluster-worker2`.

```bash
~ kubectl get nodes
```

Now, drain one of the worker nodes. The `--ignore-daemonsets` flag is needed because of the Kind CNI daemonset.

```bash
~ kubectl drain local-cluster-worker --ignore-daemonsets --delete-emptydir-data
```

You will see that the pods on that node are terminated.

After the drain is complete, you need to uncordon the node to make it schedulable again:

```bash
~ kubectl uncordon local-cluster-worker
```

## 3. Create a Pod Disruption Budget

Create a file named `resilient-workload-app-pdb.yaml`:

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: resilient-workload-app
spec:
  minAvailable: 3
  selector:
    matchLabels:
      app: resilient-workload-app
```

This PDB will ensure that at least 1 replica of the `resilient-workload-app` is always available.

Apply it: `kubectl apply -f resilient-workload-app-pdb.yaml`

## 4. Test the PDB

Now, with the PDB in place, try draining one of the worker nodes again:

```bash
~ kubectl drain local-cluster-worker --ignore-daemonsets --delete-local-data
```

This time, you will see that `kubectl` respects the PDB. The drain command will not proceed until the pods from the drained node are rescheduled and running on the other worker node, ensuring no downtime.

Don't forget to uncordon the node afterwards:
```bash
~ kubectl uncordon local-cluster-worker
```
