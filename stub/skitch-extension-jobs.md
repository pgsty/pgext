## Usage

Sources:

- [Official upstream README](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/jobs/readme.md)
- [Official extension control file (skitch-extension-jobs.control)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/jobs/skitch-extension-jobs.control)
- [Official extension SQL (skitch-extension-jobs--0.0.7.sql)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/jobs/sql/skitch-extension-jobs--0.0.7.sql)

`skitch-extension-jobs` — An asynchronous job queue schema for ACID compliant job creation. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "skitch-extension-jobs";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `app_jobs.add_job(identifier varchar, payload json)` is an extension function and returns `app_jobs`.
- `app_jobs.add_job(identifier varchar, queue_name varchar, payload json)` is an extension function and returns `app_jobs`.
- `app_jobs.complete_job(worker_id varchar, job_id int)` is an extension function and returns `app_jobs`.
- `app_jobs.do_notify()` is an extension function and returns `trigger`.
- `app_jobs.fail_job(worker_id varchar, job_id int, error_message varchar)` is an extension function and returns `app_jobs`.
- `app_jobs.get_job(worker_id varchar, identifiers varchar[])` is an extension function and returns `app_jobs`.
- `app_jobs.schedule_job(identifier varchar, queue_name varchar, payload json, run_at timestamptz)` is an extension function and returns `app_jobs`.
- `app_jobs.tg__add_job_for_row()` is an extension function and returns `trigger`.
- `app_jobs.tg_decrease_job_queue_count()` is an extension function and returns `trigger`.
- `app_jobs.tg_increase_job_queue_count()` is an extension function and returns `trigger`.
- `app_jobs.job_queues` is a table installed or managed by the extension.
- `app_jobs.jobs` is a table installed or managed by the extension.
- `app_jobs` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.7`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgcrypto`, `uuid-ossp`, `skitch-extension-defaults`, `skitch-extension-verify`, `skitch-extension-utils`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
