## 用法

来源：

- [官方扩展控制文件](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/hive_fdw.control)
- [官方上游文档](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/README.md)

`hive_fdw` — 用于从 PostgreSQL 查询 Apache Hive 的外部数据包装器。

已复核目录快照记录的版本为 `4.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "hive_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hive_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
