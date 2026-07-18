## 用法

来源：

- [官方上游文档](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/README.md)

`pg_ulid` — 通过 C 函数在 PostgreSQL 中生成 ULID 字符串。

已复核目录快照记录的版本为 `0.0.4`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_ulid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ulid';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
