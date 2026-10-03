## 用法

来源：

- [packages/jobs/launchql-extension-jobs.control](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/jobs/launchql-extension-jobs.control)
- [packages/jobs/sql/launchql-extension-jobs--0.0.2.sql](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/jobs/sql/launchql-extension-jobs--0.0.2.sql)
- [packages/jobs/readme.md](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/jobs/readme.md)

`launchql-extension-jobs` 0.0.2 在 PostgreSQL 中以事务方式存储异步任务，实际应用任务由外部工作进程执行。

### 核心工作流

创建扩展前安装依赖，并准备好 `administrator` 角色。工作进程以唯一名称认领任务，处理完成后确认成功，或报告返回任务 ID 对应的失败。

```sql
CREATE EXTENSION "launchql-extension-jobs" CASCADE;
SELECT app_jobs.add_job('send_email', '{"recipient":"reader@example.com"}'::json);
SELECT * FROM app_jobs.get_job('worker_1');
```

### 对象与依赖

`app_jobs.jobs` 与 `app_jobs.job_queues` 保存任务及队列状态。`app_jobs.complete_job(worker_id, job_id)` 确认已认领任务完成；`app_jobs.fail_job(worker_id, job_id, error_message)` 记录失败。`app_jobs.add_scheduled_job`、`app_jobs.get_scheduled_job` 与 `app_jobs.run_scheduled_job` 管理定时任务。依赖为 `plpgsql`、`pgcrypto`、`uuid-ossp`。

### 边界

安装器会向已有管理员角色授权，但该角色并未声明为扩展依赖。两个 LaunchQL 任务扩展会创建相同的 `app_jobs` 对象，不能同时安装。无需共享库或预加载。安装 SQL 函数需要相应模式权限；控制文件允许非超级用户安装，不等于补齐所缺的角色或权限。工作进程需要处理重试、超时认领和任务保留。不能仅因使用 SQL 就推断 PostgreSQL 主版本支持范围。
