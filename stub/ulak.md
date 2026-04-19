
## Usage

> Sources: [README](https://github.com/zeybek/ulak/blob/main/README.md), [wiki](https://github.com/zeybek/ulak/wiki), [v0.0.2 release](https://github.com/zeybek/ulak/releases/tag/v0.0.2)

`ulak` implements the transactional outbox pattern inside PostgreSQL. Messages are inserted transactionally, then background workers deliver them asynchronously with retries, circuit breaking, and dead-letter handling.

### Enable the Extension

```sql
-- postgresql.conf
shared_preload_libraries = 'ulak'

CREATE EXTENSION ulak;
```

The README also shows setting `ulak.database` when running the included Docker example.

### Define Endpoints and Send Messages

```sql
SELECT ulak.create_endpoint(
  'my-webhook',
  'http',
  '{"url":"https://httpbin.org/post","method":"POST"}'::jsonb
);

BEGIN;
  INSERT INTO orders (id, total) VALUES (1, 99.99);
  SELECT ulak.send('my-webhook', '{"order_id":1,"total":99.99}'::jsonb);
COMMIT;
```

The same API pattern is documented for `kafka`, `mqtt`, `redis`, `amqp`, and `nats` endpoints.

### Delivery Controls and Pub/Sub

```sql
SELECT ulak.send_with_options(
  'my-webhook',
  '{"event":"order.created"}'::jsonb,
  5,
  NOW() + INTERVAL '10 minutes',
  'order-123-created',
  '550e8400-e29b-41d4-a716-446655440000'::uuid,
  NOW() + INTERVAL '1 hour',
  'order-123'
);

SELECT ulak.send_batch('my-webhook', ARRAY['{"id":1}'::jsonb, '{"id":2}'::jsonb]);

SELECT ulak.create_event_type('order.created', 'New order placed');
SELECT ulak.subscribe('order.created', 'my-webhook');
SELECT ulak.publish('order.created', '{"order_id":123}'::jsonb);
```

### Monitoring and Recovery

```sql
SELECT * FROM ulak.health_check();
SELECT * FROM ulak.get_worker_status();
SELECT * FROM ulak.get_endpoint_health();
SELECT * FROM ulak.dlq_summary();

SELECT ulak.redrive_message(42);
SELECT ulak.redrive_endpoint('my-webhook');
SELECT ulak.redrive_all();
SELECT ulak.replay_message(100);
```

### Key Settings

The README highlights these `ulak.` GUCs:

- `workers`
- `poll_interval`
- `batch_size`
- `max_queue_size`
- `circuit_breaker_threshold`
- `circuit_breaker_cooldown`
- `http_timeout`
- `dlq_retention_days`

The `v0.0.2` release notes only document a Docker publish workflow fix, so the current README and wiki remain the authoritative user-facing usage sources.
