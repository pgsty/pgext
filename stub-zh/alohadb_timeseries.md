## 用法

来源：

- [官方上游 README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [官方扩展控制文件 (alohadb_timeseries.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_timeseries/alohadb_timeseries.control)
- [官方扩展 SQL (alohadb_timeseries--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_timeseries/alohadb_timeseries--1.0.sql)

`alohadb_timeseries` — 使用后台工作者进行时间序列表的自动分区管理。用于相应的调度、时间或时间序列工作流。上游将此功能描述为实验性功能。

### 核心工作流

```sql
CREATE EXTENSION alohadb_timeseries;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `alohadb_time_bucket(bucket_width interval, ts timestamp)` 是一个扩展函数，返回 `timestamp`。
- `alohadb_time_bucket(bucket_width interval, ts timestamptz)` 是一个扩展函数，返回 `timestamptz`。
- `alohadb_timeseries_maintain_now()` 是一个扩展函数，返回 `void`。
- `alohadb_timeseries_manage(p_table_name regclass, p_partition_column text, p_partition_interval interval, p_retention_interval interval DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `alohadb_timeseries_status(OUT table_name text, OUT partition_column text, OUT partition_interval interval, OUT retention_interval interval, OUT partition_count int, OUT enabled boolean)` 是一个扩展函数，返回 `SETOF`。
- `alohadb_timeseries_unmanage(p_table_name regclass)` 是一个扩展函数，返回 `void`。
- `alohadb_timeseries_config` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
