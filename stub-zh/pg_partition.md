## 用法

来源：

- [官方扩展控制文件](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/pg_partition.control)

`pg_partition` — 生成并管理基于时间戳的 PostgreSQL 表分区

已复核目录快照记录的版本为 `0.2.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_partition";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_partition';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
