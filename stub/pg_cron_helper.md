## Usage

Sources:

- [Official upstream README](https://github.com/splendiddata/pg_cron_helper/blob/5bc4ea422971babfc4a1e8ba5a8f5b2f20e8c62a/README.md)
- [Official extension control file (pg_cron_helper.control)](https://github.com/splendiddata/pg_cron_helper/blob/5bc4ea422971babfc4a1e8ba5a8f5b2f20e8c62a/pg_cron_helper.control)
- [Official extension SQL (pg_cron_helper--0.1.sql)](https://github.com/splendiddata/pg_cron_helper/blob/5bc4ea422971babfc4a1e8ba5a8f5b2f20e8c62a/pg_cron_helper--0.1.sql)

`pg_cron_helper` — The pg_cron_helper database extension is supposed to help running "jobs" on a simple time-based schedule. Of course you are far better of using an external scheduler, but if you must schedule jobs inside a Postgres database then this extension may be of some help. Use it for the corresponding scheduling, temporal, or time-series workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_cron_helper;

select * from  cron.list_jobs();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_job_state(job_name varchar(128) , user_name name default current_user)` is an extension function and returns `text`.
- `list_jobs()` is an extension function and returns `setof`.
- `create_job` is an extension procedure.
- `cron.stop_job` is an extension procedure.
- `disable_job` is an extension procedure.
- `drop_job` is an extension procedure.
- `enable_job` is an extension procedure.
- `run_job` is an extension procedure.
- `stop_job` is an extension procedure.
- `job_record` is an extension-defined type.
- `job_definition` is a table installed or managed by the extension.
- `job_run` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.3`.
- Install the confirmed extension dependencies first: `postgres_fdw`, `dblink`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
