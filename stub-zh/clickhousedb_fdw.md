## 用法

来源：

- [官方上游 README](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/README.md)

`clickhousedb_fdw` — ClickHouse 外部数据包装器，支持 SELECT/INSERT，以及聚合与连接下推。

已复核目录快照记录的版本为 `1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "clickhousedb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'clickhousedb_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
