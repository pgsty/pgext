## 用法

来源：

- [官方扩展控制文件](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/kafka.control)
- [官方上游文档](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/README.md)

`kafka` — 从 PostgreSQL 事务中向 Apache Kafka 异步发送消息的预加载扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "kafka";
SELECT extversion
FROM pg_extension
WHERE extname = 'kafka';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
