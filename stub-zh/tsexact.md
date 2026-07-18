## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/tsexact.control)
- [官方上游文档](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/README.md)

`tsexact` — 为 PostgreSQL 9.5 及更早版本提供精确短语匹配的全文检索辅助函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tsexact";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsexact';
```

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
