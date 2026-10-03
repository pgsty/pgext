## Usage

Sources:

- [packages/jobs-simple/launchql-ext-jobs-queue.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs-simple/launchql-ext-jobs-queue.control)
- [packages/jobs-simple/sql/launchql-ext-jobs-queue--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs-simple/sql/launchql-ext-jobs-queue--0.4.5.sql)
- [packages/jobs-simple/readme.md](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs-simple/readme.md)

`launchql-ext-jobs-queue` 0.4.5 stores asynchronous jobs transactionally in PostgreSQL. An external worker performs the actual application task.

### Core Workflow

Install dependencies and prepare the `administrator` role before creating the extension. Claim a job as a uniquely named worker, process it, then acknowledge or fail the returned job ID.

```sql
CREATE EXTENSION "launchql-ext-jobs-queue" CASCADE;
SELECT app_jobs.add_job('send_email', '{"recipient":"reader@example.com"}'::json);
SELECT * FROM app_jobs.get_job('worker_1');
```

### Objects and Dependencies

`app_jobs.jobs` and `app_jobs.job_queues` store work and queue state. `app_jobs.complete_job(worker_id, job_id)` acknowledges a claimed job; `app_jobs.fail_job(worker_id, job_id, error_message)` records a failure. `app_jobs.add_scheduled_job`, `app_jobs.get_scheduled_job` and `app_jobs.run_scheduled_job` manage scheduled work. Dependencies are `plpgsql`, `pgcrypto`, `uuid-ossp` and `launchql-ext-default-roles`.

### Boundaries

The current README uses a different example namespace; the installed API is the versioned SQL cited above. Both LaunchQL job extensions create the same `app_jobs` objects and must not be installed together. No shared library or preload is needed. SQL function installation requires suitable rights in that schema; the control’s non-superuser setting does not grant missing roles or permissions. Workers must handle retries, expired claims and job retention. No PostgreSQL-major support claim is inferred from the SQL language.
