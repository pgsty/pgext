## 用法

来源：

- [官方上游文档](https://github.com/jnidzwetzki/pg_debug_scan#readme)

`pg_debug_scan` — 使用自定义 MVCC 快照定义调试 PostgreSQL 表扫描与元组可见性。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_debug_scan";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_debug_scan';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
