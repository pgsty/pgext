## 用法

来源：

- [packages/database-jobs/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/README.md)
- [packages/database-jobs/sql/pgpm-database-jobs--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/sql/pgpm-database-jobs--0.47.0.sql)
- [packages/database-jobs/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/Makefile)
- [packages/database-jobs/pgpm-database-jobs.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/database-jobs/pgpm-database-jobs.control)

`pgpm-database-jobs` 是 Constructive 平台使用的带归属 `app_jobs` 实现，保存任务、计划和工作进程状态，并附带数据库及身份上下文。

### 核心用法

```sql
CREATE EXTENSION "pgpm-database-jobs" CASCADE;
SELECT * FROM app_jobs.jobs LIMIT 10;
SELECT * FROM app_jobs.scheduled_jobs LIMIT 10;
```

### 运行边界

入队时应使用对应版本的函数签名，保留数据库及操作者归属。`app_jobs.get_job`、`app_jobs.complete_job` 和失败、释放辅助函数管理工作进程所有权与重试。载荷由外部工作进程执行，可重试任务应具有幂等性。不要同时安装同样使用 `app_jobs` 的独立 `pgpm-jobs` 或历史 `launchql-ext-jobs` 变体。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`plpgsql`、`pgcrypto`、`pgpm-verify`、`pgpm-jwt-claims`、`errors`。 安装 SQL 要求平台角色 `administrator`, `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
