## 用法

来源：

- [官方上游 README](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/jobs/readme.md)
- [官方扩展控制文件 (skitch-extension-jobs.control)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/jobs/skitch-extension-jobs.control)
- [官方扩展 SQL (skitch-extension-jobs--0.0.7.sql)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/jobs/sql/skitch-extension-jobs--0.0.7.sql)

`skitch-extension-jobs` — 一个用于 ACID 合规作业创建的异步作业队列模式。当应用程序需要此特定数据库功能时使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION "skitch-extension-jobs";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `app_jobs.add_job(identifier varchar, payload json)` 是一个扩展函数，返回 `app_jobs`。
- `app_jobs.add_job(identifier varchar, queue_name varchar, payload json)` 是一个扩展函数，返回 `app_jobs`。
- `app_jobs.complete_job(worker_id varchar, job_id int)` 是一个扩展函数，返回 `app_jobs`。
- `app_jobs.do_notify()` 是一个扩展函数，返回 `trigger`。
- `app_jobs.fail_job(worker_id varchar, job_id int, error_message varchar)` 是一个扩展函数，返回 `app_jobs`。
- `app_jobs.get_job(worker_id varchar, identifiers varchar[])` 是一个扩展函数，返回 `app_jobs`。
- `app_jobs.schedule_job(identifier varchar, queue_name varchar, payload json, run_at timestamptz)` 是一个扩展函数，返回 `app_jobs`。
- `app_jobs.tg__add_job_for_row()` 是一个扩展函数，返回 `trigger`。
- `app_jobs.tg_decrease_job_queue_count()` 是一个扩展函数，返回 `trigger`。
- `app_jobs.tg_increase_job_queue_count()` 是一个扩展函数，返回 `trigger`。
- `app_jobs.job_queues` 是一个由扩展安装或管理的表。
- `app_jobs.jobs` 是一个由扩展安装或管理的表。
- `app_jobs` 是一个由扩展创建的模式。

### 要求与注意事项

- 审核后的控制文件声明默认版本为 `0.0.7`。
- 首先安装并验证确认的扩展依赖项：`plpgsql`, `pgcrypto`, `uuid-ossp`, `skitch-extension-defaults`, `skitch-extension-verify`, `skitch-extension-utils`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
