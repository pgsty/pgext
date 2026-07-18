## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/pg_datatype_password.control)
- [官方上游文档](https://pgxn.org/dist/pg_datatype_password/1.0.0/)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_datatype_password/)

`pg_datatype_password` — 纯 SQL 密码数据类型与比较运算符，使用 pgcrypto Blowfish 哈希，并依赖触发器在写入前加密。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pgcrypto`, `plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_datatype_password";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_datatype_password';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
