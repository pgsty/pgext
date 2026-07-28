## 用法

来源：

- [官方上游 README](https://github.com/aau-daisy/solvedb/blob/8a15559b5a08747b2217d9146e3dd122379d4de1/README.md)
- [官方扩展控制文件 (solverapi.control)](https://github.com/aau-daisy/solvedb/blob/8a15559b5a08747b2217d9146e3dd122379d4de1/SolverAPI/solverapi.control)
- [官方扩展 SQL (solverapi--1.2.sql)](https://github.com/aau-daisy/solvedb/blob/8a15559b5a08747b2217d9146e3dd122379d4de1/SolverAPI/solverapi--1.2.sql)

`solverapi` — SolveDB：一个基于 PostgreSQL 的数据库管理系统，适用于优化应用。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION solverapi;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `sl_build_dst_ctr(arg sl_solver_arg, vsout sl_viewsql_out, ctr_nr int)` 是一个扩展函数，返回 `sl_viewsql_dst`。
- `sl_build_dst_ctr_union(arg sl_solver_arg, vsout sl_viewsql_out, ctr_type text)` 是一个扩展函数，返回 `sl_viewsql_dst`。
- `sl_build_dst_obj(arg sl_solver_arg, vsout sl_viewsql_out)` 是一个扩展函数，返回 `sl_viewsql_dst`。
- `sl_build_dst_values(arg sl_solver_arg, vsout sl_viewsql_out, cast_to text DEFAULT 'text')` 是一个扩展函数，返回 `sl_viewsql_dst`。
- `sl_build_out(arg sl_solver_arg)` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_array1subst(arg sl_solver_arg, par_nr int DEFAULT 1)` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_arrayNsubst(arg sl_solver_arg, par_pos int[])` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_defcols(arg sl_solver_arg, colvalues text[][], base sl_viewsql_out DEFAULT NULL)` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_func1subst(arg sl_solver_arg, func text)` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_funcNmap(arg sl_solver_arg, base sl_viewsql_out, funcs text[])` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_funcNsubst(arg sl_solver_arg, funcs text[])` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_join(arg sl_solver_arg, base sl_viewsql_out, sql text, join_id_col text)` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_joinvalues(arg sl_solver_arg, sql text, col_varnr text, col_value text)` 是一个扩展函数，返回 `sl_viewsql_out`。
- `sl_build_out_rename(arg sl_solver_arg, base sl_viewsql_out, col_type sl_attribute_kind, col_alias text)` 是一个扩展函数，返回 `sl_viewsql_out`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.2`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
