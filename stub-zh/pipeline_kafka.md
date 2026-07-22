## 用法

来源：

- [官方 README](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/README.md)
- [扩展 SQL API](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/pipeline_kafka--1.0.0.sql)
- [Kafka worker 实现](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/pipeline_kafka.c)

`pipeline_kafka` 1.0.0 是历史性的 PipelineDB Kafka 集成。后台 worker 把 Kafka 主题消费到 PipelineDB 流中，SQL 函数或触发器则产生消息。它依赖已经淘汰的 `pipelinedb` 扩展和旧 librdkafka 接口；按设计就不能直接用于普通的当前 PostgreSQL。

### 预加载与 Broker 设置

使用上游文档指定的 librdkafka 版本构建，预加载两个模块，再重启 PostgreSQL：

```conf
shared_preload_libraries = 'pipelinedb,pipeline_kafka'
```

创建扩展并登记 broker：

```sql
CREATE EXTENSION pipeline_kafka CASCADE;

SELECT pipeline_kafka.add_broker('kafka-1:9092');
SELECT pipeline_kafka.add_broker('kafka-2:9092');
```

扩展使用固定的 `pipeline_kafka` 模式，并在其中的表里保存 broker、消费者定义和偏移量。

### 消费与生产

```sql
SELECT pipeline_kafka.consume_begin(
  'events',
  'public.event_stream',
  format      => 'text',
  delimiter   => E'\t',
  batchsize   => 10000,
  parallelism => 4
);

SELECT * FROM pipeline_kafka.consumer_lag;

SELECT pipeline_kafka.produce_message(
  'notifications',
  convert_to('hello', 'UTF8'),
  partition => 2
);

SELECT pipeline_kafka.consume_end('events', 'public.event_stream');
```

消费者目标是 PipelineDB 流，并不是任意现代数据表。无参数 `consume_begin()` 和 `consume_end()` 会启动或停止全部已记录消费者。扩展还提供分区流专用函数与 `topic_watermarks(topic, partition)`；`consumer_lag` 会比较已保存偏移与 broker 水位。

若要发送变更行，可创建 `AFTER INSERT OR UPDATE FOR EACH ROW` 触发器并使用 `pipeline_kafka.emit_tuple('topic')`。它会把元组序列化为 JSON。触发器消息和 `produce_message` 都是外部副作用，不能与 PostgreSQL 事务原子回滚。

### 遗留与安全边界

上游指定的是 PipelineDB 与 librdkafka 0.11.6 时代的 API。项目没有记录当前 Kafka TLS/SASL 配置、幂等生产、恰好一次投递或现代消费者组保证。历史 SQL 甚至把这些有副作用的函数声明为 `IMMUTABLE`，因此不能依赖面向纯函数的规划器语义。

PipelineDB 已经停止发展，其 PostgreSQL 分支、流目录和后台 worker API 都与受支持 PostgreSQL 发生分歧。除非有维护中的下游移植版，提供经过测试的兼容矩阵、安全 broker 配置、偏移恢复、重复处理以及停机/再平衡测试，否则应把该扩展视为归档项目。
