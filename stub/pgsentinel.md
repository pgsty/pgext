## Usage

Sources:

- [pgsentinel 1.5.1 README](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/README.md)
- [pgsentinel 1.5.1 release](https://github.com/pgsentinel/pgsentinel/releases/tag/v1.5.1)
- [1.5.1 upgrade SQL](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/src/pgsentinel--1.5.0--1.5.1.sql)
- [pgsentinel control file](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/src/pgsentinel.control)

`pgsentinel` records active session history by sampling `pg_stat_activity` at regular intervals and linking activity with `pg_stat_statements` query statistics. It keeps recent samples in shared-memory ring buffers managed by a background worker.

```ini
shared_preload_libraries = 'pg_stat_statements,pgsentinel'
pg_stat_statements.track = all
pgsentinel.db_name = 'postgres'
```

Restart PostgreSQL, then enable both extensions in the database used by the worker:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pgsentinel;
```

### Active Session History

```sql
SELECT ash_time, datname, usename, pid, state,
       wait_event_type, wait_event, query, queryid
FROM pg_active_session_history
ORDER BY ash_time DESC;
```

Key columns beyond `pg_stat_activity`:

| Column | Description |
|--------|-------------|
| `ash_time` | Sampling timestamp |
| `top_level_query` | Top-level statement (for PL/pgSQL) |
| `query` | Statement with actual parameter values |
| `cmdtype` | Statement type: SELECT, UPDATE, INSERT, DELETE, UTILITY, UNKNOWN, NOTHING |
| `queryid` | Links to `pg_stat_statements` |
| `blockers` | Number of blocking processes |
| `blockerpid` | PID of a blocking process |
| `blocker_state` | State of the blocker |

### Query Statistics History

When enabled, pgsentinel also samples `pg_stat_statements` concurrently:

```sql
SELECT ash_time, queryid, calls, total_exec_time, rows,
       shared_blks_hit, shared_blks_read
FROM pg_stat_statements_history
ORDER BY ash_time DESC;
```

### Example: Wait Analysis

```sql
-- Top wait events in the last hour
SELECT wait_event_type, wait_event, count(*)
FROM pg_active_session_history
WHERE ash_time > now() - interval '1 hour'
  AND wait_event IS NOT NULL
GROUP BY 1, 2
ORDER BY 3 DESC;

-- Blocking analysis
SELECT blockerpid, blocker_state, count(*)
FROM pg_active_session_history
WHERE blockers > 0
GROUP BY 1, 2
ORDER BY 3 DESC;
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgsentinel_ash.sampling_period` | 1 | Sampling period in seconds |
| `pgsentinel_ash.max_entries` | 1000 | Ring buffer size for ASH |
| `pgsentinel.db_name` | `postgres` | Database for worker connection |
| `pgsentinel_ash.track_idle_trans` | `false` | Track idle-in-transaction sessions |
| `pgsentinel_pgssh.max_entries` | 1000 | Ring buffer for pg_stat_statements history |
| `pgsentinel_pgssh.enable` | `false` | Enable pg_stat_statements history |

### Version and Privilege Changes

Version 1.5.0 separates `queryid` from `nested_queryid`: on PostgreSQL 16 and later the former comes from the activity query identifier, while the latter identifies the inner parsed statement; they match on older versions. The upgrade recreates the history view, so inspect dependent views and grants first.

Version 1.5.1 revokes PUBLIC execution of `get_parsedinfo(int)` because it can expose other backends' query text and literals. Apply the SQL upgrade after installing the new files, and grant access only to roles that need it:

```sql
ALTER EXTENSION pgsentinel UPDATE TO '1.5.1';
```

History remains in finite shared-memory ring buffers; export it for long-term retention. Sampled query text can contain sensitive values, so review access to history views as well. Upstream CI covers PostgreSQL 10–19.
