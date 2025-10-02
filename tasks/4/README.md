# Challenge 4: Exposing your Deployment with a Service

In the previous challenges, we've launched our application using a Deployment. While we can access the application by port-forwarding to a single Pod, this is not a scalable or resilient solution. If that Pod dies, we lose access.

In this challenge, we'll explore how to use a Kubernetes Service to expose our application to the outside world.

## 1. The Problem with Port-Forwarding

If you have the Deployment from the previous challenge running, try to port-forward to one of the Pods and access the application.

Now, delete that specific Pod. You'll lose access to the application, even though the Deployment will bring up a new Pod.

## 2. Create a Service

Create a Kubernetes Service of type `ClusterIP` that exposes your `resilient-workload-app` Deployment.

## 3. Access the Service

Since the Service is of type `ClusterIP`, it's only accessible from within the cluster. Use `kubectl port-forward` to forward a local port to the Service.

Now, access the application through the forwarded port. Try deleting a Pod again and see that you can still access the application without interruption.
