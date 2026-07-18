## 用法

来源：

- [官方扩展控制文件](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/csv_tam.control)

`csv_tam` — PostgreSQL 的 CSV 表访问方法，将表元组以 CSV 格式存入关系文件。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "csv_tam";
SELECT extversion
FROM pg_extension
WHERE extname = 'csv_tam';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
