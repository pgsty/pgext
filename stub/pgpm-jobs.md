## Usage

Sources:

- [packages/jobs/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/README.md)
- [packages/jobs/sql/pgpm-jobs--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/sql/pgpm-jobs--0.47.0.sql)
- [packages/jobs/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/Makefile)
- [packages/jobs/pgpm-jobs.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/pgpm-jobs.control)

`pgpm-jobs` provides the unattributed `app_jobs` queue variant for standalone installations. SQL stores jobs, schedules, locks and retries; an external worker must execute tasks.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-jobs" CASCADE;
SELECT app_jobs.add_job('demo', '{}'::json);
SELECT * FROM app_jobs.jobs LIMIT 10;
```

### Operational Boundaries

`app_jobs.get_job`, `app_jobs.complete_job` and failure/release helpers coordinate workers; LISTEN/NOTIFY wakes consumers but is not itself a durable task executor. This variant has no tenant/billing identity propagation. Use `pgpm-database-jobs` for the attributed platform variant; both create the same schema and must not be co-installed.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `plpgsql`, `pgcrypto`, `pgpm-verify`, `errors`. Its SQL expects the platform roles `administrator`, `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
