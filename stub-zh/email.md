## 用法

来源：

- [上游 README](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/README)
- [扩展 SQL](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email--0.1.sql)
- [C 实现](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email.c)
- [PGXN 发行包](https://pgxn.org/dist/email/)

`email` 增加一个变长电子邮件地址类型，以及 `getuser(email)` 与 `getdomain(email)` 辅助函数。扩展的相等运算符按字节比较地址：

```sql
CREATE EXTENSION email;

CREATE TABLE contacts (address email NOT NULL);
INSERT INTO contacts VALUES ('alice@example.com');

SELECT address, getuser(address), getdomain(address)
FROM contacts
WHERE address = 'alice@example.com'::email;
```

### 校验与索引限制

扩展版本 `0.1` 只接受实现中严格限定的小写 ASCII 模式，其中末级域名后缀必须为两到四个字母；大写地址和许多现代合法地址形式都会被拒绝。PGXN 包标记为 `0.1.0`，而 control 与安装 SQL 公开的版本为 `0.1`。上游警告哈希支持仍是临时实现：`email_hash` 使用 PL/pgSQL 编写，且没有定义排序运算符或 B-tree 运算符类。采用这一旧式类型前，应验证兼容性与校验规则。
