## 用法

来源：

- [官方上游 README](https://github.com/malisper/pgrust/blob/ed269002a1730e18446e716d2e9abbd0a4f00c82/README.md)
- [官方扩展控制文件 (injection_points.control)](https://github.com/malisper/pgrust/blob/ed269002a1730e18446e716d2e9abbd0a4f00c82/crates/contrib/injection_points/extension/injection_points.control)
- [官方扩展 SQL (injection_points--1.0.sql)](https://github.com/malisper/pgrust/blob/ed269002a1730e18446e716d2e9abbd0a4f00c82/crates/contrib/injection_points/extension/injection_points--1.0.sql)

`injection_points` — Postgres 用 Rust 重写，现已通过 100% 的 Postgres 回归测试。当应用程序需要此特定数据库功能时使用它。上游明确表示该项目尚未准备好生产使用。

### 核心工作流

```sql
CREATE EXTENSION injection_points;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `injection_points_attach(IN point_name TEXT, IN action text)` 是一个扩展函数，返回 `void`。
- `injection_points_cached(IN point_name TEXT, IN arg TEXT DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `injection_points_detach(IN point_name TEXT)` 是一个扩展函数，返回 `void`。
- `injection_points_load(IN point_name TEXT)` 是一个扩展函数，返回 `void`。
- `injection_points_run(IN point_name TEXT, IN arg TEXT DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `injection_points_set_local()` 是一个扩展函数，返回 `void`。
- `injection_points_stats_drop()` 是一个扩展函数，返回 `void`。
- `injection_points_stats_fixed(OUT numattach int8, OUT numdetach int8, OUT numrun int8, OUT numcached int8, OUT numloaded int8)` 是一个扩展函数，返回 `record`。
- `injection_points_stats_numcalls(IN point_name TEXT)` 是一个扩展函数，返回 `bigint`。
- `injection_points_wakeup(IN point_name TEXT)` 是一个扩展函数，返回 `void`。
- `removable_cutoff(rel regclass)` 是一个扩展函数，返回 `xid8`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 上游明确表示该项目尚未准备好生产使用。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
