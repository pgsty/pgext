## 用法

来源：

- [官方上游 README](https://github.com/pyramation/totp/blob/8f724aa2ba53abe49272cbdc90c2004df848b8ea/extensions/@launchql/base32/readme.md)
- [官方扩展控制文件 (launchql-base32.control)](https://github.com/pyramation/totp/blob/8f724aa2ba53abe49272cbdc90c2004df848b8ea/extensions/@launchql/base32/launchql-base32.control)
- [官方扩展 SQL (launchql-base32--0.0.3.sql)](https://github.com/pyramation/totp/blob/8f724aa2ba53abe49272cbdc90c2004df848b8ea/extensions/@launchql/base32/sql/launchql-base32--0.0.3.sql)

`launchql-base32` — 首先，你需要启动 postgres 容器（你也可以直接使用 docker-compose up -d）。在实现相应的安全、审计或访问控制工作流时，请使用它。安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-base32";

select base32.encode('foo');
-- MZXW6===


select base32.decode('MZXW6===');
-- foo
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `base32.base32_alphabet(input int)` 是一个扩展函数，返回 `char`。
- `base32.base32_alphabet_to_decimal(input text)` 是一个扩展函数，返回 `text`。
- `base32.base32_alphabet_to_decimal_int(input text)` 是一个扩展函数，返回 `int`。
- `base32.base32_to_decimal(input text)` 是一个扩展函数，返回 `text[]`。
- `base32.binary_to_int(input text)` 是一个扩展函数，返回 `int`。
- `base32.decimal_to_chunks(input text[])` 是一个扩展函数，返回 `text[]`。
- `base32.decode(input text)` 是一个扩展函数，返回 `text`。
- `base32.encode(input text)` 是一个扩展函数，返回 `text`。
- `base32.fill_chunks(input text[])` 是一个扩展函数，返回 `text[]`。
- `base32.string_nchars(text, int)` 是一个扩展函数，返回 `text[]`。
- `base32.to_ascii(input text)` 是一个扩展函数，返回 `int[]`。
- `base32.to_base32(input text[])` 是一个扩展函数，返回 `text`。
- `base32.to_binary(input int)` 是一个扩展函数，返回 `text`。
- `base32.to_binary(input int[])` 是一个扩展函数，返回 `text[]`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.3`。
- 先安装并验证确认的扩展依赖项：`pgcrypto`, `plpgsql`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
