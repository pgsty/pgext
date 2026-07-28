## 用法

来源：

- [官方上游 README](https://github.com/splendiddata/pg_cron_helper/blob/5bc4ea422971babfc4a1e8ba5a8f5b2f20e8c62a/README.md)
- [官方扩展控制文件 (pg_cron_helper.control)](https://github.com/splendiddata/pg_cron_helper/blob/5bc4ea422971babfc4a1e8ba5a8f5b2f20e8c62a/pg_cron_helper.control)
- [官方扩展 SQL (pg_cron_helper--0.1.sql)](https://github.com/splendiddata/pg_cron_helper/blob/5bc4ea422971babfc4a1e8ba5a8f5b2f20e8c62a/pg_cron_helper--0.1.sql)

`pg_cron_helper` — pg_cron_helper 数据库扩展旨在帮助在基于时间的简单调度上运行“作业”。当然，使用外部调度器会更好，但如果必须在 PostgreSQL 数据库内部安排作业，那么此扩展可能会有所帮助。请使用它进行相应的调度、时间相关或时间序列工作流。在安装扩展及其依赖项并验证它们之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION pg_cron_helper;

select * from  cron.list_jobs();
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `get_job_state(job_name varchar(128) , user_name name default current_user)` 是一个扩展函数，返回 `text`。
- `list_jobs()` 是一个扩展函数，返回 `setof`。
- `create_job` 是一个扩展过程。
- `cron.stop_job` 是一个扩展过程。
- `disable_job` 是一个扩展过程。
- `drop_job` 是一个扩展过程。
- `enable_job` 是一个扩展过程。
- `run_job` 是一个扩展过程。
- `stop_job` 是一个扩展过程。
- `job_record` 是一个扩展定义的类型。
- `job_definition` 是一个由扩展安装或管理的表。
- `job_run` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.3`。
- 请首先安装确认的扩展依赖项：`postgres_fdw`, `dblink`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
