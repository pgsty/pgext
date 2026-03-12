


## Usage

> [pgpool_recovery: recovery functions for pgpool-II](https://pgpool.net/)

The `pgpool_recovery` extension provides recovery-related functions used by Pgpool-II for online recovery of backend PostgreSQL nodes.

### Functions

```sql
-- Trigger online recovery of a backend node
-- Executes the recovery script on the primary node
SELECT pgpool_recovery(
    'recovery_1st_stage_script',   -- script name in $PGDATA
    'target_hostname',             -- hostname of node to recover
    'target_pgdata',               -- data directory on target
    'target_port'                  -- port number on target
);

-- Second stage recovery (optional, for streaming replication)
SELECT pgpool_remote_start(
    'target_hostname',             -- hostname of recovered node
    'target_pgdata'                -- data directory on target
);

-- Check if the target node is ready
SELECT pgpool_pgctl('status', 'target_pgdata');
```

### How It Works

1. Pgpool-II calls `pgpool_recovery()` on the primary node, which executes a user-defined shell script to perform base backup and setup
2. The recovery script copies data to the target node and configures replication
3. `pgpool_remote_start()` starts the recovered PostgreSQL instance
4. Pgpool-II attaches the recovered node back into the pool

### Recovery Scripts

The recovery script (e.g., `recovery_1st_stage`) must be placed in the primary node's `$PGDATA` directory. It typically performs:

- `pg_basebackup` to copy data to the target
- Configuration of `primary_conninfo` for streaming replication
- Creation of `standby.signal` on the target

The extension must be installed on all PostgreSQL backend nodes managed by Pgpool-II.
