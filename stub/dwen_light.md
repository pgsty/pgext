## Usage

Sources:

- [README.md](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/README.md)
- [dwen_light.control](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/dwen_light.control)
- [dwen_light--0.2.sql](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/dwen_light--0.2.sql)
- [LICENSE](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/LICENSE)

`dwen_light` is an early, unstable background-worker task runner for SQL and PL/pgSQL. Version 0.2 is developed and tested with PostgreSQL 18 on Ubuntu. It installs in the fixed `dwen` schema and uses a configured database role to execute queued SQL.

### Configure and Enable

The official setup preloads the library, configures its worker pool and connection, restarts PostgreSQL, creates the extension, then restarts again:

```ini
shared_preload_libraries = 'dwen_light'
dwen.max_execution_workers = 2
dwen.database = 'postgres'
dwen.role = 'dwen'
```

Preserve other entries in the existing preload list. As a superuser, connect to the configured database:

```sql
CREATE EXTENSION dwen_light;
```

The installation script creates role `dwen` if absent and grants it privileges in the extension schema, including default privileges for future objects. Review the configured role and access to task tables before enabling workers; task strings are executable SQL.

### Submit and Observe

The upstream demonstration procedure is installed with the extension:

```sql
INSERT INTO dwen.task (status, proc, execute_string, run_timestamp)
VALUES ('READY', 'X', 'call dwen.dev_test_proc()', NULL);

SELECT event_at, type, message, task_id, errcode
FROM dwen.log
ORDER BY event_at DESC;
```

The documented polling interval starts this task within approximately 15 seconds. The task table also has a `timeout_minutes` column defaulting to 60. `dwen.task_execute_string()` executes the stored command and logs success or failure; `dwen.get_engine_status_active()` and `dwen.set_engine_status_active(boolean)` expose the engine's active flag.

### Boundaries

This is a preview framework with changing behavior, not a documented durable queue contract. The reviewed source supplies an installation script for 0.2 and no versioned upgrade path. It also installs demonstration objects. Keep task submission restricted to roles authorized to run work under the configured worker identity, and test failure and restart behavior before operational use.
