## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pgzint/)

`pgzint` — pgzint 使用 Zint 条码库在 PostgreSQL 内生成 PNG 条码与二维码。

已复核目录快照记录的版本为 `0.1.4`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgzint";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgzint';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
