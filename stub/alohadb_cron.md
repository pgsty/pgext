## Usage

Sources:

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_cron/alohadb_cron--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cron/alohadb_cron--1.0.sql)
- [contrib/alohadb_cron/alohadb_cron.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cron/alohadb_cron.c)
- [contrib/alohadb_cron/alohadb_cron.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cron/alohadb_cron.control)

`alohadb_cron` 1.0 is AlohaDB’s SQL job scheduler. A background worker polls the configured database, executes due commands and records run details.

### Core Workflow

```ini
shared_preload_libraries = 'alohadb_cron'
alohadb.cron_database = 'postgres'
```

```sql
CREATE EXTENSION alohadb_cron;
SELECT cron_schedule_named('health-check', '*/5 * * * *', 'SELECT 1');
SELECT * FROM cron_job_status();
SELECT cron_unschedule_named('health-check');
```

### Operational Boundaries

Append the library to the existing preload list and restart, then create the extension in the configured database as a superuser. `alohadb.cron_database` defaults to postgres; `alohadb.cron_check_interval` controls polling. The worker expects its tables in public.

`cron_schedule` and `cron_schedule_named` create jobs; `cron_unschedule` and `cron_unschedule_named` remove them; `cron_job_status` summarizes state. Inspect `alohadb_cron_job_run_details` and define retention for growing history.

The reviewed worker connects with the default background-worker superuser identity and executes commands through SPI. Its per-job database and username arguments are not used to switch connection or role. Restrict job creation and table modification to trusted administrators; this implementation does not provide tenant isolation. This is an AlohaDB component, not a tested stock PostgreSQL package.
