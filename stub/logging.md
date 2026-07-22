## Usage

Sources:

- [Extension control file](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging.control)
- [Version 0.1 extension SQL](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging--0.1.sql)
- [C implementation](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging.c)

`logging` is a historical PostgreSQL experiment that switches a heap table and its indexes between permanent and unlogged catalog states. Its single function edits `pg_class.relpersistence` directly; it predates PostgreSQL's supported `ALTER TABLE SET LOGGED` and `ALTER TABLE SET UNLOGGED` workflow and should not be treated as a modern data-conversion tool.

### Core Workflow

The published interface is:

```sql
CREATE EXTENSION logging;

SELECT enable_logging('staging_events'::regclass::oid, false);
SELECT enable_logging('staging_events'::regclass::oid, true);
```

`false` requests the unlogged state and `true` requests the permanent, WAL-logged state. The function returns `true` when it changes the state and `false` when the relation is already in the requested state.

### Operational Boundary

- `enable_logging(relid oid, mode boolean) -> boolean` is the only SQL function.
- It requires superuser privilege and takes `AccessExclusiveLock` on the table.
- It changes the heap and each current index by updating system catalog tuples.
- It does not expose a preload requirement and version `0.1` is marked relocatable.

### Operational Notes

Do not use this extension to change production durability. Directly changing `relpersistence` bypasses the storage rewrite and safety checks provided by supported PostgreSQL DDL, and the C source uses server APIs removed from modern PostgreSQL releases. Prefer `ALTER TABLE staging_events SET UNLOGGED` or `ALTER TABLE staging_events SET LOGGED`, which also makes the durability transition explicit in schema management.

If preserving the extension for archaeology or an old server, test crash recovery, indexes, replicas, backups, and server-version compatibility on disposable data. An unlogged relation is truncated after a crash and is not replicated to standbys in the same way as a permanent relation.
