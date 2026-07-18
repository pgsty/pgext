## 用法

来源：

- [官方扩展控制文件](https://github.com/truthly/pg-webauthn/blob/c1f8099db1cbed3a89cf12a2668ed67bb10ee354/webauthn.control)

`webauthn` — 用纯 SQL 在 PostgreSQL 中实现 WebAuthn 凭据注册与断言验证。

已复核目录快照记录的版本为 `1.6`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`cbor`, `pg_ecdsa_verify`, `pgcrypto`, `plpgsql`。

```sql
CREATE EXTENSION "webauthn";
SELECT extversion
FROM pg_extension
WHERE extname = 'webauthn';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
