## 用法

来源：

- [官方上游文档](https://github.com/PrimeDataConversion/substrait-postgres/blob/d9ead4261c32d86e3a19bfcf72ce591c8359d6ed/README.md)

`pg_substrait` — 通过 PostgreSQL 函数执行 Substrait protobuf 或 JSON 查询计划

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_substrait";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_substrait';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
