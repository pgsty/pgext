## 用法

来源：

- [官方 README](https://github.com/asio/password/blob/master/README)
- [扩展 SQL](https://github.com/asio/password/blob/master/sql/password--0.0.1.sql)
- [C 实现](https://github.com/asio/password/blob/master/src/password.c)

`password` 0.0.1 是一个过时的自定义密码类型原型，内部捆绑 Blowfish crypt 代码。源码审查发现其哈希、typmod 和比较路径存在多处内存安全与类型处理缺陷。不要将它用于认证，也不要用它保存真实凭据。

### 预期 SQL 接口

下面只为审计与兼容性分析展示上游接口，并非生产建议：

```sql
CREATE EXTENSION password;

CREATE TABLE password_demo (
  account text PRIMARY KEY,
  secret password
);
```

该类型意图使用 `$2a$` Blowfish 设置对明文输入进行哈希，把已经以 `$2a$` 开头的字符串当作已存储哈希，并提供隐式文本转换。其 `=` 操作符把左操作数当作已存储密码、右操作数当作输入明文；这不是通常的对称相等语义。

### 安全与正确性风险

C 源码中的哈希辅助函数返回指向栈内存的指针，typmod 处理中错误使用整数指针，一个比较路径只复制指针大小的字节，而且没有一致地保留或检查密码学函数错误。这些缺陷可能导致崩溃、错误比较或未定义行为。比较过程也没有恒定时间保证，捆绑实现针对的是旧 `$2a$` 变体。

### 更安全的边界

应在持续维护的认证组件中验证密码；若有明确安全设计，也只能考虑当前受支持且经过独立审查的 PostgreSQL 机制，例如 `pgcrypto`。已有数据库若包含该自定义类型，应把它视为迁移与事件审查问题：仅在受控验证后导出哈希、撤销访问并退役扩展，而不能继续信任其比较结果。
