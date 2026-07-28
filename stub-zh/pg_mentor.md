## 用法

来源：

- [官方上游 README](https://github.com/danolivo/pg_mentor/blob/cd7177d756dbfa71f0ccf9a043e3b99685f4a887/README.md)
- [官方扩展控制文件 (pg_mentor.control)](https://github.com/danolivo/pg_mentor/blob/cd7177d756dbfa71f0ccf9a043e3b99685f4a887/pg_mentor.control)
- [官方扩展 SQL (pg_mentor--0.1.sql)](https://github.com/danolivo/pg_mentor/blob/cd7177d756dbfa71f0ccf9a043e3b99685f4a887/pg_mentor--0.1.sql)

`pg_mentor` — 一个轻量级扩展，利用 pg_stat_statements 扩展中存储的查询统计信息来决定使用哪种类型的计划模式（自定义、通用或自动）。在管理或自动化上述数据库行为时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_mentor;

SELECT reconsider_ps_modes();
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_mentor_nail_long_planned()` 是一个扩展函数，返回 `integer`。
- `pg_mentor_reload_conf(void)` 是一个扩展函数。
- `pg_mentor_reset()` 是一个扩展函数，返回 `integer`。
- `pg_mentor_set_plan_mode(queryId bigint, status integer, ref_total_time float8 DEFAULT NULL, ref_nblocks float8 DEFAULT NULL, fixed bool DEFAULT false)` 是一个扩展函数。
- `pg_mentor_show_prepared_statements(IN status integer, OUT queryid bigint, OUT refcounter integer, OUT plan_cache_mode int, OUT since TimestampTz, OUT fixed boolean, OUT statnum integer, OUT nblocks bigint[], OUT exec_times float8[], OUT avg_nblocks float8, OUT avg_exec_time float8, OUT ref_nblo…)` 是一个扩展函数，返回 `SETOF`。
- `reconsider_ps_modes(OUT to_generic bigint, OUT to_custom bigint, OUT unchanged bigint)` 是一个扩展函数，返回 `record`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
