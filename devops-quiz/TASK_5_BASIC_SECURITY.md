# Task 5: Basic Security 🛡️

### Essential Server Security Practices

Maintaining the security of a production server is critical. For an Ubuntu server on DigitalOcean, here are two fundamental security practices that should always be implemented.

#### 1. Configure a Host-Based Firewall

A firewall is the first line of defense, controlling what network traffic is allowed to enter or leave the server. The default firewall tool on Ubuntu is `ufw` (Uncomplicated Firewall).

**Best Practices:**
*   **Principle of Least Privilege:** Only open the ports that are absolutely necessary for your applications to function.
*   **Common Ports:** For this setup, you would typically allow SSH (port 22), HTTP (port 80), and HTTPS (port 443).

**Example `ufw` Configuration:**
```bash
# Deny all incoming traffic by default
sudo ufw default deny incoming

# Allow all outgoing traffic
sudo ufw default allow outgoing

# Allow SSH connections
sudo ufw allow ssh

# Allow web traffic
sudo ufw allow http
sudo ufw allow https

# Enable the firewall
sudo ufw enable

# Check the status to confirm the rules
sudo ufw status
```

#### 2. Enforce SSH Key-Based Authentication

Using SSH keys for authentication is significantly more secure than using passwords. Passwords can be cracked by brute-force attacks, whereas SSH keys are nearly impossible to decipher.

**Implementation Steps:**
1.  **Generate an SSH Key Pair:** Create a key pair on your local machine if you don't already have one.
2.  **Copy the Public Key to the Server:** Use the `ssh-copy-id` utility to securely install your public key on the server for the desired user.
    ```bash
    ssh-copy-id user@your_server_ip
    ```
3.  **Disable Password Authentication:** Once you have confirmed you can log in with your key, edit the SSH daemon configuration file (`/etc/ssh/sshd_config`) on the server and make the following changes:
    ```ini
    # Disallow password-based logins
    PasswordAuthentication no
    
    # Disallow logging in as the root user directly
    PermitRootLogin no
    ```
4.  **Restart the SSH Service:** Apply the changes by restarting the SSH daemon.
    ```bash
    sudo systemctl restart sshd
    ```

These two practices establish a strong security baseline for any public-facing server.