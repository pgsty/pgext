## 用法

来源：

- [上游 README](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/README.md)
- [详细使用指南](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/USAGE.md)
- [0.10.2 版控制文件](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/kadb_fdw.control)
- [偏移量存储实现](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/src/offsets.c)

kadb_fdw 0.10.2 是一个已归档的 Kafka 外部数据包装器，专门面向 Arenadata DB 与 Greenplum Database。它将 Kafka 分区分配到 MPP 段，并把消费者偏移量存储在 kadb.offsets 中。

### ADB 或 Greenplum 示例

只能在兼容的可丢弃 MPP 集群上配合测试主题使用：

```sql
CREATE EXTENSION kadb_fdw;
CREATE SERVER ka_server
FOREIGN DATA WRAPPER kadb_fdw
OPTIONS (k_brokers 'localhost:9092');
CREATE FOREIGN TABLE ka_table(payload text)
SERVER ka_server
OPTIONS (
  format 'text',
  k_topic 'test_topic',
  k_consumer_group 'test_group',
  k_seg_batch '5',
  k_timeout_ms '1000'
);
SELECT * FROM ka_table;
```

成功的 SELECT 会推进供后续读取使用的已存偏移量。它是消费操作，而不是远程表的幂等视图。

### 注意事项

- 安装与升级 SQL 使用分布式表、段执行和 MPP 选项等 Greenplum 专有语法。它不是标准 PostgreSQL FDW，无法不加修改地安装到社区 PostgreSQL。
- LIMIT 不会限制 Kafka 获取量：包装器按批次取数，并将偏移推进到最后一条已获取消息，包括未出现在返回行中的消息。
- 偏移辅助函数不为 Kafka 提供事务保证，多项分区同步操作也明确为非原子。重试与恰好一次语义应在该 FDW 之外协调。
- 部分轮询阶段在配置超时结束前无法取消，总查询时长可能跨不同分区累计多个超时间隔。
- 用户提供的 Avro 模式不会与消息进行验证；上游说明不匹配可导致无效内存分配错误。
- Kerberos 支持使用 SASL 明文协议，并要求每个段上都有相同的服务器可读 keytab 路径。应保护服务器选项、目录、文件与日志。
- 仓库于 2021 年归档，面向旧版 ADB、Greenplum、librdkafka、Avro 与 CSV 库。其安装脚本还会重置会话 client_min_messages 设置。
