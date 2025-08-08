# Task 6: Simple CI/CD for a Dockerized Application ⚙️

### Pipeline Trigger

A CI/CD pipeline is typically triggered automatically by an event in a version control repository. The most common trigger for a production deployment is a **push to the main branch** (e.g., `main` or `master`).

For other environments, different triggers can be used:
*   **Development/Staging:** A push to a `develop` branch.
*   **Feature Testing:** The creation of or a push to a pull request (PR).

This ensures that code is automatically built, tested, and deployed upon integration, streamlining the development lifecycle.

### Core Pipeline Stages

A typical CI/CD pipeline for a containerized application consists of several essential stages, each with a distinct purpose:

1.  **Checkout:** The pipeline begins by checking out the latest version of the source code from the specified branch in the repository.
2.  **Test:** Automated tests (such as unit, integration, and end-to-end tests) are executed to validate the new code. This stage acts as a quality gate, preventing bugs from proceeding further.
3.  **Build:** If all tests pass, the pipeline uses the `Dockerfile` to build the application into a Docker image. This image is a portable, self-contained package of the application and its dependencies.
4.  **Push to Registry:** The newly created Docker image is tagged (e.g., with the commit hash or a version number) and pushed to a container registry, such as Docker Hub, GitHub Container Registry (GHCR), or Amazon ECR.
5.  **Deploy:** The final stage involves deploying the new image to the server. The pipeline connects to the production environment, pulls the latest image from the registry, and updates the running application.

### Role of Docker in the Build Stage

The `Dockerfile` is central to the "build" stage. It is a text file that contains a series of instructions on how to assemble the application image layer by layer. The CI/CD system's build agent reads this file and executes the commands to create a reproducible and isolated environment for the application. This includes:
*   Starting from a base image (e.g., `node:18-alpine`).
*   Installing system dependencies.
*   Copying the application source code into the image.
*   Installing application dependencies (e.g., via `npm install`).
*   Defining the command to run when a container is started from the image.

The purpose of a **Docker image registry** is to store and distribute these images. It acts as a centralized repository where the CI/CD pipeline can push newly built images and from which the production server can pull them for deployment.

### Deployment to Server

The "deployment" stage can be automated using SSH to run commands on the target server. The script would perform the following actions to update the running service with minimal disruption:

1.  **Pull the new image:** `docker pull your-registry/your-image:latest`
2.  **Stop the running container:** `docker stop old_container_name`
3.  **Remove the old container:** `docker rm old_container_name`
4.  **Start a new container from the new image:** `docker run -d --name new_container_name -p 8080:8080 your-registry/your-image:latest`

This sequence ensures that the old version is removed cleanly before the new one starts. For zero-downtime deployments, more advanced strategies like blue-green deployments could be implemented.

#### Example CI/CD Pipeline with GitHub Actions

Here is a conceptual example of a CI/CD pipeline defined in a GitHub Actions workflow file (`.github/workflows/ci-cd.yml`).

```yaml
# Example CI/CD pipeline using GitHub Actions
name: CI/CD Pipeline

# Trigger the pipeline on every push to the main branch
on:
  push:
    branches: [ main ]

jobs:
  build-and-deploy:
    # Use the latest Ubuntu runner
    runs-on: ubuntu-latest
    steps:
      # 1. Checkout the repository code
      - name: Checkout code
        uses: actions/checkout@v3

      # 2. Set up Docker Buildx for building images
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      # 3. Log in to the Docker registry
      - name: Log in to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_HUB_USERNAME }}
          password: ${{ secrets.DOCKER_HUB_ACCESS_TOKEN }}

      # 4. Build the Docker image and push it to the registry
      - name: Build and push Docker image
        uses: docker/build-push-action@v4
        with:
          context: .
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKER_HUB_USERNAME }}/your-app:latest

      # 5. Deploy the new image to the server via SSH
      - name: Deploy to Server
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.SERVER_HOST }}
          username: ${{ secrets.SERVER_USERNAME }}
          key: ${{ secrets.SSH_PRIVATE_KEY }}
          script: |
            # Ensure the script exits if any command fails
            set -e
            
            # Pull the latest image from the registry
            docker pull ${{ secrets.DOCKER_HUB_USERNAME }}/your-app:latest
            
            # Stop and remove the old container if it exists
            docker stop your-app-container || true
            docker rm your-app-container || true
            
            # Start a new container with the updated image
            docker run -d \
              --name your-app-container \
              -p 8080:8080 \
              ${{ secrets.DOCKER_HUB_USERNAME }}/your-app:latest
```