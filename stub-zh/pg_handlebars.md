## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_handlebars/pg_handlebars-1.0.7/README.md)
- [官方扩展控制文件 (pg_handlebars.control)](https://api.pgxn.org/src/pg_handlebars/pg_handlebars-1.0.7/pg_handlebars.control)
- [官方扩展 SQL (pg_handlebars--1.0.sql)](https://api.pgxn.org/src/pg_handlebars/pg_handlebars-1.0.7/pg_handlebars--1.0.sql)

`pg_handlebars` — PostgreSQL 实现的 handlebars 模板引擎。用于相应的 SQL 或数据库实用程序工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_handlebars;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `handlebars(json JSON, template TEXT)` 是一个扩展函数，返回 `TEXT`。
- `handlebars(json JSON, template TEXT, file TEXT)` 是一个扩展函数，返回 `BOOL`。
- `handlebars_compiler_flag_all()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_alternate_decorators()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_assume_objects()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_compat()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_explicit_partial_context()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_ignore_standalone()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_known_helpers_only()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_mustache_style_lambdas()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_no_escape()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_none()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_prevent_indent()` 是一个扩展函数，返回 `void`。
- `handlebars_compiler_flag_strict()` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
