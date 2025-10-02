# Challenge: Liveness, Readiness and Startup Probes

In this challenge, we'll explore how to use liveness, readiness and startup probes to improve the reliability of our application.

## 1. Startup Probe

Our application has a `/start` endpoint that simulates a slow start. Create a new deployment for our application and configure a `startupProbe` to use this endpoint. The probe should give the application enough time to start.

## 2. Liveness Probe

Add a `livenessProbe` to the deployment that uses the `/live` endpoint. This probe will restart the container if the application becomes unresponsive.

## 3. Readiness Probe

Add a `readinessProbe` to the deployment that uses the `/ready` endpoint. This probe will remove the Pod from service endpoints if the application is not ready to serve traffic.

## 4. Testing the Probes

Deploy the application and observe the events to see the probes in action. You can use `kubectl describe pod <pod-name>` to see the probe events.

## 5. Testing Failure Scenarios

Now, let's see what happens when the probes fail.

### Liveness Probe Failure

Update the deployment to use the `/live_fail` endpoint for the `livenessProbe`. Deploy the changes and observe how Kubernetes restarts the container repeatedly.

### Readiness Probe Failure

Update the deployment to use the `/ready_fail` endpoint for the `readinessProbe`. Deploy the changes and observe how the Pod is not marked as "Ready". If you have a service pointing to this deployment, you will see that the pod is removed from the service endpoints.
