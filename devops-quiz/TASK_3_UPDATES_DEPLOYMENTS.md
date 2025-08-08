# Task 3: Updates & Deployments 🚀

### OS Package Updates

Safely updating the packages on an Ubuntu server is a routine but critical task. The process should be handled carefully to avoid unintended disruptions.

**The standard procedure involves two commands:**
1.  First, resynchronize the package index files from their sources.
    ```bash
    sudo apt update
    ```
2.  Then, install the newest versions of all packages currently installed on the system.
    ```bash
    sudo apt upgrade
    ```

**Important Considerations:**
*   **Maintenance Window:** Perform updates during a scheduled maintenance window to minimize impact on users.
*   **Reviewing Changes:** On critical systems, it's wise to first inspect which packages will be upgraded using `apt list --upgradable` before proceeding.
*   **Backups:** Ensure a recent server snapshot or backup is available in case a package update causes problems.

### Backend Application Update

Deploying a new pre-built version of the NestJS application requires a clear, step-by-step process to ensure a smooth transition.

**The general deployment workflow is as follows:**

1.  **Transfer Files:** Securely copy the new application build (e.g., a `.tar.gz` archive) to the server.
2.  **Stop the Service:** Gracefully stop the running application to prevent new connections.
    ```bash
    sudo systemctl stop your-nestjs-app.service
    ```
3.  **Backup Current Version:** Create a backup of the existing application directory. This allows for a quick rollback if needed.
    ```bash
    cp -a /path/to/app /path/to/app-backup-$(date +%F)
    ```
4.  **Deploy New Files:** Remove the old application files and extract the new version into the target directory.
5.  **Update Dependencies:** If required, install any new or updated Node.js packages.
    ```bash
    cd /path/to/app && npm install --production
    ```
6.  **Restart the Service:** Start the application using `systemd`.
    ```bash
    sudo systemctl start your-nestjs-app.service
    ```
7.  **Verify Deployment:** Check the service's status and review its logs to confirm that it started successfully and is running without errors.
    ```bash
    systemctl status your-nestjs-app.service
    journalctl -u your-nestjs-app.service -f
    ```

### Frontend Deployment with Vercel

When a Next.js frontend is connected to a Git repository and hosted on Vercel, the deployment process is highly automated and designed for efficiency and safety.

Updates are typically handled as follows:
1.  **Git Push:** A developer pushes new commits to the production branch (e.g., `main`).
2.  **Webhook Trigger:** The Git provider sends a webhook to Vercel, initiating a new deployment.
3.  **Build & Deploy:** Vercel automatically pulls the latest code, installs dependencies, builds the Next.js application, and deploys it to a unique, immutable URL.
4.  **Atomic Deployment:** Once the new deployment is built and passes health checks, Vercel atomically switches the production domain to point to the new version. This process is seamless and results in **zero downtime** for end-users.