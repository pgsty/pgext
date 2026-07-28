## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [官方扩展控制文件 (gp_toolkit.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_toolkit/gp_toolkit.control)
- [官方扩展 SQL (gp_toolkit--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_toolkit/gp_toolkit--1.0.sql)

`gp_toolkit` — 用于管理 Greenplum 家族数据库的管理视图和函数。在进行数据库管理或自动化上述描述的行为时使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_toolkit;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gp_toolkit.get_column_size(ao_oid oid, OUT segment int, OUT attnum int, OUT size bigint, OUT size_uncompressed bigint, OUT compression_ratio numeric)` 是一个扩展函数，返回 `SETOF`。
- `gp_toolkit.gp_param_setting(varchar)` 是一个扩展函数，返回 `SETOF`。
- `gp_toolkit.gp_param_settings()` 是一个扩展函数，返回 `SETOF`。
- `gp_toolkit.gp_skew_coefficient(targetoid oid, OUT skcoid oid, OUT skccoeff numeric)` 是一个扩展函数，返回 `record`。
- `gp_toolkit.gp_skew_details(oid)` 是一个扩展函数，返回 `setof`。
- `gp_toolkit.gp_skew_idle_fraction(targetoid oid, OUT sifoid oid, OUT siffraction numeric)` 是一个扩展函数，返回 `record`。
- `gp_toolkit.session_state_memory_entries_f_on_master()` 是一个扩展函数，返回 `SETOF`。
- `gp_toolkit.session_state_memory_entries_f_on_segments()` 是一个扩展函数，返回 `SETOF`。
- `gp_toolkit.gp_param_setting_t` 是一个扩展定义的类型。
- `gp_toolkit.gp_skew_analysis_t` 是一个扩展定义的类型。
- `gp_toolkit.gp_skew_details_t` 是一个扩展定义的类型。
- `gp_toolkit.gp_bloat_diag` 是一个扩展定义的视图。
- `gp_toolkit.gp_bloat_expected_pages` 是一个扩展定义的视图。
- `gp_toolkit.gp_check_missing_files` 是一个扩展定义的视图。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.6`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
