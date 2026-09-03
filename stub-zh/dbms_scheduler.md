## 用法

来源：

- [DBMS_SCHEDULER 官方参考](https://www.enterprisedb.com/docs/epas/latest/reference/oracle_compatibility_reference/epas_compat_bip_guide/03_built-in_packages/15_dbms_scheduler/)
- [EDB Job Scheduler 官方文档](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/)
- [官方 scheduler 配置](https://www.enterprisedb.com/docs/pg_extensions/edb_job_scheduler/configuring/)

`dbms_scheduler` 提供 EPAS Oracle-compatible `DBMS_SCHEDULER` package，用于管理具名 program、schedule、argument 与 job。

### 启用

配置 EDB Job Scheduler worker、重启，然后创建两个扩展：

```ini
edb_job_scheduler.database_list = 'appdb'
shared_preload_libraries = 'edb_job_scheduler'
```

```sql
CREATE EXTENSION edb_job_scheduler;
CREATE EXTENSION dbms_scheduler;
```

该 provider-only 扩展在 EDB Postgres Advanced Server 上提供。

### 定义并运行 Job

`DBMS_SCHEDULER.CREATE_PROGRAM` 定义可执行 program 与 argument。`DBMS_SCHEDULER.CREATE_SCHEDULE` 定义时间，`DBMS_SCHEDULER.CREATE_JOB` 绑定它们或直接指定 action。

```sql
BEGIN
  DBMS_SCHEDULER.CREATE_JOB(
    job_name        => 'hourly_rollup',
    job_type        => 'STORED_PROCEDURE',
    job_action      => 'refresh_rollup',
    start_date      => SYSDATE,
    repeat_interval => 'FREQ=HOURLY',
    enabled         => TRUE
  );
END;
```

使用 `RUN_JOB`、`ENABLE`、`DISABLE`、`DROP_JOB`、`DROP_PROGRAM` 与 `DROP_SCHEDULE` 管理 lifecycle。

### 兼容性边界

EDB 实现的是文档化的 Oracle `DBMS_SCHEDULER` 子集；不能假定未列出的 Oracle attribute 或语义。Job action 使用数据库权限运行，若 procedure 不具备事务性或幂等性，可能留下部分结果。启用 recurring job 前，应监控 scheduler table，并用 `EVALUATE_CALENDAR_STRING` 测试 calendar string。

