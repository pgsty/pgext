## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/json_accessors/json_accessors-1.3.6/README.md)
- [官方扩展控制文件 (json_accessors.control)](https://api.pgxn.org/src/json_accessors/json_accessors-1.3.6/json_accessors.control)
- [官方扩展 SQL (json_accessors.sql)](https://api.pgxn.org/src/json_accessors/json_accessors-1.3.6/sql/json_accessors.sql)

`json_accessors` — JSON 访问器函数 for PostgreSQL ======================================. 用于在 SQL 需要这些特殊函数或聚合时使用。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION json_accessors;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `json_array_to_bigint_array(text)` 是一个扩展函数，返回 `bigint[]`。
- `json_array_to_boolean_array(text)` 是一个扩展函数，返回 `boolean[]`。
- `json_array_to_int_array(text)` 是一个扩展函数，返回 `int[]`。
- `json_array_to_numeric_array(text)` 是一个扩展函数，返回 `numeric[]`。
- `json_array_to_object_array(text)` 是一个扩展函数，返回 `text[]`。
- `json_array_to_text_array(text)` 是一个扩展函数，返回 `text[]`。
- `json_array_to_timestamp_array(text)` 是一个扩展函数，返回 `timestamp`。
- `json_get_bigint(text, text)` 是一个扩展函数，返回 `bigint`。
- `json_get_bigint_array(text, text)` 是一个扩展函数，返回 `bigint[]`。
- `json_get_boolean(text, text)` 是一个扩展函数，返回 `boolean`。
- `json_get_boolean_array(text, text)` 是一个扩展函数，返回 `boolean[]`。
- `json_get_int(text, text)` 是一个扩展函数，返回 `int`。
- `json_get_int_array(text, text)` 是一个扩展函数，返回 `int[]`。
- `json_get_numeric(text, text)` 是一个扩展函数，返回 `numeric`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.3.6`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
