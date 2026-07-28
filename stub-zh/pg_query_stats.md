## 用法

来源：

- [官方上游 README](https://github.com/emrullahayaz9/pg_query_stats/blob/0d67af30c513d27e250a85cfd262268001fd9b1a/README.md)
- [官方扩展控制文件 (pg_query_stats.control)](https://github.com/emrullahayaz9/pg_query_stats/blob/0d67af30c513d27e250a85cfd262268001fd9b1a/pg_query_stats.control)
- [官方扩展 SQL (pg_query_stats--1.0.0.sql)](https://github.com/emrullahayaz9/pg_query_stats/blob/0d67af30c513d27e250a85cfd262268001fd9b1a/pg_query_stats--1.0.0.sql)

`pg_query_stats` — **pg_query_stats** 是一个轻量级的 PostgreSQL 扩展，用于收集并暴露 SQL 查询的执行统计信息。它帮助开发人员和 DBA 以最小的开销和简单的架构监控查询性能。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_query_stats;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_query_stats()` 是一个扩展函数，返回 `SETOF`。
- `pg_query_stats_reset()` 是一个扩展函数，返回 `void`。
- `pg_query_stats` 是一个扩展定义的视图。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
