## 用法

来源：

- [官方上游 README](https://github.com/gaoweichang/auto-job/blob/2daa1df1c64ac1fd57652b2ccc5e714d319af0ee/README.md)
- [官方扩展控制文件 (auto_job.control)](https://github.com/gaoweichang/auto-job/blob/2daa1df1c64ac1fd57652b2ccc5e714d319af0ee/auto_job.control)
- [官方扩展 SQL (auto_job--1.0.0.sql)](https://github.com/gaoweichang/auto-job/blob/2daa1df1c64ac1fd57652b2ccc5e714d319af0ee/sql/auto_job--1.0.0.sql)

`auto_job` — Auto job 扩展是用于通过后台工作者自动调度和运行存储过程的 PostgreSQL 扩展。使用它来进行相应的调度、时间序列或时间工作流。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION auto_job;

SELECT pid, application_name, state, query
FROM pg_stat_activity
WHERE datname = 'postgres';
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `add_job(proc_name TEXT, schedule_interval INTERVAL)` 是一个扩展函数，返回 `INTEGER`。
- `delete_job(job_id INTEGER)` 是一个扩展函数，返回 `VOID`。
- `run_job(job_id INTEGER)` 是一个扩展函数，返回 `VOID`。
- `_auto_job_catalog.job_info` 是一个由扩展定义的视图。
- `_auto_job_catalog.jobs` 是一个由扩展安装或管理的表。
- `public.auto_job_registry` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审核后的控制文件声明默认版本为 `1.0.0`。
- 首先安装确认的扩展依赖项：`dblink`。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
