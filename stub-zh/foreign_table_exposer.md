## 用法

来源：

- [官方扩展控制文件](https://github.com/komamitsu/foreign_table_exposer/blob/1c57a554d1c50c872fac3b103317b11dec10e95f/foreign_table_exposer.control)
- [官方上游文档](https://pgxn.org/dist/foreign_table_exposer/1.0.0/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/foreign_table_exposer/)

`foreign_table_exposer` — 让基于系统目录发现表的客户端把外部表识别为普通表。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "foreign_table_exposer";
SELECT extversion
FROM pg_extension
WHERE extname = 'foreign_table_exposer';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
