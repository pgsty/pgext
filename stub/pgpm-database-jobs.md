## Usage

Sources:

- [packages/database-jobs/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/README.md)
- [packages/database-jobs/sql/pgpm-database-jobs--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/sql/pgpm-database-jobs--0.47.0.sql)
- [packages/database-jobs/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/Makefile)
- [packages/database-jobs/pgpm-database-jobs.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/pgpm-database-jobs.control)

`pgpm-database-jobs` is the attributed `app_jobs` implementation used by the Constructive platform. It stores task, schedule and worker state with database and identity context.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-database-jobs" CASCADE;
SELECT * FROM app_jobs.jobs LIMIT 10;
SELECT * FROM app_jobs.scheduled_jobs LIMIT 10;
```

### Operational Boundaries

Use the version-specific enqueue signatures to preserve database and actor attribution. `app_jobs.get_job`, `app_jobs.complete_job` and failure/release helpers manage worker ownership and retries. An external worker executes the payload; retryable tasks must be idempotent. Do not co-install the standalone `pgpm-jobs` or legacy `launchql-ext-jobs` variants, which also own `app_jobs`.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `plpgsql`, `pgcrypto`, `pgpm-verify`, `pgpm-jwt-claims`, `errors`. Its SQL expects the platform roles `administrator`, `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
