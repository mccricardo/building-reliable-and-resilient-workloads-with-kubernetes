# Challenge: Pod Disruption Budgets

In this challenge, we'll explore how to use Pod Disruption Budgets (PDBs) to protect our application from voluntary disruptions. Your Kind cluster has one control-plane and two worker nodes.

## 1. The Scenario: Draining a Node

First, let's create a deployment with multiple replicas of our application.

Next, we'll simulate a voluntary disruption by draining one of the worker nodes.

## 2. The Problem

Drain one of the worker nodes and observe what happens to the pods. You will see that all pods on that node are terminated at the same time, which could cause a service disruption if all your pods are on that node.

## 3. Create a Pod Disruption Budget

Now, create a Pod Disruption Budget that ensures at least one replica of our application is always available.

## 4. Test the PDB

With the PDB in place, try draining one of the worker nodes again. Observe how Kubernetes now respects the PDB, and the drain command will hang until the pods are rescheduled to the other node, ensuring that the service remains available.
