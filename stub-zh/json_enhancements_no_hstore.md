## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/README.md)
- [官方扩展控制文件 (json_enhancements_no_hstore.control)](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/json_enhancements_no_hstore.control)
- [官方扩展 SQL (json_enhancements_no_hstore.sql)](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/sql/json_enhancements_no_hstore.sql)

`json_enhancements_no_hstore` — Json 增强功能 for PostgreSQL 9.2 ====================================。当应用程序需要此特定数据库功能时使用它。审核的上游项目已归档或不再维护。

### 核心工作流

```sql
CREATE EXTENSION json_enhancements_no_hstore;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `json_agg_finalfn(internal)` 是一个扩展函数，返回 `json`。
- `json_agg_transfn(internal, anyelement)` 是一个扩展函数，返回 `internal`。
- `json_array_element(the_json json, element integer)` 是一个扩展函数，返回 `json`。
- `json_array_element_text(the_json json, element integer)` 是一个扩展函数，返回 `text`。
- `json_array_elements(the_json json)` 是一个扩展函数，返回 `TABLE`。
- `json_array_length(the_json json)` 是一个扩展函数，返回 `int`。
- `json_each(the_json json, key out text, out value json)` 是一个扩展函数，返回 `SETOF record`。
- `json_each_text(the_json json, key out text, value out text)` 是一个扩展函数，返回 `SETOF record`。
- `json_extract_path(the_json json, variadic path_elements text[])` 是一个扩展函数，返回 `json`。
- `json_extract_path_op(the_json json, path_elements text[])` 是一个扩展函数，返回 `json`。
- `json_extract_path_text(the_json json, variadic path_elements text[])` 是一个扩展函数，返回 `text`。
- `json_extract_path_text_op(the_json json, path_elements text[])` 是一个扩展函数，返回 `text`。
- `json_object_field(json, text)` 是一个扩展函数，返回 `json`。
- `json_object_field_text(json, text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
