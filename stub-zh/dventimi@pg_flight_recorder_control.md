## 用法

来源：

- [Official database.dev 包页面](https://database.dev/dventimi/pg_flight_recorder_control)

`dventimi@pg_flight_recorder_control` — 控制 pg_flight_recorder 的函数。在收集或解释相应的 PostgreSQL 统计信息时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "dventimi@pg_flight_recorder_control";
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `flight_recorder._get_table_autovacuum_settings` 是一个扩展函数。
- `flight_recorder.compute_recommended_scale_factor(p_relid OID)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder.dead_tuple_growth_rate(p_relid OID, p_window INTERVAL)` 是一个扩展函数并返回 `NUMERIC`。
- `flight_recorder.dead_tuple_trend(p_relid OID, p_window INTERVAL)` 是一个扩展函数并返回 `NUMERIC`。
- `flight_recorder.time_to_budget_exhaustion(p_relid OID, p_budget BIGINT)` 是一个扩展函数并返回 `INTERVAL`。
- `flight_recorder.vacuum_control_mode(p_relid OID)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder.vacuum_control_report(p_start_time TIMESTAMPTZ, p_end_time TIMESTAMPTZ)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder.vacuum_diagnostic(p_relid OID)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.bloat_report` 是一个扩展函数。
- `flight_recorder_reporting.dead_tuple_growth_rate(p_relid OID, p_window INTERVAL)` 是一个扩展函数并返回 `NUMERIC`。
- `flight_recorder_reporting.estimate_table_bloat(p_relid OID DEFAULT NULL)` 是一个扩展函数并返回 `TABLE`。
- `flight_recorder_reporting.oid_consumption_rate(p_window INTERVAL)` 是一个扩展函数并返回 `NUMERIC`。
- `flight_recorder_reporting.table_size_growth_rate(p_relid OID, p_window INTERVAL)` 是一个扩展函数并返回 `NUMERIC`。
- `flight_recorder_reporting.time_to_budget_exhaustion(p_relid OID, p_budget BIGINT)` 是一个扩展函数并返回 `INTERVAL`。

### 要求与注意事项

- 表记录版本为 `2.26.3`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行确认。
