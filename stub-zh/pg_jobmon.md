


## 用法

> [pg_jobmon: PostgreSQL 自主作业日志和监控](https://github.com/omniti-labs/pg_jobmon)

`pg_jobmon` 为 PostgreSQL 事务和函数提供自主（非事务性）日志记录。如果函数失败，到该时间点为止写入的所有日志信息都会被保留，而不会被回滚。

```sql
CREATE SCHEMA jobmon;
CREATE EXTENSION pg_jobmon SCHEMA jobmon;
```

### 初始化

该扩展使用 dblink 回连到同一数据库（用于非事务性日志记录）。添加凭证：

```sql
INSERT INTO jobmon.dblink_mapping_jobmon (username, pwd) VALUES ('rolename', 'rolepassword');
```

非标准端口：

```sql
INSERT INTO jobmon.dblink_mapping_jobmon (host, username, pwd, port)
VALUES ('localhost', 'rolename', 'rolepassword', '5999');
```

### 核心日志函数

```sql
-- 开始新作业
SELECT jobmon.add_job('My Job Name');

-- 向作业添加步骤
SELECT jobmon.add_step(job_id, 'Step description');

-- 更新步骤状态
SELECT jobmon.update_step(step_id, 'OK', 'Step completed successfully');
SELECT jobmon.update_step(step_id, 'WARNING', 'Something unexpected');

-- 关闭作业
SELECT jobmon.close_job(job_id);

-- 或标记作业失败
SELECT jobmon.fail_job(job_id);
```

### 监控函数

```sql
-- 检查失败的作业
SELECT * FROM jobmon.check_job_status();

-- 查看作业历史
SELECT * FROM jobmon.job_log ORDER BY start_time DESC;

-- 查看步骤详情
SELECT * FROM jobmon.job_detail WHERE job_id = 123;
```

自主日志记录确保即使父事务回滚，作业日志条目也会被保留以便排障。
