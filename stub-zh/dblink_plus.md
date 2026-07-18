## 用法

来源：

- [官方扩展控制文件](https://github.com/ossc-db/dblink_plus/blob/cb0efffe7f9098e352ccf5da65738630cf40dc5a/dblink_plus.control)
- [官方上游文档](https://ossc-db.github.io/dblink_plus/index.html)
- [官方项目或服务商页面](https://ossc-db.github.io/dblink_plus/)

`dblink_plus` — dblink_plus：在 PostgreSQL 内连接 PostgreSQL、Oracle、MySQL、SQLite 等外部数据库并执行 SQL。

已复核目录快照记录的版本为 `1.0.11`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dblink_plus";
SELECT extversion
FROM pg_extension
WHERE extname = 'dblink_plus';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
