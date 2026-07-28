## 用法

来源：

- [官方上游 README](https://github.com/shicheng0104/opengauss/blob/0495a2328db5f409b6e29eda1d671964f35168d4/contrib/README)
- [官方扩展控制文件 (shark.control)](https://github.com/shicheng0104/opengauss/blob/0495a2328db5f409b6e29eda1d671964f35168d4/contrib/shark/shark.control)
- [官方扩展 SQL (shark--1.0.sql)](https://github.com/shicheng0104/opengauss/blob/0495a2328db5f409b6e29eda1d671964f35168d4/contrib/shark/shark--1.0.sql)

`shark` — 扩展用于 D 语言兼容性。在移植或模拟相应数据库 API 时使用。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION shark;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `dbcc_check_ident_no_reseed(varchar, boolean, boolean)` 是一个扩展函数，返回 `varchar`。
- `dbcc_check_ident_reseed(varchar, int16, boolean)` 是一个扩展函数，返回 `varchar`。
- `fetch_status()` 是一个扩展函数，返回 `int`。
- `objectproperty(id INT, property VARCHAR)` 是一个扩展函数，返回 `INT`。
- `rowcount()` 是一个扩展函数，返回 `int`。
- `rowcount_big()` 是一个扩展函数，返回 `bigint`。
- `spid()` 是一个扩展函数，返回 `bigint`。
- `sys.day(abstime)` 是一个扩展函数，返回 `float8`。
- `sys.day(date)` 是一个扩展函数，返回 `float8`。
- `sys.day(timestamp(0) with time zone)` 是一个扩展函数，返回 `float8`。
- `sys.day(timestamptz)` 是一个扩展函数，返回 `float8`。
- `sys.object_id(IN object_name VARCHAR, IN object_type VARCHAR DEFAULT '')` 是一个扩展函数，返回 `integer`。
- `sys.pltsql_call_handler()` 是一个扩展函数，返回 `language_handler`。
- `sys.pltsql_inline_handler(internal)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `2.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
