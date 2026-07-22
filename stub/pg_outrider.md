## Usage

Sources:

- [Version 1.0 SQL interface](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider--1.0.sql)
- [Dynamic background-worker implementation](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.c)
- [Extension control file](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.control)

`pg_outrider` launches a dynamic background worker that watches one relation's free-space map and physically extends its main fork when a free-space watermark is not observed near the tail. It is experimental storage-management code built on PostgreSQL internals, not an ordinary capacity reservation API.

### Core Workflow

Only a database administrator should launch a worker after validating the relation and sizes:

```sql
CREATE EXTENSION pg_outrider;
REVOKE EXECUTE ON FUNCTION pg_outrider_launch(regclass, bigint, bigint) FROM PUBLIC;

CREATE TABLE event_queue (id bigint, payload jsonb);
SELECT pg_outrider_launch('event_queue'::regclass, 16, 32);
```

`pg_outrider_launch(regclass, bigint, bigint)` interprets the final arguments as the extension increment and watermark in MiB; defaults are 1 and 2. It returns the PID of the registered worker. The worker is dynamic, so `shared_preload_libraries` is not required, but each launch consumes a background-worker slot.

### Security and Operational Boundaries

The SQL function checks only SELECT privilege on the relation, yet its worker modifies physical storage, and extension functions are executable by PUBLIC unless privileges are changed. Revoke access and limit launches to trusted administrators. The C implementation narrows `bigint` sizes into C integers without adequate positive-range validation; reject zero, negative, and oversized values before calling it.

There is no SQL stop, status, or duplicate-worker function. A worker does not automatically restart after failure and can grow storage unexpectedly. The free-space map is a heuristic, not an exact free-byte counter. The 2017 code directly uses heap, page, fork, and free-space-map internals, so validate relation types, crash recovery, replication, backup, bloat, and exact major-version compatibility in a disposable cluster.
