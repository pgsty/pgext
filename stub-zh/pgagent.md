

## 用法

> [pgagent: PostgreSQL 作业调度器](https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html)

pgAgent 是 PostgreSQL 的作业调度代理，能够在复杂的调度计划上运行多步骤的批处理/Shell 脚本和 SQL 任务。它以守护进程方式运行，并将作业定义存储在数据库中。

### 核心概念

- **作业（Job）**：包含一个或多个步骤和调度计划的可调度命名单元
- **步骤（Step）**：作业中的单个操作（SQL 脚本或 OS 批处理/Shell 命令）
- **调度（Schedule）**：定义作业何时运行，具有类似 cron 的灵活性

### 通过 SQL 管理作业

pgAgent 将其配置存储在 `pgagent` 模式中。作业可以通过 pgAdmin 或直接通过 SQL 管理。

```sql
-- 查看所有作业
SELECT jobid, jobname, jobenabled, jobdesc
FROM pgagent.pga_job;

-- 查看作业步骤
SELECT jstid, jstjobid, jstname, jstenabled, jstkind, jstcode
FROM pgagent.pga_jobstep;

-- 查看作业调度
SELECT jscid, jscjobid, jscname, jscenabled,
       jscstart, jscend, jscminutes, jschours,
       jscweekdays, jscmonthdays, jscmonths
FROM pgagent.pga_schedule;

-- 查看作业执行日志
SELECT * FROM pgagent.pga_joblog
WHERE jlgjobid = 1 ORDER BY jlgstart DESC;

-- 查看步骤执行日志
SELECT * FROM pgagent.pga_jobsteplog
WHERE jsljlgid IN (SELECT jlgid FROM pgagent.pga_joblog WHERE jlgjobid = 1)
ORDER BY jslstart DESC;
```

### 步骤类型

| 类型 | 描述 |
|------|-------------|
| `s` | 针对数据库执行的 SQL 脚本 |
| `b` | 在操作系统上执行的批处理/Shell 命令 |

### 调度字段

| 字段 | 描述 |
|-------|-------------|
| `jscstart` / `jscend` | 调度的有效日期范围 |
| `jscminutes` | 布尔数组[60]：在哪些分钟运行 |
| `jschours` | 布尔数组[24]：在哪些小时运行 |
| `jscweekdays` | 布尔数组[7]：在一周的哪几天运行 |
| `jscmonthdays` | 布尔数组[32]：在一月的哪几天运行 |
| `jscmonths` | 布尔数组[12]：在哪些月份运行 |

### 安全性

pgAgent 守护进程使用存储的连接字符串连接到数据库。只有数据库超级用户或被授予 `pgagent` 模式表适当权限的用户才能管理作业。
