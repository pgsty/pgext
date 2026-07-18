## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/openssl/openssl-0.0.1/openssl.control)
- [官方上游文档](https://pgxn.org/dist/openssl/0.0.1/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/openssl/)

`openssl` — openssl：为 PostgreSQL 提供基于 OpenSSL 的加密函数与类型。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "openssl";
SELECT extversion
FROM pg_extension
WHERE extname = 'openssl';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
