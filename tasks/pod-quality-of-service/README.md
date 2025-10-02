# Challenge: Pod Quality of Service

In this challenge, you will create three Pods, each with a different [Quality of Service class](https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/):

*   **Guaranteed**
*   **Burstable**
*   **BestEffort**

## Tasks

1.  **Guaranteed QoS:** Create a Pod named `resilient-workload-app-guaranteed` with a single container.

2.  **Burstable QoS:** Create a Pod named `resilient-workload-app-burstable` with a single container.

3.  **BestEffort QoS:** Create a Pod named `resilient-workload-app-besteffort` with a single container.

## Verification

After creating the Pods, you can verify their QoS class by running the following command and checking the `qosClass` field in the output:

```bash
~ kubectl get pod <pod-name> -o jsonpath='{.status.qosClass}'
```
