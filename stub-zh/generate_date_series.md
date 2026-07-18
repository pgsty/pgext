## 用法

来源：

- [官方扩展控制文件](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/generate_date_series.control)
- [官方上游文档](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/doc/generate_date_series.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/generate_date_series/)

`generate_date_series` — 为日期范围提供类似 generate_series() 的集合返回函数，以整数天数作为步长。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "generate_date_series";
SELECT extversion
FROM pg_extension
WHERE extname = 'generate_date_series';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
