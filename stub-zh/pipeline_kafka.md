## 用法

来源：

- [官方扩展控制文件](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/pipeline_kafka.control)
- [官方上游文档](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/README.md)

`pipeline_kafka` — 为 PipelineDB 集成 Kafka，提供后台消费、消息生产、偏移量与延迟监控

已复核目录快照记录的版本为 `1.0.0`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pipelinedb`。

```sql
CREATE EXTENSION "pipeline_kafka";
SELECT extversion
FROM pg_extension
WHERE extname = 'pipeline_kafka';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
