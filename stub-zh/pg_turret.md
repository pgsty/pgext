## 用法

来源：

- [Official README](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/README.md)
- [Extension control file](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/pg_turret.control)
- [GUC and background-worker implementation](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/src/lib.rs)
- [Log ring-buffer implementation](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/src/log_capture.rs)
- [Cargo version and PostgreSQL features](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/Cargo.toml)

`pg_turret` 挂接 PostgreSQL 日志管线，把日志记录转换为结构化事件，缓存在共享内存中，再由后台 Worker 分批导出到 HTTP、Kafka 或 Sentry。它避免在产生日志的后端中执行网络 I/O，但属于异步可观测性通道，并非持久审计日志。

### 预加载与配置

钩子、共享内存和 Worker 必须随服务器启动加载，并需要重启：

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

重启后创建扩展以登记 SQL 函数，并检查投递计数：

```sql
CREATE EXTENSION pg_turret;
SELECT * FROM get_metrics();
```

Kafka 使用 `pg_turret.kafka.brokers`、`topic`、凭据、超时和批大小。Sentry 使用 `pg_turret.sentry.dsn`、environment/release/server name、采样率、最低级别、SQL/PII 开关和每 Worker 事件速率限制。除非策略明确许可，应保持 SQL 和 PII 导出关闭。

### 传输语义与监控

`get_metrics()` 报告 `logs_captured`、`logs_dropped`、`logs_sent`、`logs_retry_failed`、`logs_pending` 和环形缓冲容量。必须持续监控 dropped 与 retry-failed 计数。

环形缓冲写满时，实现会覆盖最旧事件。重试队列有界且只存在内存中；进程或主机故障会丢失待处理事件，重试也可能重复发送已被外部接收的事件。应把 PostgreSQL 正常日志输出作为权威记录，并让消费者容忍缺口和重复。

### 安全与版本边界

日志字段可能包含查询文本、数据库/用户名、错误详情、上下文、文件名及其他敏感值；网络目标会扩大数据边界。启用任何适配器前，必须测试 TLS、认证、脱敏、限速、背压和目标故障行为。

已复核源码注册凭据 GUC 时没有显式的仅超级用户可见标志。应在目标构建上验证哪些角色能通过 `pg_settings` 或 `SHOW` 读取这些值；如果不可信角色可以看到，就不要把长期密钥放入其中。

版本元数据不一致：目录快照记录 0.0.0，Cargo 记录 0.1.0，control 文件则依赖构建期的 Cargo 版本替换。README 写 PostgreSQL 13–17，Cargo 却声明 13–18 特性。应记录实际制品版本和经过测试的大版本，不能从单一文件推断兼容性。
