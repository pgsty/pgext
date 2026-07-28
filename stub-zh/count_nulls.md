## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/count_nulls/count_nulls-0.9.7/README.md)
- [官方扩展控制文件 (count_nulls.control)](https://api.pgxn.org/src/count_nulls/count_nulls-0.9.7/count_nulls.control)
- [官方扩展 SQL (count_nulls--0.9.0--0.9.2.sql)](https://api.pgxn.org/src/count_nulls/count_nulls-0.9.7/sql/count_nulls--0.9.0--0.9.2.sql)

`count_nulls` — 确保已经安装了 pg_config 并且在路径中。如果你使用了包管理系统如 RPM 安装 PostgreSQL，请确保也安装了 -devel 包。如果需要，告诉构建过程 pg_config 的位置：使用它来运行 SQL 中需要这些特殊函数或聚合的代码。使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 版本上进行测试。

### 核心工作流

```sql
CREATE EXTENSION count_nulls;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `not_null_count(argument json)` 是一个扩展函数，返回 `int`。
- `not_null_count(argument jsonb)` 是一个扩展函数，返回 `int`。
- `not_null_count(VARIADIC argument anyarray)` 是一个扩展函数，返回 `int`。
- `not_null_count_trigger()` 是一个扩展函数，返回 `trigger`。
- `null_count(argument json)` 是一个扩展函数，返回 `int`。
- `null_count(argument jsonb)` 是一个扩展函数，返回 `int`。
- `null_count(VARIADIC argument anyarray)` 是一个扩展函数，返回 `int`。
- `null_count_trigger()` 是一个扩展函数，返回 `trigger`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.9.6`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
