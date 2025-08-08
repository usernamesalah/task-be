# Task 1: System Health & Monitoring 🩺

### Backend Service Status

**To check the current status of the NestJS backend service managed by `systemd`, you would use the `systemctl status` command:**

```bash
systemctl status your-nestjs-app.service
```
This command provides detailed output, including whether the service is active, its process ID (PID), and recent log entries.

**If the service were inactive, you could attempt to restart it with `systemctl start`:**

```bash
sudo systemctl start your-nestjs-app.service
```

### Database Container Status

**To list all currently running Docker containers, the standard command is `docker ps`:**

```bash
docker ps
```

**To confirm that the PostgreSQL and MongoDB containers are healthy, a few steps are recommended:**

1.  **Check the container logs** for any apparent errors using `docker logs`:
    ```bash
    docker logs <postgresql_container_name>
    docker logs <mongodb_container_name>
    ```

2.  **Perform a health check** by executing a client command inside the container.
    *   For PostgreSQL, `pg_isready` is a quick way to verify that the database is accepting connections:
        ```bash
        docker exec -it <postgresql_container_name> pg_isready -U <user>
        ```
    *   For MongoDB, you can run a simple command to check the connection status:
        ```bash
        docker exec -it <mongodb_container_name> mongo --eval 'db.runCommand({ connectionStatus: 1 })'
        ```

### Server Resource Monitoring

**Two of the most critical metrics to monitor for server health are CPU utilization and memory usage.** Consistently high CPU can indicate an inefficient process, while running out of memory can cause services to fail. Disk space is another important metric to watch.

**Several command-line tools can be used to quickly check these resources:**

*   **`htop`**: An interactive process viewer that gives a real-time overview of CPU, memory usage, and running processes. It's often preferred over the standard `top`.
*   **`free -h`**: Displays the total, used, and free memory on the system in a human-readable format.
*   **`df -h`**: Shows the disk space usage for all mounted filesystems in a human-readable format.

### Backend Unresponsive: Initial Diagnosis

If users report that the backend API is unresponsive, a systematic approach is needed to diagnose the issue. Here are the initial steps:

1.  **Check Service Status:** First, verify that the NestJS `systemd` service is running.
    ```bash
    systemctl status your-nestjs-app.service
    ```
2.  **Inspect Recent Logs:** Review the service's logs for any errors or unusual activity.
    ```bash
    journalctl -u your-nestjs-app.service -n 100 --no-pager
    ```
3.  **Assess Server Resources:** Use `htop` or `top` to check for resource exhaustion. A spike in CPU or memory usage could be the root cause.
4.  **Verify Network Ports:** Ensure the application is listening on the expected network port.
    ```bash
    sudo netstat -tulnp | grep <port_number>
    ```
5.  **Check Firewall Rules:** Confirm that the server's firewall is not blocking incoming traffic on the required port.
    ```bash
    sudo ufw status
    ```
6.  **Review Application-Level Issues:** If the service is running and accessible, the problem may lie within the application itself, such as a database connection issue or an unhandled exception. The application logs from `journalctl` are the best place to find these clues.