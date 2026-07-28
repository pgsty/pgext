## 用法

来源：

- [官方上游 README](https://github.com/pschlump/pgverhoeff/blob/f433b15dbde363eeea0a918517adef7088a61dec/README.md)
- [官方扩展控制文件 (verhoeff.control)](https://github.com/pschlump/pgverhoeff/blob/f433b15dbde363eeea0a918517adef7088a61dec/verhoeff.control)
- [官方扩展 SQL (verhoeff--1.0.sql)](https://github.com/pschlump/pgverhoeff/blob/f433b15dbde363eeea0a918517adef7088a61dec/verhoeff--1.0.sql)

`verhoeff` — Postgres 提供了强大的可扩展性功能，允许开发者增强其功能。当 SQL 需要这些特殊函数或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION verhoeff;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `generate_verhoeff(inp text)` 是一个扩展函数，返回 `text`。
- `validate_verhoeff(inp text)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
