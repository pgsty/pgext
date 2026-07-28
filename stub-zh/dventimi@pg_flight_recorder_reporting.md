## 用法

来源：

- [Official database.dev 包页面](https://database.dev/dventimi/pg_flight_recorder_reporting)

`dventimi@pg_flight_recorder_reporting` — pg_flight_recorder 的报告和分析函数。在收集或解释相应的 PostgreSQL 统计信息时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "dventimi@pg_flight_recorder_reporting";
```

在目标数据库中安装扩展，当可用时运行上方最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `flight_recorder_reporting._diagnose_regression_causes` 是一个扩展函数。
- `flight_recorder_reporting.anomaly_report(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.blast_radius(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.blast_radius_report(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` 是一个扩展函数并返回 `TEXT`。
- `flight_recorder_reporting.capacity_report(p_time_window INTERVAL DEFAULT interval '24 hours')` 是一个扩展函数并返回 `TEXT`。
- `flight_recorder_reporting.capacity_summary(p_time_window INTERVAL DEFAULT interval '24 hours')` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.check_alerts(p_lookback_interval INTERVAL DEFAULT '1 hour')` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.config_at(p_timestamp TIMESTAMPTZ, p_category TEXT DEFAULT NULL)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.config_changes(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.config_health_check()` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.db_role_config_at(p_timestamp TIMESTAMPTZ, p_database TEXT DEFAULT NULL, p_role TEXT DEFAULT NULL, p_prefix TEXT DEFAULT NULL)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.db_role_config_changes(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.db_role_config_summary()` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.detect_query_storms(p_lookback INTERVAL DEFAULT NULL, p_threshold_multiplier NUMERIC DEFAULT NULL)` 是一个扩展函数并返回 `TABLE`。

### 要求与注意事项

- 该目录记录版本 `2.26.3`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
