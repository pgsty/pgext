## Usage

Sources:

- [Official upstream README](https://github.com/victornagibator/pg_scheduler/blob/48f3bdb0119a457d7af410057a349bd023d7b813/README.md)
- [Official extension control file (scheduler.control)](https://github.com/victornagibator/pg_scheduler/blob/48f3bdb0119a457d7af410057a349bd023d7b813/scheduler.control)
- [Official extension SQL (scheduler--1.0.sql)](https://github.com/victornagibator/pg_scheduler/blob/48f3bdb0119a457d7af410057a349bd023d7b813/scheduler--1.0.sql)

`scheduler` — pg_scheduler is a PostgreSQL extension that enables flexible scheduling and execution of SQL and shell jobs directly within your database. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION scheduler;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `scheduler.add_or_update_job(p_name TEXT, p_type TEXT, p_cmd TEXT, p_interval INTERVAL = NULL, p_time TIMESTAMPTZ = NULL, p_max_attempts INT = 3)` is an extension function and returns `TEXT`.
- `scheduler.delete_job(p_name TEXT)` is an extension function and returns `VOID`.
- `scheduler.execute_job(target_job_id INT)` is an extension function and returns `VOID`.
- `scheduler.execute_shell_command(cmd TEXT)` is an extension function and returns `VOID`.
- `scheduler.set_next_run()` is an extension function and returns `TRIGGER`.
- `scheduler.toggle_job(p_name TEXT, p_enabled BOOLEAN)` is an extension function and returns `VOID`.
- `scheduler.update_timestamp()` is an extension function and returns `TRIGGER`.
- `scheduler.job_logs` is a table installed or managed by the extension.
- `scheduler.jobs` is a table installed or managed by the extension.
- `scheduler` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
