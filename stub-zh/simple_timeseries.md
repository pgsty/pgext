## 用法

来源：

- [官方上游 README](https://github.com/gaoweichang/timeseries-extension/blob/7b3d2ca3c31c2cad61ea09712c17874ecad28d87/README.md)
- [官方扩展控制文件 (simple_timeseries.control)](https://github.com/gaoweichang/timeseries-extension/blob/7b3d2ca3c31c2cad61ea09712c17874ecad28d87/project/simple_timeseries.control)
- [官方扩展 SQL (simple_timeseries--1.0.sql)](https://github.com/gaoweichang/timeseries-extension/blob/7b3d2ca3c31c2cad61ea09712c17874ecad28d87/project/sql/simple_timeseries--1.0.sql)

`simple_timeseries` — 时间序列扩展。用于相应的调度、时间或时间序列工作流。在安装该扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION simple_timeseries;

CREATE TABLE sensor_data (
    time TIMESTAMPTZ NOT NULL,
    sensor_id INTEGER,
    temperature DOUBLE PRECISION,
    humidity DOUBLE PRECISION
);

SELECT create_hypertable('sensor_data', 'time', INTERVAL '1 day');
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `apply_retention_policies()` 是一个扩展函数，返回 `VOID`。
- `check_hypertable_exists(schema_name text, table_name text)` 是一个扩展函数，返回 `boolean`。
- `compress_chunk(chunk_name REGCLASS)` 是一个扩展函数，返回 `VOID`。
- `create_continuous_aggregate(view_name TEXT, hypertable REGCLASS, view_sql TEXT, bucket_width INTERVAL, refresh_interval INTERVAL DEFAULT NULL)` 是一个扩展函数，返回 `VOID`。
- `create_hypertable(table_name REGCLASS, time_column_name TEXT, chunk_time_interval INTERVAL)` 是一个扩展函数，返回 `VOID`。
- `display_all_chunks()` 是一个扩展函数，返回 `TABLE`。
- `display_all_triggers()` 是一个扩展函数，返回 `TABLE`。
- `drop_chunks(hypertable REGCLASS, older_than INTERVAL)` 是一个扩展函数，返回 `INTEGER`。
- `drop_continuous_aggregate(view_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `drop_hypertable(table_name REGCLASS)` 是一个扩展函数，返回 `VOID`。
- `get_chunks_by_hid(p_hypertable_id int)` 是一个扩展函数，返回 `TABLE`。
- `get_current_timestamp_seconds()` 是一个扩展函数，返回 `bigint`。
- `get_hypertable_info(s_name text, t_name text)` 是一个扩展函数，返回 `text`。
- `refresh_continuous_aggregate(view_name TEXT, start_time TIMESTAMPTZ, end_time TIMESTAMPTZ)` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0`。
- 先安装并验证确认的扩展依赖项：`dblink`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
