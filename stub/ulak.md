## Usage

Sources:

- [README](https://api.pgxn.org/src/ulak/ulak-0.3.0/README.md)
- [Control](https://api.pgxn.org/src/ulak/ulak-0.3.0/ulak.control)
- [Changelog](https://api.pgxn.org/src/ulak/ulak-0.3.0/CHANGELOG.md)
- [Install SQL](https://api.pgxn.org/src/ulak/ulak-0.3.0/sql/ulak--0.3.0.sql)
- [Upgrade SQL](https://api.pgxn.org/src/ulak/ulak-0.3.0/sql/ulak--0.2.0--0.3.0.sql)

`ulak` implements a transactional outbox: enqueue messages in the same transaction as application data, then deliver them through background workers. Delivery to remote systems is at least once; consumers must tolerate duplicates.

### Enablement

PostgreSQL 14–18 is supported. Set `shared_preload_libraries` and `ulak.database`, restart PostgreSQL, then create the extension in that database. HTTP is built in; Kafka, MQTT, Redis Streams, AMQP and NATS require their optional compiled adapters and client libraries.

```ini
shared_preload_libraries = 'ulak'
ulak.database = 'appdb'
```

```sql
CREATE EXTENSION ulak;
```

### Enqueue and publish

Create endpoints as an administrator; grant application access through `ulak_application`. `ulak_admin` manages endpoints and operations; `ulak_monitor` reads monitoring data. Use `ulak.endpoint_status` for endpoint status without exposing endpoint credentials. A rolled-back transaction leaves no message.

```sql
SELECT ulak.create_endpoint('orders-http', 'http',
  '{"url":"https://example.com/webhook","method":"POST"}'::jsonb);
BEGIN;
SELECT ulak.send('orders-http', '{"event":"order.created","id":123}'::jsonb);
COMMIT;

SELECT ulak.send_batch('orders-http', ARRAY['{"id":1}'::jsonb, '{"id":2}'::jsonb]);
SELECT ulak.create_event_type('order.created', 'Order created');
SELECT ulak.subscribe('order.created', 'orders-http');
SELECT ulak.publish('order.created', '{"id":123}'::jsonb);
```

### Delivery limits and retries

`ulak.send_with_options` adds priority, scheduling, idempotency, expiry and ordering controls. In 0.3.0 all enqueue APIs enforce `ulak.max_payload_size`; oversized payloads raise `54000`. Batch APIs count each row against `ulak.max_queue_size` and raise `53400` on overflow. Stale processing rows consume retry budget and eventually enter the dead-letter queue. Workers use `synchronous_commit = local` and `lock_timeout = 10s`.

### Protocol and upgrade changes

0.3.0 retries HTTP 408/425, Redis OOM/READONLY and unroutable mandatory AMQP messages. Redis TLS now checks hostnames. `ulak.mqtt_default_qos` defaults to 1; multiple workers suffix configured MQTT client IDs with their worker number. NATS message IDs include cluster and database identity; AMQP message IDs use the queue row ID. Review downstream deduplication assumptions. Install the matching library and SQL, restart to load the worker library, and upgrade the extension. External listeners need `ulak.wake_notify = on` to receive `ulak_new_msg`; native workers use shared-memory wakeups.

```sql
ALTER EXTENSION ulak UPDATE TO '0.3.0';
```

### Monitoring and recovery

Tune `ulak.workers`, `ulak.poll_interval`, `ulak.batch_size`, retry and circuit-breaker settings for the workload. Inspect failures before redriving; redrive and replay can deliver an already-observed message again. HTTP SSRF restrictions apply unless explicitly relaxed.

```sql
SELECT * FROM ulak.health_check();
SELECT * FROM ulak.get_worker_status();
SELECT * FROM ulak.get_endpoint_health();
SELECT * FROM ulak.dlq_summary();
SELECT ulak.redrive_endpoint('orders-http');
```
