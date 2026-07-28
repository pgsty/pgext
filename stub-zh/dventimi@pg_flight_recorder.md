## 用法

来源：

- [Official database.dev 包页面](https://database.dev/dventimi/pg_flight_recorder)

`dventimi@pg_flight_recorder` — PostgreSQL 性能监控的服务器端飞行记录器。在收集或解释相应的 PostgreSQL 统计信息时使用它。上游审查的材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION "dventimi@pg_flight_recorder";
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `flight_recorder._check_and_adjust_mode` 是一个扩展函数。
- `flight_recorder._check_catalog_ddl_locks` 是一个扩展函数。
- `flight_recorder._check_circuit_breaker` 是一个扩展函数。
- `flight_recorder._check_schema_size` 是一个扩展函数。
- `flight_recorder._check_statements_health` 是一个扩展函数。
- `flight_recorder._collect_config_snapshot` 是一个扩展函数。
- `flight_recorder._collect_db_role_config_snapshot` 是一个扩展函数。
- `flight_recorder._collect_index_stats` 是一个扩展函数。
- `flight_recorder._collect_table_stats` 是一个扩展函数。
- `flight_recorder._get_config` 是一个扩展函数。
- `flight_recorder._get_ring_buffer_slots` 是一个扩展函数。
- `flight_recorder._get_ring_retention_interval` 是一个扩展函数。
- `flight_recorder._get_setting_from_snapshots` 是一个扩展函数。
- `flight_recorder._has_pg_stat_statements` 是一个扩展函数。

### 要求与注意事项

- 表记录版本为 `2.26.3`。
- 首先安装确认的扩展依赖项：`pg_cron`。
- 上游材料包含显式的弃用边界。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
