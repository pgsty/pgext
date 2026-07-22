## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/README)
- [0.1 版本 SQL 对象](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email--0.1.sql)
- [类型实现](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email.c)
- [扩展 control 文件](https://github.com/asotolongo/email/blob/50422767d36208c1252e2696f3ad5525003a0c23/email.control)
- [官方 PGXN 发行页](https://pgxn.org/dist/email/)

`email` 是一种旧式可变长电子邮件地址类型，提供等值、不等值、哈希操作符类，以及拆分本地部分与域名的辅助函数。其验证规则只识别很窄的地址子集，已复核的 C 二进制 I/O 路径也不安全。不能把它用作符合标准的验证器或生产数据类型。

### 核心流程

只能在隔离的兼容性测试中使用文本输入：

```sql
CREATE EXTENSION email;

CREATE TABLE contacts (
  address email NOT NULL
);

INSERT INTO contacts VALUES ('alice@example.com');

SELECT address, getuser(address), getdomain(address)
FROM contacts
WHERE address = 'alice@example.com'::email;
```

C 辅助函数是 `getuser(email)` 和 `getdomain(email)`。SQL 脚本还会安装名为 `get_user(email)` 与 `get_domain(email)` 的 PL/pgSQL 版本。等值比较按字节进行且区分大小写。默认哈希操作符类是 `email_equal_ops`；没有排序操作符或 B-tree 操作符类。

### 验证与安全边界

0.1 版本只接受符合实现正则表达式的小写 ASCII 字符，末级字母后缀长度必须为二至四个字符。它会拒绝大写输入、国际化地址与许多有效的现代域名；语法通过也不能证明邮箱或域名存在。

哈希支持函数是安装脚本中标为 volatile 的临时 PL/pgSQL 计算。更严重的是，C 接收函数会复制到未初始化指针，而发送函数把 varlena 数据当作单个字符读取。因此，二进制协议输入输出与 binary copy 可能破坏内存或导致后端崩溃。应避免二进制 I/O、依赖二进制表示的备份以及真实工作负载，并把数据迁移到 `text` 配合适合应用的验证逻辑。
