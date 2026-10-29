## Usage

Sources:
- [README](https://github.com/fraruiz/pgmb/blob/release/README.md)
- [Control](https://github.com/fraruiz/pgmb/blob/release/pgmb.control)
- [SQL](https://github.com/fraruiz/pgmb/blob/release/pgmb--1.0.0.sql)
- [META.json](https://github.com/fraruiz/pgmb/blob/release/META.json)

`pgmb` implements a SQL message broker with worker endpoints, routing-key patterns, queues and HTTP dispatch scheduled through `pg_cron`. The upstream repository is archived.

### Installation and compatibility

The RPM revision `1.0.0-2PGSTY` requires `pgsql_http_N` and removes a redundant schema-creation statement. Earlier RPMs and existing DEB packages retain the upstream SQL and can fail at `CREATE EXTENSION pgmb` because the control file and install script both create the schema. Package installation alone does not establish that the extension can be created.

Install as a superuser. The required extensions are `pg_cron` and `http`; `pgmb` itself does not require preload. Configure the scheduler for the database in which the extension will be used:

```ini
shared_preload_libraries = 'pg_cron'
cron.database_name = 'postgres'
```

Restart PostgreSQL, connect to that database, and create the dependencies before the extension:

```sql
CREATE EXTENSION pg_cron;
CREATE EXTENSION http;
CREATE EXTENSION pgmb;
```

### Register a queue and send a message

The worker endpoint receives HTTP requests from the database server. Use a trusted endpoint and restrict who can register workers.

```sql
SELECT pgmb.create(
    'orders', 'order.*', 3,
    pgmb.worker('order_worker', 'http://127.0.0.1:8080/worker', 100)
);

SELECT pgmb.send(
    gen_random_uuid(), 'order.created', '{"order_id": 123}'::jsonb
);

SELECT * FROM pgmb.workers;
SELECT * FROM pgmb.queues;
SELECT * FROM pgmb.orders_queue;
SELECT * FROM pgmb.orders_dead_letter_queue;
```

The argument order for `pgmb.create` is name, binding pattern, maximum retries, then worker UUID. A trigger routes messages to matching queues, and queue creation registers a scheduler job.

### API

| Function | Purpose |
|---|---|
| `pgmb.worker(name, endpoint, rps)` | Register a worker and return its UUID; rps limits the batch size per dispatch. |
| `pgmb.create(name, binding_key, max_retries, worker_id)` | Create a queue, its dead-letter table and scheduler job; return the queue UUID. |
| `pgmb.send(id, routing_key, body)` | Store a JSONB message and route it to matching queues. |
| `pgmb.send` overloads | Accept optional headers and an integer-seconds or timestamp delay. |
| `pgmb.dispatch_messages(queue_id)` | Dispatch a batch; mark successful HTTP responses acknowledged, retry failures and move exhausted messages to the dead-letter table. |

### Limits of version 1.0.0

- Delays populate `enqueued_at`, but the dispatcher does not filter out future timestamps; the overloads do not guarantee delayed delivery.
- Headers are stored with messages, but the dispatcher sends only the body using `http_post`; they are not forwarded as custom HTTP headers.
- Retries use a counter on subsequent dispatches. The SQL does not implement exponential backoff.
- HTTP effects are external to the database transaction. Review delivery and retention requirements before relying on this archived project.
