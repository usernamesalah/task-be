# Task 2: Database Maintenance & Backup 💾

### PostgreSQL Backup

To perform a manual backup of a PostgreSQL database running inside a Docker container, you can execute the `pg_dump` utility from the host machine using `docker exec`. This command streams the database dump to a local file.

A simple command to achieve this is:

```bash
docker exec <postgresql_container_name> pg_dumpall -U <user> > backup.sql
```

For a more robust solution, a shell script can be used to automate the process, including timestamping the backup file.

#### Example Backup Script for PostgreSQL

```bash
#!/bin/bash
# conceptual_postgres_backup.sh

# --- Configuration ---
# The name of the running PostgreSQL Docker container.
CONTAINER_NAME="your_postgres_container"
# The PostgreSQL user for the backup operation.
DB_USER="your_postgres_user"
# The local directory where backups will be stored.
BACKUP_DIR="/path/to/backups"
# --- End Configuration ---

# Generate a timestamp for the backup file.
DATE=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="$BACKUP_DIR/postgres_backup_$DATE.sql"

# Ensure the backup directory exists.
mkdir -p "$BACKUP_DIR"

# Execute pg_dumpall within the container and redirect the output to the backup file.
echo "Starting PostgreSQL backup for container '$CONTAINER_NAME'..."
docker exec "$CONTAINER_NAME" pg_dumpall -U "$DB_USER" > "$BACKUP_FILE"

# Optional: Compress the backup file to save space.
echo "Compressing backup file..."
gzip "$BACKUP_FILE"

echo "PostgreSQL backup successfully created at: $BACKUP_FILE.gz"
```

### MongoDB Backup

Similarly, a MongoDB database running in a Docker container can be backed up using `mongodump`. The process involves executing the command inside the container, which writes the backup to a directory within the container, and then copying that directory back to the host machine.

#### Example Backup Script for MongoDB

```bash
#!/bin/bash
# conceptual_mongodb_backup.sh

# --- Configuration ---
# The name of the running MongoDB Docker container.
CONTAINER_NAME="your_mongo_container"
# The name of the database to back up.
DB_NAME="your_mongo_db"
# The local directory on the host where the final backup archive will be stored.
BACKUP_DIR="/path/to/backups/mongo"
# --- End Configuration ---

# Generate a timestamp for the backup.
DATE=$(date +"%Y-%m-%d_%H-%M-%S")
# Define the temporary backup path inside the container.
BACKUP_PATH_INSIDE_CONTAINER="/backup/$DATE"
# Define the path on the host where the backup will be copied.
BACKUP_PATH_ON_HOST="$BACKUP_DIR/$DATE"

# Ensure the backup directory on the host exists.
mkdir -p "$BACKUP_PATH_ON_HOST"

# Execute mongodump inside the container.
echo "Starting MongoDB backup for database '$DB_NAME'..."
docker exec "$CONTAINER_NAME" mongodump --db "$DB_NAME" --out "$BACKUP_PATH_INSIDE_CONTAINER"

# Copy the backup files from the container to the host.
echo "Copying backup from container to host..."
docker cp "$CONTAINER_NAME:$BACKUP_PATH_INSIDE_CONTAINER" "$BACKUP_PATH_ON_HOST"

# Create a compressed tarball of the backup directory for easy management.
echo "Creating compressed archive..."
tar -czf "$BACKUP_PATH_ON_HOST.tar.gz" -C "$BACKUP_PATH_ON_HOST" .

# Clean up by removing the raw backup directory from the host.
rm -rf "$BACKUP_PATH_ON_HOST"

# Clean up by removing the temporary backup directory from the container.
docker exec "$CONTAINER_NAME" rm -rf "$BACKUP_PATH_INSIDE_CONTAINER"

echo "MongoDB backup successfully created at: $BACKUP_PATH_ON_HOST.tar.gz"
```

### Secure Backup Storage

It is highly recommended to store database backups in a secure, off-site location. Storing them on the same server as the database exposes them to the same risks, such as hardware failure or security breaches.

**Recommended Solutions:**
*   **Cloud Storage Services (e.g., Amazon S3, Google Cloud Storage, Azure Blob Storage):** These services are ideal because they offer high durability, availability, and scalability. They also provide robust security features, including encryption at rest and in transit, and fine-grained access control (IAM) to ensure only authorized personnel can access the backups.
*   **Dedicated Backup Servers:** A separate, physically isolated server can also be used, but requires more manual setup for security and redundancy compared to cloud options.

**Key Rationale:**
*   **Disaster Recovery:** Off-site storage is fundamental for recovering from a server-wide failure or a data center outage.
*   **Security:** Separating backups from the production environment limits the impact of a security compromise.
*   **Durability:** Cloud providers replicate data across multiple locations, offering a level of data durability that is difficult to achieve with a single on-premises solution.