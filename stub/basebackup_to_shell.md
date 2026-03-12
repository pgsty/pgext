


## Usage

> [basebackup_to_shell: adds a custom basebackup target called shell](https://www.postgresql.org/docs/current/basebackup-to-shell.html)

The `basebackup_to_shell` module adds a custom `shell` backup target to `pg_basebackup`, allowing administrators to pipe backup archives through arbitrary shell commands.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'basebackup_to_shell'

# Command to execute for each archive; receives archive data on stdin
basebackup_to_shell.command = 'gzip > /backup/%f.gz'

# Optional: restrict usage to a specific role
basebackup_to_shell.required_role = 'backup_admin'
```

### Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `basebackup_to_shell.command` | string | Shell command to execute; receives archive on stdin |
| `basebackup_to_shell.required_role` | string | Role required to use the shell target (empty = any replication user) |

### Command Placeholders

| Placeholder | Replaced With |
|-------------|---------------|
| `%f` | Archive filename (e.g., `base.tar`) |
| `%d` | User-provided target detail string |
| `%%` | Literal `%` |

### Examples

```bash
# Compress backups to a local directory
# postgresql.conf: basebackup_to_shell.command = 'gzip > /backup/%f.gz'
pg_basebackup --target=shell

# Upload to S3 with a detail parameter
# postgresql.conf: basebackup_to_shell.command = 'aws s3 cp - s3://bucket/%d/%f'
pg_basebackup --target=shell:myprefix

# Custom processing pipeline
# postgresql.conf: basebackup_to_shell.command = 'zstd | ssh backup-host "cat > /backup/%f.zst"'
pg_basebackup --target=shell
```

The `%d` placeholder requires a target detail to be provided via `--target=shell:DETAIL`. If `%d` is not in the command, providing a detail is prohibited. Detail strings may only contain alphanumeric characters.
