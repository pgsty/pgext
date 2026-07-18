## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/salesforce_id/)

`salesforce_id` — 用于 Salesforce ID 的 12 字节紧凑 PostgreSQL 类型，提供类型转换、比较运算符以及 B-tree/Hash 运算符类。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "salesforce_id";
SELECT extversion
FROM pg_extension
WHERE extname = 'salesforce_id';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
