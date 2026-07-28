## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg-json/pg-json-0.0.1/README)
- [官方扩展控制文件 (pg-json.control)](https://api.pgxn.org/src/pg-json/pg-json-0.0.1/pg-json.control)
- [官方扩展 SQL (pg-json--0.0.1.sql)](https://api.pgxn.org/src/pg-json/pg-json-0.0.1/pg-json--0.0.1.sql)

`pg-json` — pg-json - PostgreSQL 的 JSON 支持  =====================================。当应用程序数据需要此类型、域或其操作符时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "pg-json";
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `json_equals(this json, that json)` 是一个扩展函数，返回 `boolean`。
- `json_get_value(data json, path text)` 是一个扩展函数，返回 `text`。
- `json_in(cstring)` 是一个扩展函数，返回 `json`。
- `json_not_equals(this json, that json)` 是一个扩展函数，返回 `boolean`。
- `json_out(json)` 是一个扩展函数，返回 `cstring`。
- `json` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
