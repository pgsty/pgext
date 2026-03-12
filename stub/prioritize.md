


## Usage

> [prioritize: get and set the priority of PostgreSQL backends](https://github.com/schmiddy/pg_prioritize)

The `prioritize` extension exposes `getpriority()` and `setpriority()` system calls for PostgreSQL backends, allowing you to `renice` backend processes from SQL.

### Get Backend Priority

```sql
SELECT get_backend_priority(pg_backend_pid());
```

Any user may query the priority of any backend.

### Set Backend Priority

```sql
SELECT set_backend_priority(pg_backend_pid(), 10);
```

Superusers can set the priority of any backend. Unprivileged users can only adjust backends with the same role.

Note: priority can only be raised (higher numeric value = lower OS priority). Only root can lower the numeric priority value, and PostgreSQL processes should not run as root.

### Batch Operations

```sql
-- Increase priority of all current user's backends by 5
SELECT set_backend_priority(pid, get_backend_priority(pid) + 5)
  FROM pg_stat_activity
  WHERE usename = CURRENT_USER;
```
