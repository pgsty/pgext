## 用法

来源：

- [官方扩展控制文件](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/schema_triggers.control)
- [官方上游文档](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/README.md)

`schema_triggers` — 已归档的 C 钩子扩展，为 PostgreSQL 事件触发器增加关系、列及触发器的模式变更事件。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "schema_triggers";
SELECT extversion
FROM pg_extension
WHERE extname = 'schema_triggers';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
