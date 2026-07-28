## 用法

来源：

- [官方扩展控制文件 (json_model.control)](https://api.pgxn.org/src/json_model/json_model-2.0.0-alpha3/json_model.control)
- [官方扩展 SQL (json_model--2.0.sql)](https://api.pgxn.org/src/json_model/json_model-2.0.0-alpha3/json_model--2.0.sql)

`json_model` — JSON Model PL/pgSQL 运行时 - JSON 值验证。当应用程序需要此特定数据库功能时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION json_model;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `jm_array_is_unique(val JSONB, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_call(fun TEXT, val JSONB, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_check_constraint(val JSONB, op TEXT, cst ANYELEMENT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_cmap_get(TEXT, JSONB)` 是一个扩展函数，返回 `TEXT`。
- `jm_is_valid_date(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_datetime(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_email(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_extreg(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_regex(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_time(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_url(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_is_valid_uuid(val TEXT, path TEXT[], rep jm_report_entry[])` 是一个扩展函数，返回 `BOOLEAN`。
- `jm_object_size(val JSONB)` 是一个扩展函数，返回 `INT`。
- `jm_report_add_entry` 是一个扩展存储过程。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `2.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码进行比对。
