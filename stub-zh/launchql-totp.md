## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/totp/readme.md)
- [官方扩展控制文件 (launchql-totp.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/totp/launchql-totp.control)
- [官方扩展 SQL (launchql-totp--0.4.5.sql)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/totp/sql/launchql-totp--0.4.5.sql)

`launchql-totp` — TOTP 实现纯 PostgreSQL plpgsql。此扩展提供 RFC 6238 中规定的 HMAC Time-Based One-Time Password Algorithm (TOTP) 作为纯 plpgsql 函数。在实现相应的安全、审计或访问控制工作流时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-totp";

SELECT totp.generate('mysecret');

-- you can also specify totp_interval, and totp_length
SELECT totp.generate('mysecret', 30, 6);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `totp.base32_to_hex(input text)` 是一个扩展函数，返回 `text`。
- `totp.generate(secret text, period int DEFAULT 30, digits int DEFAULT 6, time_from timestamptz DEFAULT now(), hash text DEFAULT 'sha1', encoding text DEFAULT 'base32', clock_offset int DEFAULT 0)` 是一个扩展函数，返回 `text`。
- `totp.generate_secret(hash text DEFAULT 'sha1')` 是一个扩展函数，返回 `bytea`。
- `totp.hotp(key bytea, c int, digits int DEFAULT 6, hash text DEFAULT 'sha1')` 是一个扩展函数，返回 `text`。
- `totp.pad_secret(input bytea, len int)` 是一个扩展函数，返回 `bytea`。
- `totp.random_base32(_length int DEFAULT 20)` 是一个扩展函数，返回 `text`。
- `totp.url(email text, totp_secret text, totp_interval int, totp_issuer text)` 是一个扩展函数，返回 `text`。
- `totp.urlencode(in_str text)` 是一个扩展函数，返回 `text`。
- `totp.verify(secret text, check_totp text, period int DEFAULT 30, digits int DEFAULT 6, time_from timestamptz DEFAULT now(), hash text DEFAULT 'sha1', encoding text DEFAULT 'base32', clock_offset int DEFAULT 0)` 是一个扩展函数，返回 `boolean`。
- `totp` 是由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.4.5`。
- 先安装并验证确认的扩展依赖项：`pgcrypto`, `plpgsql`, `launchql-base32`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
