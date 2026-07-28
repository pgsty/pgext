## 用法

来源：

- [官方上游 README](https://github.com/percona/percona_pg_telemetry/blob/6c8c778d7eca74189770bbad5919d5eabc8cb99e/README.md)
- [官方扩展控制文件 (percona_pg_telemetry.control)](https://github.com/percona/percona_pg_telemetry/blob/6c8c778d7eca74189770bbad5919d5eabc8cb99e/percona_pg_telemetry.control)
- [官方扩展 SQL (percona_pg_telemetry--1.0.sql)](https://github.com/percona/percona_pg_telemetry/blob/6c8c778d7eca74189770bbad5919d5eabc8cb99e/percona_pg_telemetry--1.0.sql)

`percona_pg_telemetry` — > [!CAUTION] > 该扩展已被弃用并替换为一个向后兼容的占位符。 > > 将不再收集任何遥测数据，它将不再维护，并且不应在新部署中使用。仅在收集或解释相应的 PostgreSQL 统计信息时使用。上游审查材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION percona_pg_telemetry;

ALTER SYSTEM SET percona_pg_telemetry.enabled = 0;
    SELECT pg_reload_conf();
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `percona_pg_telemetry_status(OUT latest_output_filename text, OUT pt_enabled boolean)` 是一个扩展函数，返回 `record`。
- `percona_pg_telemetry_version()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.2`。
- 控制文件将扩展标记为可重定位。
- 上游材料包含显式的弃用边界。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
