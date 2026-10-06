## 用法

来源：

- [packages/jobs/readme.md](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/readme.md)
- [packages/jobs/sql/launchql-ext-jobs--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/sql/launchql-ext-jobs--0.4.5.sql)
- [packages/jobs/Makefile](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/Makefile)
- [packages/jobs/launchql-ext-jobs.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jobs/launchql-ext-jobs.control)

`launchql-ext-jobs` 在 `app_jobs` 中保存任务、计划任务和队列锁，任务由外部工作进程领取并执行。

### 核心用法

```sql
CREATE EXTENSION "launchql-ext-jobs" CASCADE;
SELECT * FROM app_jobs.jobs LIMIT 10;
SELECT * FROM app_jobs.scheduled_jobs LIMIT 10;
```

### 运行边界

`app_jobs.add_job` 接收数据库 UUID、任务标识符和 JSON 载荷。`app_jobs.get_job`、`app_jobs.complete_job`、`app_jobs.fail_job` 与 `app_jobs.release_jobs` 协调任务归属及重试，可重试任务应具有幂等性。此历史模式与 `pgpm-jobs` 和 `pgpm-database-jobs` 冲突，不应同时安装。SQL 要求配套默认角色与 JWT 声明集成。

须先安装声明的依赖：`plpgsql`, `uuid-ossp`, `pgcrypto`, `launchql-ext-default-roles`, `launchql-jwt-claims`。这是 SQL/PLpgSQL 代码，没有自己的共享库，也不要求预加载。控制文件允许非超级用户安装，但没有标记为 trusted；仍须满足依赖和模式权限。
