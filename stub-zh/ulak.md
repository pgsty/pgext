## 用法

来源：

- [README](https://api.pgxn.org/src/ulak/ulak-0.3.0/README.md)
- [Control](https://api.pgxn.org/src/ulak/ulak-0.3.0/ulak.control)
- [Changelog](https://api.pgxn.org/src/ulak/ulak-0.3.0/CHANGELOG.md)
- [Install SQL](https://api.pgxn.org/src/ulak/ulak-0.3.0/sql/ulak--0.3.0.sql)
- [Upgrade SQL](https://api.pgxn.org/src/ulak/ulak-0.3.0/sql/ulak--0.2.0--0.3.0.sql)

`ulak` 实现事务型 Outbox：消息与应用数据在同一个事务中入队，再由后台工作进程投递。远端投递为至少一次，消费者必须容忍重复消息。

### 启用

支持 PostgreSQL 14–18。设置 `shared_preload_libraries` 与 `ulak.database` 后重启 PostgreSQL，再在该库创建扩展。内置 HTTP；Kafka、MQTT、Redis Streams、AMQP 和 NATS 需要编译相应可选适配器及其客户端库。

```ini
shared_preload_libraries = 'ulak'
ulak.database = 'appdb'
```

```sql
CREATE EXTENSION ulak;
```

### 入队与发布

由管理员创建端点，通过 `ulak_application` 授予应用访问权。`ulak_admin` 管理端点和运维操作；`ulak_monitor` 读取监控数据。使用 `ulak.endpoint_status` 查看端点状态，避免暴露端点凭据。事务回滚后消息不会保留。

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

### 投递限制与重试

`ulak.send_with_options` 支持优先级、定时、幂等、过期和顺序控制。0.3.0 的所有入队接口均执行 `ulak.max_payload_size` 限制，超大载荷报 `54000`。批量接口按每行计入 `ulak.max_queue_size`，超限报 `53400`。卡在处理中的消息会消耗重试次数，耗尽后进入死信队列。工作进程使用 `synchronous_commit = local` 和 `lock_timeout = 10s`。

### 协议与升级变化

0.3.0 会重试 HTTP 408/425、Redis OOM/READONLY 以及无法路由的 mandatory AMQP 消息；Redis TLS 现会检查主机名。`ulak.mqtt_default_qos` 默认值为 1；多个工作进程会为配置的 MQTT 客户端 ID 附加进程编号。NATS 消息 ID 包含集群和数据库标识，AMQP 消息 ID 使用队列行 ID，应复核下游去重假设。安装匹配的库与 SQL，重启以加载工作进程库，再升级扩展。外部监听者需要设置 `ulak.wake_notify = on` 才会收到 `ulak_new_msg`；内置工作进程使用共享内存唤醒。

```sql
ALTER EXTENSION ulak UPDATE TO '0.3.0';
```

### 监控与恢复

按负载调整 `ulak.workers`、`ulak.poll_interval`、`ulak.batch_size`、重试和熔断配置。重新投递前应检查失败原因；重投和回放可能再次投递远端已经收到的消息。HTTP SSRF 限制默认生效，除非显式放宽。

```sql
SELECT * FROM ulak.health_check();
SELECT * FROM ulak.get_worker_status();
SELECT * FROM ulak.get_endpoint_health();
SELECT * FROM ulak.dlq_summary();
SELECT ulak.redrive_endpoint('orders-http');
```
