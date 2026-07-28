## 用法

来源：

- [官方上游 README](https://github.com/andy31002/opentenbase/blob/ff795be78c8583b4129fa9b745597c4fa2e122c8/contrib/README)
- [官方扩展控制文件 (pg_stat_cluster_activity.control)](https://github.com/andy31002/opentenbase/blob/ff795be78c8583b4129fa9b745597c4fa2e122c8/contrib/pg_stat_cluster_activity/pg_stat_cluster_activity.control)
- [官方扩展 SQL (pg_stat_cluster_activity--1.0.sql)](https://github.com/andy31002/opentenbase/blob/ff795be78c8583b4129fa9b745597c4fa2e122c8/contrib/pg_stat_cluster_activity/pg_stat_cluster_activity--1.0.sql)

`pg_stat_cluster_activity` — 用于在整个集群范围内跟踪执行统计信息。在收集或解释相应的 PostgreSQL 统计信息时使用此功能。上游将此功能描述为实验性功能。

### 核心工作流

```sql
CREATE EXTENSION pg_stat_cluster_activity;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_cancel_session(text)` 是一个扩展函数。
- `pg_signal_session(text, integer, bool)` 是一个扩展函数。
- `pg_terminate_session(text)` 是一个扩展函数。
- `pg_stat_cluster_activity` 是一个扩展定义视图。
- `pg_stat_cluster_activity_cn` 是一个扩展定义视图。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
