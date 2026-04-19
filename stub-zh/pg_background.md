## 用法

来源: [official README](https://github.com/vibhorkum/pg_background/blob/master/README.md), [v1.9.2 release](https://github.com/vibhorkum/pg_background/releases/tag/v1.9.2)

`pg_background` 在 PostgreSQL 后台工作进程中执行 SQL。工作进程运行在服务器内部，并使用自己的事务，因此适合异步维护、自主副作用，以及不希望阻塞调用方的长时间任务。

```sql
CREATE EXTENSION pg_background;

SELECT * FROM pg_background_launch_v2(
  'SELECT count(*) FROM large_table',
  65536,
  'count-large-table'
) AS h;

SELECT * FROM pg_background_result_v2(h.pid, h.cookie) AS (count bigint);
```

### 核心 API

- `pg_background_launch_v2(sql, queue_size, label)`：启动可跟踪的工作进程，并返回 `(pid, cookie)`。
- `pg_background_submit_v2(sql, queue_size, label)`：即发即忘，适合只需要副作用的 SQL。
- `pg_background_result_v2(pid, cookie)`：一次性消费工作进程的结果集。
- `pg_background_wait_v2(...)` 和 `pg_background_wait_v2_timeout(...)`：等待任务完成。
- `pg_background_cancel_v2(...)`：停止执行；`pg_background_detach_v2(...)`：停止跟踪但让任务继续运行。
- `pg_background_list_v2()`、`pg_background_stats_v2()` 和 `pg_background_get_progress_v2(...)`：查看工作进程状态和进度。

### 典型模式

在不保持客户端会话打开的情况下运行维护任务：

```sql
SELECT * FROM pg_background_submit_v2(
  'VACUUM (ANALYZE) public.events',
  65536,
  'vacuum-events'
);
```

在应用 SQL 中触发一个自主副作用：

```sql
SELECT * FROM pg_background_submit_v2(
  $$INSERT INTO audit_log(ts, event) VALUES (clock_timestamp(), 'job queued')$$
);
```

### GUC 与安全

- `pg_background.max_workers` 限制每个会话的并发工作进程数。
- `pg_background.default_queue_size` 控制共享内存队列大小。
- `pg_background.worker_timeout` 设置执行超时；`0` 表示不限制。
- 扩展会创建专用的 `pg_background_worker` NOLOGIN 角色，并提供辅助函数来授予或撤销执行权限。

### 注意事项

- 优先使用 V2 API。旧的 V1 API 仍然保留用于兼容，但缺少基于 cookie 的 PID 重用保护。
- `v1.9.2` 是仅涉及二进制构建的补丁版本，对 assert-enabled PostgreSQL 构建做了修复。SQL 扩展版本仍是 `1.9`，因此相比 `1.9.1` 没有新的 SQL 升级脚本或面向用户的函数变化。
