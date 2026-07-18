## 用法

来源：

- [已复核 commit 的 pg_kafka_events README](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/README.md)
- [已复核 commit 的 pg_kafka_events.control](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/pg_kafka_events.control)
- [后台工作进程实现](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/src/pg_kafka_events.c)
- [构建依赖](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/Makefile)
- [版本 1.0 的安装 SQL](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/pg_kafka_events--1.0.sql)

`pg_kafka_events` 是一个预加载后台工作进程，它会 fork `pg_recvlogical`，读取按换行分隔的逻辑解码输出，并通过 librdkafka 发布每项变更。默认解码器为 `decoding_json`；一个已配置数据库会被流式发送到一个 Kafka 主题。

### 服务器配置

```conf
shared_preload_libraries = 'pg_kafka_events'
wal_level = logical
max_replication_slots = 4
max_wal_senders = 4
kafka.database_name = 'appdb'
kafka.topic = 'row_changes'
kafka.bootstrap_servers = 'kafka1:9092'
```

重启 PostgreSQL 前，应安装 librdkafka、选定的逻辑解码插件和配置的 `pg_recvlogical` 二进制。重启后，以超级用户登记扩展：

```sql
CREATE EXTENSION pg_kafka_events;
```

安装 SQL 为空；预加载会启动工作进程，而 `CREATE EXTENSION` 只在该数据库中登记版本 `1.0`。

### 注意事项

- 工作进程始终使用复制槽名 `kafka_events`。消费者停止或失败可能无限保留 WAL；应监控复制槽延迟，并定义运维恢复流程。
- 这个 2017 年的实现不提供事务分组、暴露给 PostgreSQL 的确认、与复制位置协调的重试或 exactly-once 保证。必须验证 Kafka 和进程故障下的重复与丢失行为。
- 解码器行被累积到固定 64 KiB 静态缓冲区，且没有长度保护。大型行变更消息可能使其溢出；修复前不要部署已复核代码。
- 已发布载荷还会作为 PostgreSQL notice 输出，可能把行数据复制到服务器日志。尽管控制文件和目录使用 `1.0`，上游唯一 tag 是 `0.1`。
