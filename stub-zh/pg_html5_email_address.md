

## 用法

> [pg_html5_email_address: PostgreSQL 的 HTML5 电子邮件地址验证](https://github.com/bigsmoke/pg_html5_email_address)

提供与 HTML5 `<input type="email">` 规范一致的电子邮件地址验证。

### 域类型：`html5_email`

一个强制执行 HTML5 电子邮件验证规则的域类型，支持大小写不敏感比较：

```sql
SELECT 'user@example.com'::html5_email;

-- 大小写不敏感的相等比较：
SELECT 'User@Example.com'::html5_email = 'user@example.com'::html5_email;
-- t

-- 无效的电子邮件地址会触发 check_violation 错误：
SELECT 'user @example.com'::html5_email;
-- ERROR: check_violation
```

### 函数：`html5_email_regexp()`

返回匹配有效 HTML5 电子邮件地址的正则表达式：

```sql
-- 检查字符串是否为有效的 HTML5 电子邮件地址
SELECT 'user@example.com' ~ html5_email_regexp();
-- t

SELECT 'user @example.com' ~ html5_email_regexp();
-- f
```

使用可选的捕获组获取本地部分和域名部分：

```sql
SELECT (regexp_matches('user@example.com', html5_email_regexp(true)))[1];
-- 'user'
SELECT (regexp_matches('user@example.com', html5_email_regexp(true)))[2];
-- 'example.com'
```

### 验证规则

- 不允许包含空格
- 不允许使用非 ASCII 字符（本地部分和域名部分均不允许）
- `@` 后面必须有内容
- 本地部分允许使用特殊字符，如 `!#$%&'*+/=?^_`{|}~-`
