## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/check_digits/check_digits-0.3.3/README.md)
- [官方扩展控制文件 (check_digits.control)](https://api.pgxn.org/src/check_digits/check_digits-0.3.3/check_digits.control)
- [官方扩展 SQL (check_digits--0.3.2.sql)](https://api.pgxn.org/src/check_digits/check_digits-0.3.3/sql/check_digits--0.3.2.sql)

`check_digits` — 最后，在 psql 中创建扩展。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION check_digits;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `check_digits_inn(inn text)` 是一个扩展函数，返回 `boolean`。
- `check_digits_isbn(isbn text)` 是一个扩展函数，返回 `boolean`。
- `check_digits_ogrn(ogrn text)` 是一个扩展函数，返回 `boolean`。
- `check_digits_okpo(okpo text)` 是一个扩展函数，返回 `boolean`。
- `check_digits_snils(snils text)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.3.3`。
- 控制文件将扩展标记为可重定位。
- 2026-07-28 审查期间，前 GitHub 仓库 URL 返回 404；请将上述固定 PGXN 发行版视为可用源边界。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
