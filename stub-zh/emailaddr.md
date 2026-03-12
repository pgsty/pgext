

## 用法

> [emailaddr: PostgreSQL 的电子邮件地址数据类型](https://github.com/petere/pgemailaddr)

`emailaddr` 扩展提供了一种数据类型，用于存储和验证符合 RFC 5322 中 `addr-spec` 格式定义的电子邮件地址。

```sql
CREATE EXTENSION emailaddr;

CREATE TABLE accounts (
    id    int PRIMARY KEY,
    name  text,
    email emailaddr
);

INSERT INTO accounts VALUES (1, 'Peter Eisentraut', 'peter@eisentraut.org');
```

### 数据类型

`emailaddr` 类型在输入时按照 RFC 5322 `addr-spec` 规则验证电子邮件地址。接受 `user@domain.com` 等简单格式，不支持 `"User Name" <user@domain.com>` 等显示名称语法。

### 运算符

支持标准比较运算符：`=`、`<>`、`<`、`>`、`<=`、`>=`。

### 索引支持

可使用 Btree 索引对 `emailaddr` 列进行高效查找和排序。
