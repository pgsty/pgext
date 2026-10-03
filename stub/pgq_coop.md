## Usage

Sources:

- [README.rst](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/README.rst)
- [pgq_coop.control](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/pgq_coop.control)
- [Makefile](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/Makefile)
- [mk/common-pgxs.mk](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/mk/common-pgxs.mk)
- [sql/pgq_coop_test.sql](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/sql/pgq_coop_test.sql)
- [functions/pgq_coop.next_batch.sql](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/functions/pgq_coop.next_batch.sql)
- [functions/pgq_coop.finish_batch.sql](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/functions/pgq_coop.finish_batch.sql)
- [.github/workflows/ci.yml](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/.github/workflows/ci.yml)

`pgq_coop` 3.4 lets multiple connections consume one PgQ consumer cooperatively. The current upstream CI configuration covers PostgreSQL 14–18.

### Core Workflow

```sql
CREATE EXTENSION pgq;
CREATE EXTENSION pgq_coop;
SELECT pgq.create_queue('work');
SELECT pgq_coop.register_subconsumer('work', 'workers', 'worker_a');
SELECT pgq.insert_event('work', 'example', 'payload');
SELECT pgq.ticker();
SELECT pgq_coop.next_batch('work', 'workers', 'worker_a');
```

### Consumption and Maintenance

`pgq_coop.next_batch` returns a batch ID or NULL when no work is ready. Read its events using the normal PgQ batch API, complete application processing, then call `pgq_coop.finish_batch(batch_id)` with that returned ID. `pgq_coop.next_batch_custom` controls batch thresholds; the four-argument next-batch form accepts an inactivity interval for takeover. `pgq_coop.unregister_subconsumer` removes a worker registration.

Install as a superuser with the `pgq` dependency already available. The non-relocatable SQL extension creates the `pgq_coop` schema; its control selects `pg_catalog` as the installation schema. No new shared library or preload is required. PgQ still needs its normal ticker and maintenance. Worker failure and takeover can require reprocessing, so make handlers retry-safe; this extension does not replace a worker process.
