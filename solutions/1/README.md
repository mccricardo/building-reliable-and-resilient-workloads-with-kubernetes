# Solution for Task 1: Deploy the Application as a Pod

This document provides the solution for deploying the application as a Kubernetes Pod.

## Steps:

1.  **Build the Docker Image:**

    Navigate to the `app` directory and build the Docker image.

    ```bash
    ~ docker build -t resilient-workload-app .
    ```

2.  **Load the Image into `kind`:**

    If you are using `kind` for your local Kubernetes cluster, you can load the image directly into the cluster nodes.

    ```bash
    ~ kind -n local-cluster load docker-image resilient-workload-app
    ```

3.  **Create a Pod Manifest:**

    Create a file named `resilient-workload-app-pod.yaml` with the following content:

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: resilient-workload-app
    spec:
      containers:
      - name: resilient-workload-app
        image: resilient-workload-app:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3000
    ```

4.  **Deploy the Pod:**

    Apply the manifest to your cluster:

    ```bash
    ~ kubectl apply -f pod.yaml
    ```

5.  **Verify the Pod:**

    Check the status of the Pod:

    ```bash
    ~ kubectl get pods
    ```

    You should see the `resilient-workload-app` pod in a `Running` state. You can also check the logs:

    ```bash
    ~ kubectl logs resilient-workload-app
    ```

6.  **Access the Application:**

    [Port-forward](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/) to the Pod:
    ```bash
    ~ kubectl port-forward reslient-workload-app 3000:3000
    ```

    Access the application at [localhost:3000](localhost:3000). User your browser or cURL.