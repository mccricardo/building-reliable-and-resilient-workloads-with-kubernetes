# Challenge: PriorityClass

In this challenge, we'll explore how to use PriorityClass to ensure that high-priority workloads are scheduled even when resources are scarce.

## 1. The Scenario: A Full Cluster

First, we need to create a situation where our cluster is "full". We'll create a deployment that consumes a significant amount of resources on the worker nodes, leaving no room for new pods.

## 2. The Problem

Try to deploy a new pod. You will see that it gets stuck in the `Pending` state because there are not enough resources.

## 3. Create PriorityClasses

Now, let's create two PriorityClasses: `high-priority` and `low-priority`.

## 4. Test Preemption

Update the "filler" deployment to use the `low-priority` class.

Create a new deployment for our application and assign it the `high-priority` class.

Deploy the high-priority application and observe how Kubernetes preempts (evicts) the low-priority pods to make room for the high-priority ones.
