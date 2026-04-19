
## 用法

> 来源：[README](https://github.com/zeybek/ulak/blob/main/README.md), [wiki](https://github.com/zeybek/ulak/wiki), [v0.0.2 release](https://github.com/zeybek/ulak/releases/tag/v0.0.2)

`ulak` 在 PostgreSQL 内实现 transactional outbox pattern。消息会在事务中写入，然后由后台 worker 异步投递，并带有重试、circuit breaking 和 dead-letter 处理。

### 启用扩展

```sql
-- postgresql.conf
shared_preload_libraries = 'ulak'

CREATE EXTENSION ulak;
```

README 还展示了在自带 Docker 示例里设置 `ulak.database`。

### 定义 Endpoint 并发送消息

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

同一套 API 也记录了 `kafka`、`mqtt`、`redis`、`amqp` 和 `nats` endpoint 的用法。

### 投递控制与发布订阅

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

### 监控与恢复

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

### 关键配置

README 点名列出了这些 `ulak.` GUC：

- `workers`
- `poll_interval`
- `batch_size`
- `max_queue_size`
- `circuit_breaker_threshold`
- `circuit_breaker_cooldown`
- `http_timeout`
- `dlq_retention_days`

### 说明

`v0.0.2` release notes 只记录了 Docker publish workflow 修复，因此当前 README 和 wiki 仍是用户侧最权威的使用来源。
