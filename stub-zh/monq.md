## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/monq.control)
- [官方上游文档](https://github.com/postgrespro/monq/blob/0c245caeeee1057144d91c57c2d7d229310afe7e/README.md)

`monq` — monq：为 PostgreSQL jsonb 数据提供类似 MongoDB 的查询支持。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`jsquery`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "monq";
SELECT extversion
FROM pg_extension
WHERE extname = 'monq';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
