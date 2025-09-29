# Resilient Workload App

This document provides instructions on how to build, run, and publish the Docker container image for this application.

## Prerequisites

*   Docker is installed and running on your local machine.
*   [You are authenticated with the GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry). 

## Build the Image

To build the Docker image, navigate to the `app` directory and run the following command:

```sh
~ docker build -t resilient-workload-app .
```

## Run the Image

To run the Docker image, use the following command:

```sh
~ docker run -d -p 3000:3000 --name resilient-workload-app resilient-workload-app
```

## Stop and Delete the IMage

To stop and delete the Docker image, use the following command:

```sh
~ docker rm -f resilient-workload-app
```

This will start the container in detached mode and map port 3000 of the container to port 3000 on your host machine. You can access the application at [http://localhost:3000](http://localhost:3000).

## Push the Image to GitHub Container Registry

To push the image to the GitHub Container Registry, you first need to tag it appropriately. Replace `OWNER` with your GitHub username or organization name.

```sh
~ docker tag resilient-workload-app ghcr.io/OWNER/resilient-workload-app:latest
```

Now, you can push the image:

```sh
~ docker push ghcr.io/OWNER/resilient-workload-app:latest
```



export CR_PAT=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
