## 用法

来源：

- [官方扩展控制文件](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/pg_txn_status.control)
- [官方上游文档](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/README.md)

`pg_txn_status` — 提供一字节的 txn_status 数据类型，用于存储事务状态值。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_txn_status";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_txn_status';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
