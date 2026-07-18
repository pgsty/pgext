## 用法

来源：

- [官方扩展控制文件](https://github.com/nvorobev/pg_metricus_c/blob/085cf982da50211938cf014551cfe194ebc48b27/pg_metricus.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_metricus/)

`pg_metricus` — 从 PostgreSQL 代码通过 UDP 套接字向 Brubeck、Graphite 等聚合器发送自定义指标。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_metricus";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_metricus';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
