## 用法

来源：

- [官方上游 README](https://github.com/pgsql-io/cassandra_fdw/blob/2ee3d7950954f99464c6f3249d622136e558cfc8/README.md)

`cassandra_fdw` — 用于从 PostgreSQL 11+ 读写 Cassandra 3+ 的外部数据包装器。

已复核目录快照记录的版本为 `3.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "cassandra_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'cassandra_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
