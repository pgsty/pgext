


## Usage

> [qos: QoS resource governor extension for PostgreSQL sessions and queries](https://github.com/appstonia/pg_qos)

The `qos` extension provides Quality of Service resource governance for PostgreSQL, allowing administrators to set per-role and per-database limits on memory usage, CPU cores, and concurrent transactions/statements.

### Configuration Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `qos.work_mem_limit` | bytes | Maximum effective `work_mem` per session |
| `qos.cpu_core_limit` | integer | Maximum CPU cores available to a session |
| `qos.max_concurrent_tx` | integer | Maximum concurrent transactions |
| `qos.max_concurrent_select` | integer | Maximum concurrent SELECT statements |
| `qos.max_concurrent_update` | integer | Maximum concurrent UPDATE statements |
| `qos.max_concurrent_delete` | integer | Maximum concurrent DELETE statements |
| `qos.max_concurrent_insert` | integer | Maximum concurrent INSERT statements |

### Per-Role Limits

```sql
ALTER ROLE app_user SET qos.work_mem_limit = '32MB';
ALTER ROLE app_user SET qos.cpu_core_limit = '2';
ALTER ROLE app_user SET qos.max_concurrent_select = '100';
```

### Per-Database Limits

```sql
ALTER DATABASE appdb SET qos.max_concurrent_tx = '200';
```

### Combined Role + Database Limits

```sql
ALTER ROLE app_user IN DATABASE appdb SET qos.work_mem_limit = '4MB';
ALTER ROLE app_user IN DATABASE appdb SET qos.max_concurrent_update = '10';
```

### Enforcement Behavior

- **Work memory**: Intercepts `SET work_mem` and rejects values exceeding configured limits
- **CPU limiting** (Linux only): Binds backend to N CPU cores via CPU affinity; on non-Linux platforms, limits parallel workers instead
- **Concurrency**: Executor hooks track active transactions/statements by type; violations block execution

### Observability

```sql
SET client_min_messages = 'debug1';  -- enable debug output for QoS events
```

The most restrictive combination of role-level and database-level settings takes effect. Requires `shared_preload_libraries = 'qos'` and PostgreSQL 15+.
