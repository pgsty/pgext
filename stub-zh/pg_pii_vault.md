## 用法

来源：

- [官方扩展控制文件](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/pg_pii_vault.control)

`pg_pii_vault` — 通过 HashiCorp Vault Transit Engine 提供面向 GDPR 的列级加密与加密擦除。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_pii_vault";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_pii_vault';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
