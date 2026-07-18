## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/pg_oltp_bench.control)

`pg_oltp_bench` — 用于运行类 sysbench OLTP 测试的辅助函数与 pgbench 脚本。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_oltp_bench";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_oltp_bench';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
