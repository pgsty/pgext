## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_utility_trigger_functions/pg_utility_trigger_functions-1.9.3/README.md)
- [官方扩展控制文件 (pg_utility_trigger_functions.control)](https://api.pgxn.org/src/pg_utility_trigger_functions/pg_utility_trigger_functions-1.9.3/pg_utility_trigger_functions.control)
- [官方扩展 SQL (pg_utility_trigger_functions--1.0.0.sql)](https://api.pgxn.org/src/pg_utility_trigger_functions/pg_utility_trigger_functions-1.9.3/sql/pg_utility_trigger_functions--1.0.0.sql)

`pg_utility_trigger_functions` — The pg_utility_trigger_functions PostgreSQL 扩展将一些作者 BigSmoke 喜欢在各种 PostgreSQL 项目中使用的宠物触发函数打包在一起。使用它来进行相应的 SQL 或数据库实用工具工作流。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_utility_trigger_functions;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `coalesce_sibling_fields()` 是一个扩展函数，返回 `trigger`。
- `copy_fields_from_foreign_table()` 是一个扩展函数，返回 `trigger`。
- `fallback_to_fields_from_foreign_table()` 是一个扩展函数，返回 `trigger`。
- `no_delete()` 是一个扩展函数，返回 `trigger`。
- `nullify_columns()` 是一个扩展函数，返回 `trigger`。
- `overwrite_composite_field_in_referencing_table()` 是一个扩展函数，返回 `trigger`。
- `overwrite_fields_in_referencing_table()` 是一个扩展函数，返回 `trigger`。
- `pg_utility_trigger_functions_meta_pgxn()` 是一个扩展函数，返回 `jsonb`。
- `pg_utility_trigger_functions_readme()` 是一个扩展函数，返回 `text`。
- `set_installed_extension_version_from_name()` 是一个扩展函数，返回 `trigger`。
- `test__mock.now()` 是一个扩展函数，返回 `timestamptz`。
- `update_updated_at()` 是一个扩展函数，返回 `trigger`。
- `test__coalesce_sibling_fields` 是一个扩展过程。
- `test__copy_fields_from_foreign_table` 是一个扩展过程。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.9.3`。
- 先安装并验证确认的扩展依赖项：`hstore`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
