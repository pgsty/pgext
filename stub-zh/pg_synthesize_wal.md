## 用法

来源：

- [官方上游文档](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/README.md)

`pg_synthesize_wal` — 生成不同大小的 PostgreSQL WAL 记录用于开发测试

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_walinspect`。
整理后的兼容版本集合为 `15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_synthesize_wal";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_synthesize_wal';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
