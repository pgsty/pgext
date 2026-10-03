## Usage

Sources:

- [packages/jobs/launchql-extension-jobs.control](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/jobs/launchql-extension-jobs.control)
- [packages/jobs/sql/launchql-extension-jobs--0.0.2.sql](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/jobs/sql/launchql-extension-jobs--0.0.2.sql)
- [packages/jobs/readme.md](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/jobs/readme.md)

`launchql-extension-jobs` 0.0.2 stores asynchronous jobs transactionally in PostgreSQL. An external worker performs the actual application task.

### Core Workflow

Install dependencies and prepare the `administrator` role before creating the extension. Claim a job as a uniquely named worker, process it, then acknowledge or fail the returned job ID.

```sql
CREATE EXTENSION "launchql-extension-jobs" CASCADE;
SELECT app_jobs.add_job('send_email', '{"recipient":"reader@example.com"}'::json);
SELECT * FROM app_jobs.get_job('worker_1');
```

### Objects and Dependencies

`app_jobs.jobs` and `app_jobs.job_queues` store work and queue state. `app_jobs.complete_job(worker_id, job_id)` acknowledges a claimed job; `app_jobs.fail_job(worker_id, job_id, error_message)` records a failure. `app_jobs.add_scheduled_job`, `app_jobs.get_scheduled_job` and `app_jobs.run_scheduled_job` manage scheduled work. Dependencies are `plpgsql`, `pgcrypto`, `uuid-ossp`.

### Boundaries

The installer grants privileges to an existing administrator role even though that role is not a declared extension dependency. Both LaunchQL job extensions create the same `app_jobs` objects and must not be installed together. No shared library or preload is needed. SQL function installation requires suitable rights in that schema; the control’s non-superuser setting does not grant missing roles or permissions. Workers must handle retries, expired claims and job retention. No PostgreSQL-major support claim is inferred from the SQL language.
