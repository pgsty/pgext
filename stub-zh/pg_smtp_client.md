

## 用法

### 启用扩展

连接到 PostgreSQL 并执行以下命令。

```sql
CREATE EXTENSION IF NOT EXISTS pg_smtp_client CASCADE;
```

使用 `smtp_client.send_email()` 函数发送电子邮件。

### 函数参数

| 参数          | 类型    | 说明                                               | 系统配置（可选）                |
|---------------|---------|----------------------------------------------------|---------------------------------|
| subject       | text    | 邮件主题                                           |                                 |
| body          | text    | 邮件正文                                           |                                 |
| html          | boolean | 正文是否为 HTML 格式（true）或纯文本格式（false）  |                                 |
| from_address  | text    | 发件人邮箱地址                                     | `smtp_client.from_address`      |
| recipients    | text[]  | 收件人邮箱地址列表                                 |                                 |
| ccs           | text[]  | 抄送邮箱地址列表                                   |                                 |
| bccs          | text[]  | 密送邮箱地址列表                                   |                                 |
| smtp_server   | text    | 使用的 SMTP 服务器地址                             | `smtp_client.server`            |
| smtp_port     | integer | SMTP 服务器端口                                    | `smtp_client.port`              |
| smtp_tls      | boolean | 是否使用 TLS 加密                                  | `smtp_client.tls`               |
| smtp_username | text    | SMTP 服务器登录用户名                              | `smtp_client.username`          |
| smtp_password | text    | SMTP 服务器登录密码                                | `smtp_client.password`          |

### 默认配置

您可以为上表中标注的部分参数配置系统级默认值，方法如下：

```
ALTER SYSTEM SET smtp_client.server TO 'smtp.example.com';
ALTER SYSTEM SET smtp_client.port TO 587;
ALTER SYSTEM SET smtp_client.tls TO true;
ALTER SYSTEM SET smtp_client.username TO 'MySmtpUsername';
ALTER SYSTEM SET smtp_client.password TO 'MySmtpPassword';
ALTER SYSTEM SET smtp_client.from_address TO 'from@example.com';
SELECT pg_reload_conf();
```

### 使用示例

发送邮件：
```sql
SELECT smtp_client.send_email('test subject', 'test body', false, 'from@example.com', array['to@example.com'], null, null, 'smtp.example.com', 587, true, 'username', 'password');
```

使用已配置的默认值发送邮件：
```sql
SELECT smtp_client.send_email('test subject', 'test body', false, null, array['to@example.com']);
```

或者使用命名参数：
```sql
SELECT smtp_client.send_email('test subject', 'test body', recipients => array['to@example.com']);
```
