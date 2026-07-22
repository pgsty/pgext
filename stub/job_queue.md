## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/README.md)
- [Version 1.0 SQL objects](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/job_queue--1.0.sql)
- [Background-worker implementation](https://github.com/michelmilezzi/pg_job_queue/blob/b742a0713ebe46b15a154a53074a52107338a620/job_queue.c)

`job_queue` is an archived proof of concept that starts a dynamic background worker after a statement inserts into `jobs_queue`. The worker claims one row with `FOR UPDATE SKIP LOCKED`, deletes it, and executes the stored `proc` text with its optional JSON `args`.

```sql
CREATE EXTENSION job_queue;
INSERT INTO jobs_queue (proc, args)
VALUES ('SELECT process_order($1)', '{"order_id": 42}');
```

This is not a safe production queue. `proc` is concatenated into SQL rather than resolved as a typed procedure, and arbitrary writers can therefore execute SQL with the worker's database privileges. Revoke table and launch-function access from `PUBLIC` and do not expose either field to untrusted input.

The dequeue, deletion, and job body run in one worker transaction, but the implementation has an explicit unfinished error path: an exception from the job is not caught to update `error_count` or `last_error`. The source transaction only registers the worker; it does not wait for job success, and worker registration can fail when dynamic worker slots are exhausted.

Use a maintained queue with explicit retry, dead-letter, observability, cancellation, idempotency, and privilege semantics. Keep this extension to disposable PostgreSQL-background-worker experiments.
