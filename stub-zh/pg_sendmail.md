## 用法

来源：

- [官方 README](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/README.md)
- [扩展控制文件](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail.control)
- [扩展 SQL](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail--0.0.2.sql)
- [邮件实现](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail.c)

`pg_sendmail` 从 PostgreSQL 后端把生成的邮件文本通过管道发送给 `/usr/sbin/sendmail`。固定的 0.0.2 实现在地址处理处存在命令注入漏洞，因此不能向不受信输入或用户开放；未经代码修复不应部署。

### 安装接口

若要审计现有安装，应在任何功能测试之前先进行隔离：

```sql
CREATE EXTENSION pg_sendmail;

REVOKE ALL ON FUNCTION mail(text, text, text, text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION sendmail(text, text, text, text) FROM PUBLIC;
```

`mail(mailfrom,rcptto,subject,msg_body,headers)` 是 C 入口。`sendmail(mailfrom,rcptto,subject,msg_body)` 是对主题编码并提供 MIME 头的 PL/pgSQL 包装函数。两者都返回本地 sendmail 调用的布尔结果；该值不能证明邮件最终送达。

### 严重安全与事务注意事项

C 代码把 `mailfrom` 与 `rcptto` 直接插入字符串，再把该字符串交给 `popen` 执行 shell 命令。这些值既没有正确引用，也没有作为安全参数向量传递，因此构造的输入可以用 PostgreSQL 服务账号执行操作系统命令。邮件头和主题数据也能改变消息结构。从 `PUBLIC` 撤销只是隔离措施，不是完整修复；启用功能前必须改用非 shell API，并严格校验地址与邮件头。

发送邮件是同步操作，可能阻塞数据库后端，而且属于 PostgreSQL 回滚无法撤销的外部副作用。应在事务中写入应用通知队列，再由独立的最小权限工作进程发送。SQL 脚本还会把 `sendmail(text,text,text,text)` 的所有者改成硬编码的 `postgres` 角色，可能导致安装失败或产生意外所有者。必须在实际主机上审查软件包代码、角色所有权、可执行文件路径、MTA 策略、超时、审计日志与失败处理。
