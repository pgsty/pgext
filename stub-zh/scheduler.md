## 用法

来源：

- [官方上游 README](https://github.com/victornagibator/pg_scheduler/blob/48f3bdb0119a457d7af410057a349bd023d7b813/README.md)
- [官方扩展控制文件 (scheduler.control)](https://github.com/victornagibator/pg_scheduler/blob/48f3bdb0119a457d7af410057a349bd023d7b813/scheduler.control)
- [官方扩展 SQL (scheduler--1.0.sql)](https://github.com/victornagibator/pg_scheduler/blob/48f3bdb0119a457d7af410057a349bd023d7b813/scheduler--1.0.sql)

`scheduler` — pg_scheduler 是一个 PostgreSQL 扩展，允许在数据库中灵活地调度和执行 SQL 和 shell 作业。使用它来实现相应的调度、时间相关或时间序列工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION scheduler;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `scheduler.add_or_update_job(p_name TEXT, p_type TEXT, p_cmd TEXT, p_interval INTERVAL = NULL, p_time TIMESTAMPTZ = NULL, p_max_attempts INT = 3)` 是一个扩展函数，返回 `TEXT`。
- `scheduler.delete_job(p_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `scheduler.execute_job(target_job_id INT)` 是一个扩展函数，返回 `VOID`。
- `scheduler.execute_shell_command(cmd TEXT)` 是一个扩展函数，返回 `VOID`。
- `scheduler.set_next_run()` 是一个扩展函数，返回 `TRIGGER`。
- `scheduler.toggle_job(p_name TEXT, p_enabled BOOLEAN)` 是一个扩展函数，返回 `VOID`。
- `scheduler.update_timestamp()` 是一个扩展函数，返回 `TRIGGER`。
- `scheduler.job_logs` 是一个由扩展安装或管理的表。
- `scheduler.jobs` 是一个由扩展安装或管理的表。
- `scheduler` 是一个由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
