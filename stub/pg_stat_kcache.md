

## Usage

> [pg_stat_kcache: kernel cache statistics for PostgreSQL](https://github.com/powa-team/pg_stat_kcache)

pg_stat_kcache gathers statistics about real reads and writes done by the filesystem layer, as well as CPU usage per query. It requires `pg_stat_statements`.

### Views

**`pg_stat_kcache`** -- aggregated per-database statistics:

| Column | Type | Description |
|--------|------|-------------|
| `datname` | name | Database name |
| `exec_user_time` | double precision | User CPU time executing statements (seconds) |
| `exec_system_time` | double precision | System CPU time executing statements (seconds) |
| `exec_reads` | bigint | Bytes read by filesystem layer |
| `exec_reads_blks` | bigint | 8K blocks read by filesystem layer |
| `exec_writes` | bigint | Bytes written by filesystem layer |
| `exec_writes_blks` | bigint | 8K blocks written by filesystem layer |
| `plan_user_time` | double precision | User CPU time planning (if tracking enabled) |
| `plan_system_time` | double precision | System CPU time planning (if tracking enabled) |

**`pg_stat_kcache_detail`** -- per-query statistics including query text, role, and database.

### Functions

```sql
-- Reset all collected statistics
SELECT pg_stat_kcache_reset();
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_stat_kcache.linux_hz` | -1 | Linux CONFIG_HZ (auto-detected) |
| `pg_stat_kcache.track` | `top` | Track: `top`, `all`, or `none` |
| `pg_stat_kcache.track_planning` | `off` | Track planning statistics (PG 13+) |

### Example

```sql
SELECT datname, exec_user_time, exec_system_time, exec_reads, exec_writes
FROM pg_stat_kcache;

SELECT query, exec_user_time, exec_system_time, exec_reads
FROM pg_stat_kcache_detail
ORDER BY exec_user_time DESC
LIMIT 10;
```
