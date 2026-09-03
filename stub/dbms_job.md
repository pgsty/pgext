## Usage

Sources:

- [Official DBMS_JOB reference](https://www.enterprisedb.com/docs/epas/latest/reference/oracle_compatibility_reference/epas_compat_bip_guide/03_built-in_packages/05_dbms_job/)
- [Official EDB Job Scheduler documentation](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/)
- [Official Clone Schema setup](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/setting_up_edb_clone_schema/)

`dbms_job` provides the EPAS Oracle-compatible `DBMS_JOB` package for creating, changing, running, and removing stored procedure jobs.

### Enablement

`dbms_job` relies on the preloaded EDB Job Scheduler. Configure and create the scheduler first, then create the package extension:

```ini
edb_job_scheduler.database_list = 'appdb'
shared_preload_libraries = 'edb_job_scheduler'
```

```sql
CREATE EXTENSION edb_job_scheduler;
CREATE EXTENSION dbms_job;
```

This provider-only extension is available on EDB Postgres Advanced Server, not community PostgreSQL.

### Create and Manage a Job

```sql
DECLARE
  job_id bigint;
BEGIN
  DBMS_JOB.SUBMIT(
    job_id,
    'job_proc;',
    SYSDATE,
    'SYSDATE + 1/24'
  );
  COMMIT;
END;
```

Use `DBMS_JOB.CHANGE`, `DBMS_JOB.INTERVAL`, `DBMS_JOB.NEXT_DATE`, `DBMS_JOB.RUN`, `DBMS_JOB.BROKEN`, and `DBMS_JOB.REMOVE` for lifecycle management.

### Operational Boundaries

`what` is stored procedure text executed later by a worker. Restrict package privileges, avoid constructing it from untrusted input, and make recurring work idempotent. `next_date` and `interval` interact: the interval expression is reevaluated before each run and becomes the next scheduled time. Monitor `sys.job_run_details` for failures rather than assuming submission implies completion.

