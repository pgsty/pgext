

## 用法

> [chkpass: 自动加密的密码数据类型](https://github.com/lacanoid/chkpass)

`chkpass` 扩展提供了一种用于存储加密密码的数据类型。该类型最初随 PostgreSQL 一起发布（在 PG 11 中移除），这是适用于现代 PostgreSQL 的独立版本。

```sql
CREATE EXTENSION chkpass;
```

### 数据类型

`chkpass` 类型在输入时使用 Unix `crypt()` 自动加密密码，仅存储加密后的形式。

```sql
CREATE TABLE accounts (
    username text PRIMARY KEY,
    password chkpass
);

INSERT INTO accounts VALUES ('admin', 'mysecretpassword');
```

### 运算符

`=` 运算符用于将明文密码与存储的加密值进行比对：

```sql
SELECT * FROM accounts WHERE password = 'mysecretpassword';
-- 如果密码正确则返回匹配行
```

### 行为特性

- 密码在输入时自动加密——明文永远不会被存储
- 输出显示加密后的哈希值，而非原始密码
- `=` 比较会先加密右侧操作数，然后比较哈希值
- 使用标准 Unix `crypt()` 函数进行加密
