## 用法

来源：

- [官方扩展控制文件](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/hbase_fdw.control)

`hbase_fdw` — 用于从 PostgreSQL 查询 HBase 的外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "hbase_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hbase_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
