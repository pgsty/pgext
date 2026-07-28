## 用法

来源：

- [官方上游 README](https://github.com/tursodatabase/pg_turso/blob/a7a9f28176044e49de514f6541461822eefabd99/README.md)
- [官方扩展控制文件 (pg_turso.control)](https://github.com/tursodatabase/pg_turso/blob/a7a9f28176044e49de514f6541461822eefabd99/extension/pg_turso.control)
- [官方扩展 SQL (pg_turso--1.0.sql)](https://github.com/tursodatabase/pg_turso/blob/a7a9f28176044e49de514f6541461822eefabd99/extension/pg_turso--1.0.sql)

`pg_turso` — Postgres 输出插件，用于将数据复制到 Turso。在从 PostgreSQL 移动、转换或集成相应数据时使用它。经过审核的上游材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION pg_turso;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `turso_generate_create_table_for_mv(mv_name text)` 是一个扩展函数，返回 `text`。
- `turso_generate_create_table_for_table(table_name text)` 是一个扩展函数，返回 `text`。
- `turso_migrate_mv_schema(mv_name text)` 是一个扩展函数，返回 `text`。
- `turso_migrate_table_schema(table_name text)` 是一个扩展函数，返回 `text`。
- `turso_schedule_mv_replication(view_name text, refresh_interval text)` 是一个扩展函数，返回 `integer`。
- `turso_schedule_table_replication(table_name text, refresh_interval text)` 是一个扩展函数，返回 `integer`。
- `turso_send(url text, token text, data text)` 是一个扩展函数，返回 `text`。
- `turso_replicate_mv` 是一个扩展过程。
- `turso_replicate_table` 是一个扩展过程。

### 要求与注意事项

- 经过审核的控制文件声明默认版本为 `1.0`。
- 先安装确认的扩展依赖项：`pg_cron`。
- 控制文件将扩展标记为可重定位。
- 上游材料包含显式的弃用边界。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
