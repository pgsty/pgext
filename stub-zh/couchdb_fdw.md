## 用法

来源：

- [官方扩展控制文件](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/couchdb_fdw.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/couchdb_fdw/)
- [官方项目或服务商页面](https://pgxn.org/extension/couchdb_fdw)

`couchdb_fdw` — couchdb_fdw：用于从 PostgreSQL 查询 CouchDB 文档数据库的外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "couchdb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'couchdb_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
