## 用法

来源：

- [官方 README](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/README.md)
- [1.0 版扩展 SQL](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/schedulerx--1.0.sql)
- [后台工作进程实现](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/scheduler.c)

`schedulerx` 是 PostgreSQL 9.5 时代用于无状态 SQL 作业的后台工作进程调度器。一个中央调度进程为一次性、周期性或通知驱动命令启动数据库工作进程，也能为同一作业并行启动多个进程。它有意不建模依赖图或重试工作流。

### 预加载与安装

该库在启动钩子中注册后台工作进程并申请共享内存，因此创建扩展前要把它加入 `shared_preload_libraries` 并重启 PostgreSQL：

```conf
shared_preload_libraries = 'schedulerx'
```

```sql
CREATE EXTENSION schedulerx;
```

创建扩展会生成 `job_schedule`、其约束触发器、辅助函数以及激活当前数据库的 `jobscheduler` 安全标签。可用 `enable_scheduler` 与 `disable_scheduler` 更改该数据库级标签。

### 注册作业

使用绝对开始时间注册一次性作业，或通过重复间隔注册周期作业：

```sql
SELECT register_job_at(
  clock_timestamp() + interval '1 minute',
  'CALL app.refresh_once()',
  current_user,
  'refresh-once'
);

SELECT register_periodic_job(
  clock_timestamp(),
  interval '5 minutes',
  'CALL app.drain_queue()',
  current_user,
  'drain-queue'
);
```

两个辅助函数都会向 `job_schedule` 插入一行。重要列包括 `suspended`、`job_start`、`job_repeat_after`、`job_user`、`job_cmd`、`max_workers`、`run_workers`、`idle_timeout`、`life_time`、`job_start_timeout` 与 `job_timeout`。非空 `job_name` 必须唯一。

`register_listener_job` 原本用于在 `listen_channel` 收到 `NOTIFY` 后启动作业，并可通过 `job_start_delay` 合并通知。但发布的 1.0 版 SQL 在该辅助函数内检查了未定义的标识符；依赖通知作业前，应在确切构建上测试，并优先使用经过审查的 `job_schedule` 直接插入或本地修复。

### 权限与限制

普通用户可以自己的身份调度命令。为其他角色创建、更新或删除作业时，当前用户必须是数据库所有者或超级用户。保存的 SQL 以 `job_user` 权限运行，因此要限制调度表及辅助函数的访问，并审计命令文本。

作业是无状态执行。调度器不提供“重试 N 次”、依赖链或“A 失败后运行 B”等语义。定时与工作进程设置允许时，并发执行可能重叠，因此作业 SQL 必须幂等或自行实现锁。容量同时受 PostgreSQL 后台工作进程资源和每个作业设置限制。

1.0 版较旧，其控制文件没有声明实际运维所需的预加载。应在目标 PostgreSQL 版本上测试启动、重启行为、故障转移、超时执行、配置重载与作业恢复。禁用作业并移除扩展后，要删除预加载项仍需再次重启。
