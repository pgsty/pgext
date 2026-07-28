## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/orafce_mail/orafce_mail-1.2.3/README.md)
- [官方扩展控制文件 (orafce_mail.control)](https://api.pgxn.org/src/orafce_mail/orafce_mail-1.2.3/orafce_mail.control)
- [官方扩展 SQL (orafce_mail--1.2.sql)](https://api.pgxn.org/src/orafce_mail/orafce_mail-1.2.3/orafce_mail--1.2.sql)

`orafce_mail` — 在移植或模拟相应的数据库 API 之前，应先安装 Orafce 扩展。使用此扩展时，请确保其扩展依赖项已安装并验证。

### 核心工作流

```sql
CREATE EXTENSION orafce_mail;

set orafce_mail.smtp_server_url to 'smtps://smtp.gmail.com:465';
set orafce_mail.smtp_server_userpwd to 'pavel.stehule@gmail.com:yourgoogleapppassword';

call utl_mail.send(sender => 'pavel.stehule@gmail.com',
                   recipients => 'pavel.stehule@gmail.com',
                   subject => 'ahoj, nazdar, žlutý kůň',
                   message => e'test, \nžlutý kůň');

do $$
declare
  myimage bytea = (select img from foo limit 1);
begin
  call utl_mail.send_attach_raw(sender => 'pavel.stehule@gmail.com',
                                recipients => 'pavel.stehule@gmail.com',
                                subject => 'mail with picture',
                                message => 'I am sending some picture',
                                attachment => myimage,
                                att_mime_type => 'image/png',
                                att_filename => 'screenshot.png');
end
$$;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `dbms_mail.send` 是一个扩展过程。
- `utl_mail.send` 是一个扩展过程。
- `utl_mail.send_attach_raw` 是一个扩展过程。
- `utl_mail.send_attach_varchar2` 是一个扩展过程。
- `dbms_mail` 是由扩展创建的模式。
- `utl_mail` 是由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.2`。
- 请先安装确认的扩展依赖项：`orafce`。
- 控制文件将该扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
