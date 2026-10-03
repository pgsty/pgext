## 用法

来源：

- [packages/jobs-simple/launchql-ext-jobs-queue.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs-simple/launchql-ext-jobs-queue.control)
- [packages/jobs-simple/sql/launchql-ext-jobs-queue--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs-simple/sql/launchql-ext-jobs-queue--0.4.5.sql)
- [packages/jobs-simple/readme.md](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs-simple/readme.md)

`launchql-ext-jobs-queue` 0.4.5 在 PostgreSQL 中以事务方式存储异步任务，实际应用任务由外部工作进程执行。

### 核心工作流

创建扩展前安装依赖，并准备好 `administrator` 角色。工作进程以唯一名称认领任务，处理完成后确认成功，或报告返回任务 ID 对应的失败。

```sql
CREATE EXTENSION "launchql-ext-jobs-queue" CASCADE;
SELECT app_jobs.add_job('send_email', '{"recipient":"reader@example.com"}'::json);
SELECT * FROM app_jobs.get_job('worker_1');
```

### 对象与依赖

`app_jobs.jobs` 与 `app_jobs.job_queues` 保存任务及队列状态。`app_jobs.complete_job(worker_id, job_id)` 确认已认领任务完成；`app_jobs.fail_job(worker_id, job_id, error_message)` 记录失败。`app_jobs.add_scheduled_job`、`app_jobs.get_scheduled_job` 与 `app_jobs.run_scheduled_job` 管理定时任务。依赖为 `plpgsql`、`pgcrypto`、`uuid-ossp` 与 `launchql-ext-default-roles`。

### 边界

当前 README 的示例使用另一命名空间；实际 API 应以上述版本化 SQL 为准。两个 LaunchQL 任务扩展会创建相同的 `app_jobs` 对象，不能同时安装。无需共享库或预加载。安装 SQL 函数需要相应模式权限；控制文件允许非超级用户安装，不等于补齐所缺的角色或权限。工作进程需要处理重试、超时认领和任务保留。不能仅因使用 SQL 就推断 PostgreSQL 主版本支持范围。
