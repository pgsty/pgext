## 用法

来源：

- [官方 README](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/README.md)
- [用法文档](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/doc/kafka.md)
- [扩展 SQL](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/sql/kafka.sql)
- [生产者实现](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/src/pg_kafka.c)
- [扩展控制文件](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/kafka.control)

`kafka` 0.0.1 来自已归档的 `pg_kafka` 项目，通过 librdkafka 从 PostgreSQL 后端直接向 Apache Kafka 发送文本消息。它面向 PostgreSQL 9.2+ 与 Kafka/librdkafka 0.8 时代的 API。尽管 SQL 注释声称消息与提交绑定，实现并没有提供可靠的事务性交付；应改用事务发件箱加外部生产者。

### 历史工作流

上游要求预加载动态库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_kafka.so'
```

安装扩展、登记代理并排入一条消息：

```sql
CREATE EXTENSION kafka;

INSERT INTO kafka.broker (host, port)
VALUES ('kafka.internal', 9092);

SELECT kafka.produce('audit-events', '{"event":"example"}');
SELECT kafka.close();
```

`kafka.produce(varchar, varchar)` 返回真，只表示 librdkafka 已把消息接受到本地生产者队列。`kafka.close()` 关闭当前数据库后端持有的生产者。代理配置只在该后端首次创建生产者时读取；修改 `kafka.broker` 后应关闭生产者或重新连接。

### 交付与事务注意事项

C 函数会立即调用 `rd_kafka_produce`。事务回调在提交时什么也不做，中止时只销毁生产者；它没有 PostgreSQL 提交记录、发件箱、Kafka 事务、幂等键、重试台账或两阶段协议。成功排队后仍可能在稍后失败，已中止的 SQL 事务也可能产生外部效果，后端或代理故障则可能丢失或重复消息。

复核的代码在成功排队后既不轮询也不刷新，因此交付回调与队列排空不能作为可靠确认路径。它还把有外部副作用的 `kafka.produce` 和 `kafka.close` 声明成 `IMMUTABLE`，使 PostgreSQL 可能以不适用于副作用的方式折叠、缓存、重排或省略调用。SQL 把 `kafka.close` 声明为返回布尔值，C 实现却返回 `void`。这些是源码缺陷，不是应用层可调参数。

### 运维边界

不要在现代集群部署这个已归档的原生模块，除非维护分支已经修正易变性、返回类型、生命周期、轮询、交付确认、内存/错误路径、TLS/认证以及当前 librdkafka/PostgreSQL 兼容性。触发器会让数据库写入依赖 Kafka 延迟和故障行为，却仍不能带来原子交付。若需要持久的变更发布，应在同一事务写入发件箱行并异步发布，或使用有监控偏移量与重放语义的受支持逻辑解码/CDC 工具。
