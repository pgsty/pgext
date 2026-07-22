## Usage

Sources:

- [Official README](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/README.md)
- [Extension control file](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/pg_turret.control)
- [GUC and background-worker implementation](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/src/lib.rs)
- [Log ring-buffer implementation](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/src/log_capture.rs)
- [Cargo version and PostgreSQL features](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/Cargo.toml)

`pg_turret` hooks PostgreSQL's logging pipeline, converts log records to structured events, buffers them in shared memory, and exports batches from background workers to HTTP, Kafka, or Sentry. It avoids network I/O in the emitting backend, but it is an asynchronous observability path rather than a durable audit log.

### Preload and Configure

The hook, shared memory, and workers require startup loading and a restart:

```conf
shared_preload_libraries = 'pg_turret'

pg_turret.ring_buffer_size = 4096
pg_turret.num_workers = 2
pg_turret.poll_interval_s = 2
pg_turret.filter.level_min = 19

pg_turret.http.enabled = on
pg_turret.http.endpoint = 'https://logs.example.com/v1/ingest'
pg_turret.http.batch_size = 100
pg_turret.http.timeout_ms = 5000
```

After restart, create the extension to register its SQL functions and inspect delivery counters:

```sql
CREATE EXTENSION pg_turret;
SELECT * FROM get_metrics();
```

Kafka uses `pg_turret.kafka.brokers`, `topic`, credentials, timeout, and batch size. Sentry uses `pg_turret.sentry.dsn`, environment/release/server name, sampling, minimum level, SQL/PII switches, and a per-worker event-rate limit. Keep SQL and PII export disabled unless policy explicitly permits them.

### Delivery Semantics and Monitoring

`get_metrics()` reports `logs_captured`, `logs_dropped`, `logs_sent`, `logs_retry_failed`, `logs_pending`, and ring capacity. Monitor dropped and retry-failed counters continuously.

When the ring is full, the implementation overwrites the oldest event. Retry queues are bounded and in memory; process or host failure loses pending events, and retries can duplicate externally accepted events. Use PostgreSQL's normal log sink as the authoritative record and design consumers to tolerate gaps and duplicates.

### Security and Version Boundaries

Log fields can include query text, database/user names, error detail, context, filenames, and other sensitive values. Network endpoints widen the data boundary. Test TLS, authentication, redaction, rate limiting, backpressure, and destination outage behavior before enabling any adapter.

Credential GUCs are registered without an explicit superuser-only visibility flag in the reviewed source. Verify which roles can read `pg_settings` or `SHOW` those values on the target build; do not place long-lived secrets there if untrusted roles can observe them.

Version metadata is inconsistent: the catalog snapshot says 0.0.0, Cargo says 0.1.0, and the control file relies on build-time Cargo version substitution. The README names PostgreSQL 13–17 while Cargo declares features for 13–18. Record the actual artifact version and tested major rather than inferring compatibility from one file.
