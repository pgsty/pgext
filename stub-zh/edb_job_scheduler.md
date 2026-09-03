## 用法

来源：

- [EDB Job Scheduler 官方文档](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/configuring/)
- [官方用户 API](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/user_api/)

`edb_job_scheduler` 在 EDB Postgres Advanced Server 中用 background worker 运行已保存数据库 job，并提供 `dbms_job` 与 `dbms_scheduler` 使用的 runtime。

### 启用

列出每个需要调度的数据库、预加载 worker、重启，并在每个已列数据库中创建扩展：

```ini
edb_job_scheduler.database_list = 'appdb'
shared_preload_libraries = 'edb_job_scheduler'
```

```sql
CREATE EXTENSION edb_job_scheduler;
```

扩展创建 `sys.jobs` 与 `sys.job_run_details`。生产使用前应配置 `edb_job_scheduler.max_jobs_per_database`、`edb_job_scheduler.max_workers_per_database` 以及 database-list capacity。

### 提交并检查 Job

应用通常通过 `DBMS_JOB` 或 `DBMS_SCHEDULER` 操作；scheduler 也公开 `sys` 表用于状态检查。

```sql
SELECT jobid, jobnextrun, jobcommand, jobuser
FROM sys.jobs
ORDER BY jobnextrun;

SELECT jobid, runid, status, starttime, endtime, error
FROM sys.job_run_details
ORDER BY runid DESC;
```

### 运维边界

Job 以配置的 job user 执行 SQL，并可改变数据库状态。应限制 package execution、校验存储命令、约束 worker、监控失败运行，并设计幂等 job。修改 `database_list` 后需要 reload 还是 restart，取决于 EPAS 版本与配置容量；未列出的数据库不会获得 scheduler worker。

