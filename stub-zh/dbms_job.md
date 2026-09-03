## 用法

来源：

- [DBMS_JOB 官方参考](https://www.enterprisedb.com/docs/epas/latest/reference/oracle_compatibility_reference/epas_compat_bip_guide/03_built-in_packages/05_dbms_job/)
- [EDB Job Scheduler 官方文档](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/)
- [Clone Schema 官方设置指南](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/setting_up_edb_clone_schema/)

`dbms_job` 提供 EPAS Oracle-compatible `DBMS_JOB` package，用于创建、修改、运行与删除 stored procedure job。

### 启用

`dbms_job` 依赖预加载的 EDB Job Scheduler。先配置并创建 scheduler，再创建 package 扩展：

```ini
edb_job_scheduler.database_list = 'appdb'
shared_preload_libraries = 'edb_job_scheduler'
```

```sql
CREATE EXTENSION edb_job_scheduler;
CREATE EXTENSION dbms_job;
```

该 provider-only 扩展只在 EDB Postgres Advanced Server 上提供，不适用于 community PostgreSQL。

### 创建并管理 Job

```sql
DECLARE
  job_id bigint;
BEGIN
  DBMS_JOB.SUBMIT(
    job_id,
    'job_proc;',
    SYSDATE,
    'SYSDATE + 1/24'
  );
  COMMIT;
END;
```

使用 `DBMS_JOB.CHANGE`、`DBMS_JOB.INTERVAL`、`DBMS_JOB.NEXT_DATE`、`DBMS_JOB.RUN`、`DBMS_JOB.BROKEN` 与 `DBMS_JOB.REMOVE` 管理 lifecycle。

### 运维边界

`what` 是稍后由 worker 执行的 stored procedure text。应限制 package privilege，避免从不可信输入拼接它，并让 recurring work 保持幂等。`next_date` 与 `interval` 会交互：每次运行前重新计算 interval expression，并把结果作为下一次调度时间。应监控 `sys.job_run_details` 的失败记录，不能把成功提交等同于成功完成。

