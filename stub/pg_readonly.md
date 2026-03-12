


## Usage

> [pg_readonly: cluster database read only](https://github.com/pierreforstmann/pg_readonly)

pg_readonly sets all databases in a PostgreSQL cluster to read-only mode at the SQL level. It must be loaded via `shared_preload_libraries`. The read-only status is managed in shared memory with a global flag (not persisted across restarts).

### Check Read-Only Status

```sql
SELECT get_cluster_readonly();
-- Returns false (read-write) or true (read-only)
```

### Set Cluster Read-Only

```sql
SELECT set_cluster_readonly();
```

In read-only mode, SELECT statements are allowed (unless they call writing functions), but DML (INSERT, UPDATE, DELETE), DDL (including TRUNCATE), and DCL (GRANT, REVOKE) are blocked:

```sql
SELECT * FROM t;              -- OK
UPDATE t SET x = 33;          -- ERROR: pg_readonly: invalid statement because cluster is read-only
CREATE TABLE tmp(c text);     -- ERROR: pg_readonly: invalid statement because cluster is read-only
```

Note: `set_cluster_readonly()` terminates all open transactions.

### Set Cluster Read-Write

```sql
SELECT unset_cluster_readonly();
```

Note: background processes (checkpointer, bgwriter, walwriter, autovacuum) continue running in read-only mode -- the restriction is at the SQL statement level only.
