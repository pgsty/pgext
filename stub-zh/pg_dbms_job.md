

## 用法

> [pg_dbms_job: 为 PostgreSQL 添加 Oracle DBMS_JOB 完整兼容的扩展](https://github.com/MigOpsRepos/pg_dbms_job)

### 启用

```sql
CREATE EXTENSION pg_dbms_job;
```

需要为每个数据库运行专用调度守护进程：

```bash
pg_dbms_job -c /etc/pg_dbms_job/mydb-dbms_job.conf
```

### SUBMIT - 调度作业

```sql
BEGIN;
-- 调度作业：每天运行一个过程
CALL dbms_job.submit(
    job       => jobid,
    what      => 'CALL my_procedure();',
    next_date => current_timestamp + interval '1 minute',
    interval  => 'current_timestamp + ''1 day''::interval'
);
COMMIT;
```

省略 `next_date` 和 `interval` 时，作业立即异步执行。

### BROKEN - 禁用/启用作业

```sql
BEGIN;
CALL dbms_job.broken(12345, true);   -- 禁用作业
CALL dbms_job.broken(12345, false);  -- 重新启用作业
COMMIT;
```

### CHANGE - 修改作业

```sql
BEGIN;
CALL dbms_job.change(12345, null, null, 'current_timestamp + ''3 days''::interval');
COMMIT;
```

### INTERVAL - 更改执行间隔

```sql
BEGIN;
CALL dbms_job.interval(12345, 'current_timestamp + ''1 hour''::interval');
COMMIT;
```

### NEXT_DATE - 更改下次执行日期

```sql
BEGIN;
CALL dbms_job.next_date(12345, current_timestamp + interval '30 minutes');
COMMIT;
```

### WHAT - 更改作业代码

```sql
BEGIN;
CALL dbms_job.what(12345, 'CALL new_procedure();');
COMMIT;
```

### REMOVE - 删除作业

```sql
BEGIN;
CALL dbms_job.remove(12345);
COMMIT;
```

### RUN - 立即执行

```sql
CALL dbms_job.run(12345);
```

### 查看作业

```sql
SELECT * FROM dbms_job.all_jobs;
```

### 执行历史

```sql
SELECT * FROM dbms_job.all_scheduler_job_run_details;
```
