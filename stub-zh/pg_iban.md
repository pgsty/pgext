## 用法

来源：

- [官方上游 README](https://github.com/rom8726/pg_iban/blob/1b4ea9a37823a253512f61803b1429a7037b1099/README.md)
- [官方扩展控制文件 (pg_iban.control)](https://github.com/rom8726/pg_iban/blob/1b4ea9a37823a253512f61803b1429a7037b1099/pg_iban.control)
- [官方扩展 SQL (pg_iban--1.0--1.1.sql)](https://github.com/rom8726/pg_iban/blob/1b4ea9a37823a253512f61803b1429a7037b1099/pg_iban--1.0--1.1.sql)

`pg_iban` — pg_iban 是一个提供国际银行账户号码（IBAN）数据类型以及若干验证和操作 IBAN 的实用函数的 PostgreSQL 扩展。当应用程序需要此类型、域或其操作符时，请使用此扩展。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_iban;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `iban_bban(iban)` 是一个扩展函数，返回 `text`。
- `iban_cmp(iban, iban)` 是一个扩展函数，返回 `int4`。
- `iban_country(iban)` 是一个扩展函数，返回 `text`。
- `iban_eq(iban, iban)` 是一个扩展函数。
- `iban_format(iban)` 是一个扩展函数，返回 `text`。
- `iban_ge(iban, iban)` 是一个扩展函数。
- `iban_gt(iban, iban)` 是一个扩展函数。
- `iban_hash(iban)` 是一个扩展函数，返回 `int4`。
- `iban_in(cstring)` 是一个扩展函数，返回 `iban`。
- `iban_le(iban, iban)` 是一个扩展函数。
- `iban_lt(iban, iban)` 是一个扩展函数。
- `iban_out(iban)` 是一个扩展函数，返回 `cstring`。
- `iban_valid(text)` 是一个扩展函数。
- `iban` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.1`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
