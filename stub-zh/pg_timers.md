## 用法

来源：

- [官方 README](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/README.md)
- [SQL 函数定义](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/sql/migrations/002_functions.sql)
- [后台工作进程实现](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/src/bgworker.c)

`pg_timers` 在确切时间戳或给定间隔后调度一次性 SQL 操作。单个预加载后台工作进程会在下一条记录到期时唤醒，以调度该操作的用户身份执行，并在 `timers.timers` 表记录已触发、失败或已取消状态。

### 预加载与调度

选择工作进程唯一服务的数据库，预加载模块并重启 PostgreSQL，再在该数据库创建不可重定位扩展：

```conf
shared_preload_libraries = 'pg_timers'
pg_timers.database = 'app'
```

```sql
CREATE EXTENSION pg_timers;

SELECT timers.schedule_at(
  '2026-08-01 09:00:00+08',
  $$DELETE FROM sessions WHERE expires_at < now()$$,
  0,
  30000
);

SELECT timers.schedule_in(
  interval '5 minutes',
  $$SELECT app.expire_session(42)$$
);
```

可选的 `shard_key` 支持 Citus 分布，默认值为 0。`timeout_ms` 设置每个操作的语句超时；零表示不限制。调度属于调用者事务，插入若回滚就不会到达工作进程。

### 状态与控制

`timers.timers` 以状态 0 表示等待，1 表示已触发，2 表示失败，3 表示已取消。`timers.cancel` 把仍在等待的行改为取消状态，并通过行锁与工作进程协调。成功会记录 `fired_at`，失败会记录错误文本。

```sql
SELECT timers.cancel(timer_id, shard_key)
FROM timers.timers
WHERE status = 0
  AND id = 42;
```

`timers.fire` 与 `timers.fire_all_pending` 忽略计划时间，同步执行等待中的操作，主要用于事务化 pgTAP 测试。拥有这些权限并能查看表的角色可以提前触发其他用户的计时器，因此不能授予生产应用角色。

### 安全、可靠性与容量

安装时会撤销 `PUBLIC` 对所有函数与表的权限。应显式授予调度、取消与读取权限。操作以记录的 `session_user` 身份运行，而非后台工作进程超级用户；但有权调度的用户可以安排未来以自身身份运行的任意 SQL。

每个操作在子事务中运行。数据库修改和状态转换随工作进程批次原子提交，因此提交前崩溃会让计时器保持等待并重试。不可事务回滚的外部副作用仍可能重复，必须设计为幂等。

单个工作进程在一个配置数据库中顺序处理操作。`pg_timers.max_timers_per_tick` 限制每批数量，`pg_timers.check_interval_ms` 可限制最长休眠时间；两者都可重载。慢操作会阻塞后续计时器，已完成记录也需要保留策略。

当前源码报告开发版本 0.0.0，而 README 给出 0.1.0 镜像标签示例。部署前必须固定真实发行产物并对齐版本。文档构建矩阵覆盖 PostgreSQL 15 至 18；还要在目标环境验证确切产物、Citus 放置、重启/故障转移行为、授权、时钟假设与负载延迟。
