

## 用法

> [pg_utl_smtp: PostgreSQL 的 Oracle UTL_SMTP 兼容扩展](https://github.com/hexacluster/pg_utl_smtp)

### 启用

```sql
CREATE EXTENSION plperlu;
CREATE EXTENSION pg_utl_smtp;
```

### 发送邮件

```sql
DO $$
DECLARE
    c utl_smtp.connection;
BEGIN
    c := utl_smtp.open_connection('smtp.example.com', 25);
    CALL utl_smtp.ehlo(c, 'mydomain.com');
    CALL utl_smtp.mail(c, 'sender@example.com');
    CALL utl_smtp.rcpt(c, 'recipient@example.com');
    CALL utl_smtp.open_data(c);
    CALL utl_smtp.write_data(c, 'From: sender@example.com' || E'\r\n');
    CALL utl_smtp.write_data(c, 'To: recipient@example.com' || E'\r\n');
    CALL utl_smtp.write_data(c, 'Subject: Test Email' || E'\r\n');
    CALL utl_smtp.write_data(c, E'\r\n');
    CALL utl_smtp.write_data(c, 'Hello from PostgreSQL!');
    CALL utl_smtp.close_data(c);
    CALL utl_smtp.quit(c);
END;
$$;
```

### 过程

- **OPEN_CONNECTION(host, port, tx_timeout, ...)** - 打开到 SMTP 服务器的连接。返回 `utl_smtp.connection` 类型。通过 `secure_connection_before_smtp` 支持 SSL/TLS。
- **EHLO(c, domain)** / **HELO(c, domain)** - 执行初始 SMTP 握手。
- **MAIL(c, sender)** - 发起邮件事务。
- **RCPT(c, recipient)** - 指定邮件收件人。多个收件人可多次调用。
- **OPEN_DATA(c)** - 发送 DATA 命令开始消息正文。
- **WRITE_DATA(c, data)** - 写入消息正文的一部分。
- **WRITE_RAW_DATA(c, data)** - 向消息正文写入原始数据。
- **CLOSE_DATA(c)** - 关闭数据会话。
- **QUIT(c)** - 终止 SMTP 会话并断开连接。

### 连接类型

```sql
-- utl_smtp.connection 复合类型
(host varchar(255), port integer, tx_timeout integer,
 private_tcp_con integer, private_state integer)
```

### 注意事项

- 需要系统上安装 Perl `Net::SMTP` 模块
- 使用 `E'\r\n'` 作为换行符替代 `utl_tcp.crlf`
- `wallet_path` 和 `wallet_password` 参数未使用
