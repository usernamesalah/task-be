# Task 4: Logging & Troubleshooting 🔍

### Backend Application Logs

For a service managed by `systemd`, the `journalctl` command is the standard tool for viewing logs. It provides powerful options for filtering and displaying log data.

To view the logs for the NestJS backend service, you can use the following commands:

```bash
# View the most recent 100 log entries and follow new logs in real-time.
journalctl -u your-nestjs-app.service -n 100 -f

# View all logs for the service, with output piped to a pager for navigation.
journalctl -u your-nestjs-app.service
```

### Database Container Logs

Accessing logs for Docker containers is straightforward using the `docker logs` command.

*   **For the PostgreSQL container:**
    ```bash
    # View the last 100 lines of logs and follow new logs.
    docker logs --tail 100 -f <postgresql_container_name>
    ```

*   **For the MongoDB container:**
    ```bash
    # The command is identical, just specify the MongoDB container name.
    docker logs --tail 100 -f <mongodb_container_name>
    ```

### Full-Stack Troubleshooting Scenario

**Scenario:** A user reports seeing an error page on the Next.js frontend after submitting a form.

A full-stack issue like this requires a methodical approach to isolate the problem. Here is an effective troubleshooting strategy, starting from the user-facing component and working backwards.

1.  **Examine the Frontend:**
    *   **Vercel Logs:** Start by checking the Vercel dashboard. Look at the runtime logs for the frontend application to see if any errors were captured during server-side rendering or API route execution.
    *   **Browser Developer Tools:** Replicate the issue in a web browser. Open the developer tools and inspect the **Console** for any client-side JavaScript errors. Check the **Network** tab to see the details of the form submission request. Look at the HTTP status code, request payload, and response from the server. This is often the quickest way to identify the source of the error.

2.  **Investigate the Backend:**
    *   **Backend Service Status:** SSH into the DigitalOcean server and confirm the NestJS service is active: `systemctl status your-nestjs-app.service`.
    *   **Backend Logs:** If the service is running, check its logs for errors corresponding to the time of the form submission. This will show if the request was received and how it was processed.
        ```bash
        journalctl -u your-nestjs-app.service -n 200 --no-pager
        ```

3.  **Check the Databases:**
    *   **Database Logs:** If the backend logs point to a database problem (e.g., connection error, failed query), inspect the logs of the relevant Docker container (`docker logs <db_container_name>`). This can reveal underlying issues with the database itself.

This layered approach allows you to efficiently determine whether the fault lies with the client-side code, the backend API, or the database layer.