## 用法

来源：

- [官方上游 README](https://github.com/animeshs34/timeql/blob/a3032739628758757e223d49fc67b276c6e14f8c/README.md)
- [官方扩展控制文件 (timeql.control)](https://github.com/animeshs34/timeql/blob/a3032739628758757e223d49fc67b276c6e14f8c/timeql.control)
- [官方扩展 SQL (timeql--1.0.sql)](https://github.com/animeshs34/timeql/blob/a3032739628758757e223d49fc67b276c6e14f8c/timeql--1.0.sql)

`timeql` — TimeQL 是一个 PostgreSQL 扩展，提供了原生的高性能时间数据跟踪和“时间点查询”能力。它使用 PostgreSQL 的原生范围类型（tsrange）和 GiST 索引来确保跟踪数据的历史不会影响性能。使用它来处理相应的调度、时间相关或时间序列工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION timeql;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `timeql.create_history_table(p_schema TEXT, p_table TEXT)` 是一个扩展函数，返回 `void`。
- `timeql.create_history_table(p_schema TEXT, p_table TEXT, p_retention_period INTERVAL DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `timeql.disable_temporal(p_schema TEXT, p_table TEXT)` 是一个扩展函数，返回 `void`。
- `timeql.enable_temporal(p_schema TEXT, p_table TEXT)` 是一个扩展函数，返回 `void`。
- `timeql.purge_all_history()` 是一个扩展函数，返回 `void`。
- `timeql.purge_history(p_schema TEXT, p_table TEXT)` 是一个扩展函数，返回 `void`。
- `timeql.restore_at(p_table_reg regclass, p_pk_value ANYELEMENT, p_timestamp TIMESTAMP WITHOUT TIME ZONE)` 是一个扩展函数，返回 `void`。
- `timeql.set_retention_policy(p_schema TEXT, p_table TEXT, p_retention_period INTERVAL)` 是一个扩展函数，返回 `void`。
- `timeql.tql_at(p_table_reg regclass, p_timestamp TIMESTAMP WITHOUT TIME ZONE)` 是一个扩展函数，返回 `SETOF`。
- `timeql.tql_between(p_table_reg regclass, p_start TIMESTAMP WITHOUT TIME ZONE, p_end TIMESTAMP WITHOUT TIME ZONE)` 是一个扩展函数，返回 `SETOF`。
- `timeql.track_delete()` 是一个扩展函数，返回 `TRIGGER`。
- `timeql.track_insert()` 是一个扩展函数，返回 `TRIGGER`。
- `timeql.track_update()` 是一个扩展函数，返回 `TRIGGER`。
- `timeql_version()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
