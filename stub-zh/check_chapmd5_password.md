## 用法

来源：

- [官方扩展控制文件](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/check_chapmd5_password.control)

`check_chapmd5_password` — 用于验证 CHAP-MD5 认证响应的 PostgreSQL 函数。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "check_chapmd5_password";
SELECT extversion
FROM pg_extension
WHERE extname = 'check_chapmd5_password';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
