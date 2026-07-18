## Usage

Sources:

- [Version 1.0 SQL interface](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider--1.0.sql)
- [Dynamic background-worker implementation](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.c)
- [Extension control file](https://github.com/meistervonperf/pg_outrider/blob/ae1462c9c75f5dfd95d0db02ec9d4a68675afb2b/pg_outrider.control)

`pg_outrider` launches a dynamic background worker that watches one heap relation's free-space watermark. When free space falls below the requested watermark, the worker extends the relation's main fork by the requested number of megabytes and updates its free-space map.

```sql
CREATE EXTENSION pg_outrider;

CREATE TABLE event_queue (id bigint, payload jsonb);
SELECT pg_outrider_launch('event_queue'::regclass, 16, 32);
```

The second and third arguments are the extension increment and watermark in MiB; their defaults are 1 and 2. A successful call returns the process ID of a newly registered worker. The worker is registered dynamically, so upstream does not require `shared_preload_libraries`.

### Caveats

Each call consumes a background-worker slot and starts a worker for that relation; avoid duplicate launches and monitor storage growth. The worker stops rather than automatically restarting after failure. The reviewed implementation dates from 2017, uses PostgreSQL internal storage APIs, and publishes no compatibility matrix or README, so validate it against the exact server major before deployment.
