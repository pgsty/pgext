## 用法

来源：

- [官方上游 README](https://github.com/cybertec-postgresql/postgresql_extension_template/blob/8dcf2fa61c23185e2dedec0bc21b1402ff109def/README.md)
- [官方扩展控制文件 (postgresql_extension_template.control)](https://github.com/cybertec-postgresql/postgresql_extension_template/blob/8dcf2fa61c23185e2dedec0bc21b1402ff109def/postgresql_extension_template.control)
- [官方扩展 SQL (postgresql_extension_template--1.0.sql)](https://github.com/cybertec-postgresql/postgresql_extension_template/blob/8dcf2fa61c23185e2dedec0bc21b1402ff109def/postgresql_extension_template--1.0.sql)

`postgresql_extension_template` — 这是一个 PostgreSQL C 扩展的模板仓库。该仓库包括：在需要此特定数据库功能的应用程序中使用它。使用上述链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION postgresql_extension_template;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `my_function()` 是一个扩展函数，返回 `cstring`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
