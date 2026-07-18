## 用法

来源：

- [上游 README](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/README)
- [扩展 control 文件](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/pgsmtp.control)
- [PL/Python 实现](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/pgsmtp--0.1.1.sql)
- [发行元数据](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/META.json)

`pgsmtp` 通过 Python SMTP 客户端从 PostgreSQL 发送邮件。版本 `0.1.1` 会创建保存发件人、服务器和密码记录的 `pgsmtp.user_smtp_data` 表，并提供 `pgsmtp.pg_smtp_mail()` 以及支持附件的 `pgsmtp.pg_smtp_mail_attach()`。它依赖不受信任语言 `plpython3u`，且必须由超级用户安装。

### 配置发件人并发送邮件

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pgsmtp;

INSERT INTO pgsmtp.user_smtp_data
  (smtp_user, smtp_server, smtp_port, smtp_pass)
VALUES
  ('sender@example.com', 'smtp.example.com', 465, 'replace-this-secret');

SELECT pgsmtp.pg_smtp_mail(
  'sender@example.com',
  'receiver@example.com',
  ARRAY[]::text[],
  'Status',
  'Job completed'
);
```

实现使用 `SMTP_SSL`、以明文保存密码，并在数据库后端内执行网络 I/O。更严重的是，它把调用者提供的发件人直接拼接进 PL/Python SQL 查询，存在 SQL 注入风险。`pgsmtp.pg_smtp_mail_attach()` 还会以 PostgreSQL 操作系统账号读取任意服务端路径。

上游警告仍有 bug，并把该扩展描述为仅限管理员使用。不要向应用角色授予其 schema、凭据表或函数执行权限。优先使用外部邮件队列；若不得不保留遗留用法，应参数化查询、加密或外置凭据、限制网络出口与文件访问、增加超时和审计，并在部署前测试 SMTP 故障行为。
