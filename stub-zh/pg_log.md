## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_log/pg_log-0.0.3/README.md)
- [官方扩展控制文件 (pg_log.control)](https://api.pgxn.org/src/pg_log/pg_log-0.0.3/pg_log.control)
- [官方扩展 SQL (pg_log--0.0.1.sql)](https://api.pgxn.org/src/pg_log/pg_log-0.0.3/pg_log--0.0.1.sql)

`pg_log` — PostgreSQL 扩展，用于显示 SQL 日志。在收集或解释相应的 PostgreSQL 统计信息时使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_log;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_get_logname()` 是一个扩展函数，返回 `cstring`。
- `pg_log(OUT line integer, OUT message text)` 是一个扩展函数，返回 `SETOF`。
- `pg_log_refresh()` 是一个扩展函数，返回 `void`。
- `pg_read(cstring)` 是一个扩展函数，返回 `void`。
- `log` 是一个由扩展定义的视图。
- `pglog` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 仓库记录版本 `1.0.0`，而审核过的控制文件声明默认版本 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
