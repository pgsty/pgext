## 用法

来源：

- [官方扩展控制文件](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail.control)

`pg_sendmail` — 通过调用服务器固定路径下兼容 sendmail 的可执行文件，从 SQL 发送电子邮件。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_sendmail";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sendmail';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
