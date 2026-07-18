## 用法

来源：

- [官方扩展控制文件](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/pgbrtypes.control)
- [官方上游文档](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/README.md)

`pgbrtypes` — 提供巴西 CPF 与 CNPJ 编号数据类型，支持校验、格式化、类型转换和索引。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgbrtypes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgbrtypes';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
