## 用法

来源：

- [PGXN 发行版 1.0.0 README](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/README.md)
- [版本 0.0.1 安装 SQL](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/sql/pg_datatype_password--0.0.1.sql)
- [扩展控制文件](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/pg_datatype_password.control)

`pg_datatype_password` 提供 `password` 类型，以及由 `pgcrypto` Blowfish 哈希支持的明文比较操作符。本次核对的 PGXN 发行版是 `1.0.0`，但其中可安装的扩展版本是 `0.0.1`。

### 核心流程

该类型的输入函数不会进行哈希。每张存储此类型的表都必须使用 `t_encrypt_password` BEFORE 触发器，否则会直接存入明文。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_datatype_password;

CREATE TABLE app_user (
  username text PRIMARY KEY,
  passwd password NOT NULL
);

CREATE TRIGGER app_user_encrypt_password
BEFORE INSERT OR UPDATE OF passwd ON app_user
FOR EACH ROW EXECUTE PROCEDURE t_encrypt_password('passwd', '10');

INSERT INTO app_user VALUES ('alice', 'secret');
SELECT username FROM app_user WHERE passwd = 'secret'::text;
```

可选的第二个触发器参数是 bcrypt 成本；上游默认值为 `8`。输出时显示存储的哈希。更新密码列会对新值进行哈希，因此写入已有哈希会再次哈希。

### 安全与注意事项

可用比较是在 `password` 与 `text` 之间双向使用 `=` 和 `<>`；没有文档说明排序或密码索引操作符类。应限制对表的直接访问，避免记录明文，并在认证外围采用当前的密码策略和速率限制。

尽管控制文件声明可重定位，版本化 SQL 仍将多个对象硬编码到 `public`。请安装并测试于预期模式。这是一个较老的 bcrypt 设计，默认工作因子已经过时；应评估使用受维护库在应用侧哈希密码是否更符合威胁模型。
