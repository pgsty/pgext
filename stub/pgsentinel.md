

## Usage

> [pgsentinel: active session history for PostgreSQL](https://github.com/pgsentinel/pgsentinel)

pgsentinel records active session history by sampling `pg_stat_activity` at regular intervals and linking activity with `pg_stat_statements` query statistics.

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
