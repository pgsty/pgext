## Usage

Sources:

- [external/pg_cron_preload/pg_cron_preload--1.0.sql](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/pg_cron_preload--1.0.sql)
- [external/pg_cron_preload/invalid_cache.c](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/invalid_cache.c)
- [external/pg_cron_preload/Makefile](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/Makefile)
- [external/pg_cron_preload/pg_cron_preload.control](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/fd9fc37f6a4d6fce58b61af39c18cb979fa15419/external/pg_cron_preload/pg_cron_preload.control)

`pg_cron_preload` is PolarDB’s separately packaged cron catalog component. It creates cron job/run tables and a C trigger that invalidates cached job metadata; it does not itself implement the scheduler.

### Core Workflow

```sql
CREATE EXTENSION pg_cron_preload;
SELECT jobid, schedule, database, username, active FROM cron.job;
SELECT jobid, status, start_time, end_time FROM cron.job_run_details;
```

### Operational Boundaries

Its SQL installs `cron.job`, `cron.job_run_details`, sequences, row-security policies and `cron.job_cache_invalidate`. The scheduling and unscheduling function declarations are commented out in this component. Use the matching PolarDB scheduler integration, not an assumed stock `pg_cron` API.

Create it only in the intended PolarDB installation as a superuser. The control fixes an extension schema, while SQL creates `cron`; preexisting cron objects can conflict. The C module contains no preload initialization hook, despite the name. This catalog entry makes no vanilla PostgreSQL compatibility or independent scheduler claim.
