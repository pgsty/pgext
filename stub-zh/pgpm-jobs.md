## 用法

来源：

- [packages/jobs/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/README.md)
- [packages/jobs/sql/pgpm-jobs--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/sql/pgpm-jobs--0.47.0.sql)
- [packages/jobs/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/Makefile)
- [packages/jobs/pgpm-jobs.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/jobs/pgpm-jobs.control)

`pgpm-jobs` 提供独立安装使用的 `app_jobs` 队列变体，不带身份归属字段。SQL 保存任务、计划、锁定和重试状态，实际任务须由外部工作进程执行。

### 核心用法

```sql
CREATE EXTENSION "pgpm-jobs" CASCADE;
SELECT app_jobs.add_job('demo', '{}'::json);
SELECT * FROM app_jobs.jobs LIMIT 10;
```

### 运行边界

`app_jobs.get_job`、`app_jobs.complete_job` 及失败、释放辅助函数协调工作进程；LISTEN/NOTIFY 用于唤醒消费者，本身不是持久任务执行器。此变体不传播租户或计费身份；平台带归属版本应使用 `pgpm-database-jobs`，两者创建相同模式，不能同时安装。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`plpgsql`、`pgcrypto`、`pgpm-verify`、`errors`。 安装 SQL 要求平台角色 `administrator`, `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
