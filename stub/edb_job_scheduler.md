## Usage

Sources:

- [Official EDB Job Scheduler documentation](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/configuring/)
- [Official user API](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/user_api/)

`edb_job_scheduler` runs stored database jobs in background workers for EDB Postgres Advanced Server and supplies the runtime used by `dbms_job` and `dbms_scheduler`.

### Enablement

List every scheduled database, preload the worker, restart, and create the extension in each listed database:

```ini
edb_job_scheduler.database_list = 'appdb'
shared_preload_libraries = 'edb_job_scheduler'
```

```sql
CREATE EXTENSION edb_job_scheduler;
```

The extension creates `sys.jobs` and `sys.job_run_details`. Configure `edb_job_scheduler.max_jobs_per_database`, `edb_job_scheduler.max_workers_per_database`, and the database-list capacity before production use.

### Submit and Inspect Jobs

Applications normally use `DBMS_JOB` or `DBMS_SCHEDULER`; the scheduler also exposes its `sys` tables for status inspection.

```sql
SELECT jobid, jobnextrun, jobcommand, jobuser
FROM sys.jobs
ORDER BY jobnextrun;

SELECT jobid, runid, status, starttime, endtime, error
FROM sys.job_run_details
ORDER BY runid DESC;
```

### Operational Boundaries

Jobs execute SQL as the configured job user and can change database state. Restrict package execution, validate stored commands, bound workers, monitor failed runs, and design idempotent jobs. Changes to `database_list` require reload or restart depending on EPAS version and configured capacity; databases omitted from the list do not receive scheduler workers.

