## Usage

Sources:

- [packages/jobs/readme.md](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/readme.md)
- [packages/jobs/sql/launchql-ext-jobs--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/sql/launchql-ext-jobs--0.4.5.sql)
- [packages/jobs/Makefile](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/Makefile)
- [packages/jobs/launchql-ext-jobs.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/launchql-ext-jobs.control)

`launchql-ext-jobs` stores jobs, scheduled jobs and queue locks in `app_jobs`. External workers acquire and execute the tasks.

### Core Workflow

```sql
CREATE EXTENSION "launchql-ext-jobs" CASCADE;
SELECT * FROM app_jobs.jobs LIMIT 10;
SELECT * FROM app_jobs.scheduled_jobs LIMIT 10;
```

### Operational Boundaries

`app_jobs.add_job` takes a database UUID, task identifier and JSON payload. `app_jobs.get_job`, `app_jobs.complete_job`, `app_jobs.fail_job` and `app_jobs.release_jobs` coordinate ownership and retries. Make retryable jobs idempotent. This historical schema conflicts with `pgpm-jobs` and `pgpm-database-jobs`; do not co-install them. The SQL expects the matching default roles and JWT-claim integration.

Install its declared dependencies first: `plpgsql`, `uuid-ossp`, `pgcrypto`, `launchql-ext-default-roles`, `launchql-jwt-claims`. This is SQL/PLpgSQL code with no own shared library or preload. The control permits non-superuser installation but is not marked trusted; dependency and schema privileges still apply.
