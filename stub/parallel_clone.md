## Usage

Sources:

- [Official Clone Schema setup](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/setting_up_edb_clone_schema/)
- [Official Clone Schema documentation](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/)
- [Official EDB extension catalog](https://www.enterprisedb.com/docs/pg_extensions/)

`parallel_clone` is the EPAS worker extension used by `edb_cloneschema` to copy schema objects with parallel background processing.

### Enablement

Install the matching EPAS package, preload `parallel_clone` with `edb_job_scheduler`, restart, then create it in every source or target database:

```ini
shared_preload_libraries = 'parallel_clone,edb_job_scheduler'
```

```sql
CREATE EXTENSION parallel_clone;
```

The provider-only module is supported on EDB Postgres Advanced Server and is normally installed as part of the Clone Schema workflow.

### Relationship to Clone Schema

Create `edb_cloneschema` only after `parallel_clone`, `edb_job_scheduler`, `dbms_job`, `postgres_fdw`, `dblink`, and PL/Perl prerequisites are present.

```sql
CREATE EXTENSION edb_cloneschema;
SELECT edb_util.create_clone_log_dir();
```

### Operational Boundaries

Parallel clone workers consume `max_worker_processes`, memory, WAL, locks, checkpoints, and log space. Large schema copies can run in one transaction and may require deliberate tuning. Treat `parallel_clone` as an internal runtime dependency: use documented `edb_cloneschema` APIs, monitor status logs, and do not call undocumented worker functions directly.

