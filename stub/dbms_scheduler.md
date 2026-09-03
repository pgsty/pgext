## Usage

Sources:

- [Official DBMS_SCHEDULER reference](https://www.enterprisedb.com/docs/epas/latest/reference/oracle_compatibility_reference/epas_compat_bip_guide/03_built-in_packages/15_dbms_scheduler/)
- [Official EDB Job Scheduler documentation](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/)
- [Official scheduler configuration](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/configuring/)

`dbms_scheduler` supplies the EPAS Oracle-compatible `DBMS_SCHEDULER` package for named programs, schedules, arguments, and jobs.

### Enablement

Configure the EDB Job Scheduler worker, restart, then create both extensions:

```ini
edb_job_scheduler.database_list = 'appdb'
shared_preload_libraries = 'edb_job_scheduler'
```

```sql
CREATE EXTENSION edb_job_scheduler;
CREATE EXTENSION dbms_scheduler;
```

This provider-only extension is available on EDB Postgres Advanced Server.

### Define and Run Jobs

`DBMS_SCHEDULER.CREATE_PROGRAM` defines an executable program and its arguments. `DBMS_SCHEDULER.CREATE_SCHEDULE` defines timing, and `DBMS_SCHEDULER.CREATE_JOB` binds them or directly names an action.

```sql
BEGIN
  DBMS_SCHEDULER.CREATE_JOB(
    job_name        => 'hourly_rollup',
    job_type        => 'STORED_PROCEDURE',
    job_action      => 'refresh_rollup',
    start_date      => SYSDATE,
    repeat_interval => 'FREQ=HOURLY',
    enabled         => TRUE
  );
END;
```

Use `RUN_JOB`, `ENABLE`, `DISABLE`, `DROP_JOB`, `DROP_PROGRAM`, and `DROP_SCHEDULE` for lifecycle operations.

### Compatibility Boundary

EDB implements a documented subset of Oracle `DBMS_SCHEDULER`; do not assume unsupported Oracle attributes or semantics. Job actions run with database privileges and can leave partial effects if the procedure is not transactional or idempotent. Monitor scheduler tables and test calendar strings with `EVALUATE_CALENDAR_STRING` before enabling recurring jobs.

