# Challenge 3: Exploring Replicas, Deployments, and Scaling

In this challenge, we'll explore Kubernetes ReplicaSets, Deployments, and manual scaling.

## 1. Launch a single Pod

Create a Pod definition file for the `resilient-workload-app` from the `app` folder. Launch the Pod and verify it's running.

Now, delete the Pod and see that it doesn't come back. This demonstrates that a Pod is a transient object.

## 2. Use a ReplicaSet

Create a ReplicaSet definition file that launches three replicas of the `resilient-workload-app`.

Delete one of the Pods and observe that Kubernetes automatically creates a new one to maintain the desired number of replicas.

## 3. Use a Deployment

Create a Deployment definition file that launches three replicas of the `resilient-workload-app`.

Again, delete one of the Pods and see that a new one is created.

## 4. Update and Rollback a Deployment

Update the Deployment to use a new version of the application (e.g., by changing the image tag). Observe how the Deployment performs a rolling update.

Now, imagine the new version has a bug. Rollback the Deployment to the previous version and see how the old Pods are brought back.

## 5. Manual Scaling

Use `kubectl` to manually scale the number of replicas in your Deployment up to 5. Then, scale it back down to 2.