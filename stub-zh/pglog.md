## 用法

来源：

- [官方扩展控制文件](https://github.com/2ndQuadrant/pglog/blob/5d57809e4e73913a06b4aaa6553f53f43d010a39/pglog.control)

`pglog` — pglog：外部数据访问类 PostgreSQL 扩展；上游说明为“通过 SQL 访问 PostgreSQL 日志”。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pglog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pglog';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
