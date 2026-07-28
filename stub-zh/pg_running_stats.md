## 用法

来源：

- [官方上游 README](https://github.com/chanukyasds/pg_running_stats/blob/20d42698a0c7b594d930a55bfe81ed5c7163a058/README.md)
- [官方扩展控制文件 (pg_running_stats.control)](https://github.com/chanukyasds/pg_running_stats/blob/20d42698a0c7b594d930a55bfe81ed5c7163a058/pg_running_stats.control)
- [官方扩展 SQL (pg_running_stats--1.0.sql)](https://github.com/chanukyasds/pg_running_stats/blob/20d42698a0c7b594d930a55bfe81ed5c7163a058/pg_running_stats--1.0.sql)

`pg_running_stats` — Mergeable Running Statistics (Welford/Chan) for PostgreSQL。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_running_stats;
```

在目标数据库中安装扩展，如果有可用示例，请运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `rstat_combine(a bytea, b bytea)` 是一个扩展函数，返回 `bytea`。
- `rstat_final(state bytea)` 是一个扩展函数，返回 `rstat_result_t`。
- `rstat_sfunc(state bytea, x double precision)` 是一个扩展函数，返回 `bytea`。
- `rstat_state_merge(a bytea, b bytea)` 是一个扩展函数，返回 `bytea`。
- `rstat_state_result(state bytea)` 是一个扩展函数，返回 `rstat_result_t`。
- `rstat_state` 是由扩展公开的聚合。
- `running_stats` 是由扩展公开的聚合。
- `rstat_result_t` 是一个扩展定义的类型。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
