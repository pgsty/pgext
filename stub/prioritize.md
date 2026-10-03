## Usage

Sources:

- [Official README](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/README.md)
- [SQL functions](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/prioritize.sql.in)
- [Control file](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/prioritize.control)

`prioritize` exposes operating-system priority controls for PostgreSQL backend processes. Use it to lower the scheduling priority of selected sessions; it is not a PostgreSQL query scheduler.

### Inspect and Adjust a Backend

```sql
CREATE EXTENSION prioritize;
SELECT get_backend_priority(pg_backend_pid());
SELECT set_backend_priority(pg_backend_pid(), 10);
```

Any user may query a backend's priority. A PostgreSQL superuser may request adjustments for any backend; other users may adjust only backends running under the same database role.

### Adjust Related Sessions

```sql
SELECT set_backend_priority(pid, get_backend_priority(pid) + 5)
FROM pg_stat_activity
WHERE usename = CURRENT_USER;
```

Increasing the numeric nice value lowers operating-system scheduling priority. The example therefore reduces those backends' priority, rather than speeding them up.

### Privileges and Limits

Installation requires a superuser, with no preload or restart. Operating-system permissions still apply: ordinary PostgreSQL processes normally cannot lower the numeric nice value to increase priority. Database superuser privileges do not confer root privileges. Review platform scheduling policy and process identity before using bulk adjustments. The control file uses SQL version 1.0; distribution package version 1.0.4 is separate.
