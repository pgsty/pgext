## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/range_type_functions/range_type_functions-0.0.4/README.md)
- [官方扩展控制文件 (range_type_functions.control)](https://api.pgxn.org/src/range_type_functions/range_type_functions-0.0.4/range_type_functions.control)
- [官方扩展 SQL (range_type_functions.sql)](https://api.pgxn.org/src/range_type_functions/range_type_functions-0.0.4/sql/range_type_functions.sql)
- [当前官方源代码仓库](https://github.com/decibel/range_type_functions)

`range_type_functions` — 此扩展有两个目的：1. 扩展范围函数的功能，旨在将最有用的那些函数移入核心。2. 便于从较新版本的 PostgreSQL 回退函数。当 SQL 需要这些特殊功能或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION range_type_functions;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `element_range_comp(element anyelement, range anyrange)` 是一个扩展函数，返回 `smallint`。
- `get_bound_expr(range anyrange, literal anyelement)` 是一个扩展函数，返回 `text`。
- `get_bounds_condition_expr(range anyrange, text default 'x')` 是一个扩展函数，返回 `text`。
- `get_collation_expr(range anyrange)` 是一个扩展函数，返回 `text`。
- `get_lower_bound_condition_expr(range anyrange, text default 'x')` 是一个扩展函数，返回 `text`。
- `get_subtype_element_expr(range anyrange, text default 'x')` 是一个扩展函数，返回 `text`。
- `get_upper_bound_condition_expr(range anyrange, text default 'x')` 是一个扩展函数，返回 `text`。
- `is_singleton(range anyrange)` 是一个扩展函数，返回 `boolean`。
- `to_range(elem anyelement, range anyrange)` 是一个扩展函数，返回 `anyrange`。
- `to_range(low anyelement, high anyelement, bounds text, range anyrange)` 是一个扩展函数，返回 `anyrange`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.4`。
- 控制文件将扩展标记为可重定位。
- 以前的 `moat` GitHub 地址已不可用；幸存的上游源代码仓库已链接在上文。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码进行比对。
