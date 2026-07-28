## Usage

Sources:

- [Official upstream README](https://github.com/gaoweichang/auto-job/blob/2daa1df1c64ac1fd57652b2ccc5e714d319af0ee/README.md)
- [Official extension control file (auto_job.control)](https://github.com/gaoweichang/auto-job/blob/2daa1df1c64ac1fd57652b2ccc5e714d319af0ee/auto_job.control)
- [Official extension SQL (auto_job--1.0.0.sql)](https://github.com/gaoweichang/auto-job/blob/2daa1df1c64ac1fd57652b2ccc5e714d319af0ee/sql/auto_job--1.0.0.sql)

`auto_job` — Auto job extension is the postgreSQL extension that schedules and runs stored procedures automatically using background workers. Use it for the corresponding scheduling, temporal, or time-series workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION auto_job;

SELECT pid, application_name, state, query
FROM pg_stat_activity
WHERE datname = 'postgres';
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_job(proc_name TEXT, schedule_interval INTERVAL)` is an extension function and returns `INTEGER`.
- `delete_job(job_id INTEGER)` is an extension function and returns `VOID`.
- `run_job(job_id INTEGER)` is an extension function and returns `VOID`.
- `_auto_job_catalog.job_info` is an extension-defined view.
- `_auto_job_catalog.jobs` is a table installed or managed by the extension.
- `public.auto_job_registry` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Install the confirmed extension dependencies first: `dblink`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
