## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/otp/otp-1.0.0/otp.control)
- [官方上游文档](https://pgxn.org/dist/otp/1.0.0/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/otp/)

`otp` — otp：安全类 PostgreSQL 扩展；上游说明为“与 Google Authenticator 兼容的 TOTP 实现”。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plperlu`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "otp";
SELECT extversion
FROM pg_extension
WHERE extname = 'otp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
