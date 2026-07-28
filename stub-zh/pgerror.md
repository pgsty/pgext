## 用法

来源：

- [官方扩展控制文件 (pgerror.control)](https://api.pgxn.org/src/pgerror/pgerror-0.2.1/pgerror.control)
- [官方扩展 SQL (pgerror.in.sql)](https://api.pgxn.org/src/pgerror/pgerror-0.2.1/sql/pgerror.in.sql)

`pgerror` — 工具用于更好的错误处理。在相应的 SQL 或数据库实用程序工作流中使用它。使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgerror;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `error_data(sqlstate text = '' , message text = '' , hint text = '' , detail text = '' , context text = '' , schema_name text = '' , table_name text = '' , column_name text = '' , constraint_name text = '' , type_name text = '')` 是一个扩展函数，返回 `error_data`。
- `pg_temp.raise_error_internal()` 是一个扩展函数，返回 `void`。
- `raise(error error_data , level text = 'EXCEPTION')` 是一个扩展函数，返回 `void`。
- `raise(message text , level text = 'EXCEPTION' , sqlstate text = NULL , hint text = NULL , detail text = NULL , schema_name text = NULL , table_name text = NULL , column_name text = NULL , constraint_name text = NULL , type_name text = NULL)` 是一个扩展函数，返回 `void`。
- `try(code text , OUT row_count int , OUT error error_data)` 是一个扩展函数，返回 `record`。
- `try_cursor(query text , cursor_name text DEFAULT NULL , OUT result refcursor , OUT error error_data)` 是一个扩展函数，返回 `record`。
- `try_into(code text , INOUT result anyelement , strict boolean DEFAULT false , OUT error error_data)` 是一个扩展函数，返回 `record`。
- `error_data` 是一个扩展定义的类型。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.2.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
