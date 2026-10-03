## 用法

来源：

- [v2.0.4 README](https://github.com/vibhorkum/pg_background/blob/v2.0.4/README.md)
- [v2.0.4 release](https://github.com/vibhorkum/pg_background/releases/tag/v2.0.4)
- [v2.0.4 SQL API](https://github.com/vibhorkum/pg_background/blob/v2.0.4/extension/pg_background--2.0.sql)

`pg_background` 在 PostgreSQL 后台工作进程中执行 SQL。工作进程在服务器内部运行独立事务，适合异步维护、自主副作用、有边界的长时间任务，以及可跟踪进度的作业。

2.0 版本将不带后缀的 API 作为标准接口。旧的 `_v2` 名称在 2.x 系列中仍保留为已弃用别名，但新代码应使用 `pg_background_launch`、`pg_background_result`、`pg_background_run` 等名称。

### 一次性执行

当 SQL 只有副作用，而你只需要执行元数据时，使用 `pg_background_run`：

```sql
CREATE EXTENSION pg_background;

SELECT completed, has_error, sqlstate, error_message,
       row_count, command_tag, elapsed_ms, timed_out
FROM pg_background_run(
  'INSERT INTO audit_log(ts, who) VALUES (clock_timestamp(), current_user)',
  queue_size := 0,
  timeout_ms := 30000,
  label := 'audit-login'
);
```

### 启动并获取结果

当后台 SQL 返回行时，使用 launch/result 模式。在 psql 中，`\gset` 将句柄保存为后续命令可用的变量：

```sql
SELECT pid, cookie FROM pg_background_launch(
  'SELECT count(*) FROM large_table',
  queue_size := 65536,
  label := 'count-large-table'
) \gset bg_

SELECT * FROM pg_background_result(:bg_pid, :bg_cookie) AS (count bigint);
```

结果只能消费一次。请同时保存 `pid` 和 `cookie`；cookie 用来避免后续调用受到 PID 重用影响。

### 即发即忘

对于不需要消费结果行的副作用：

```sql
SELECT * FROM pg_background_submit(
  $$VACUUM (ANALYZE) public.events$$,
  queue_size := 65536,
  label := 'vacuum-events'
);
```

### 核心 API

- `pg_background_launch(sql, queue_size, label)` 启动工作进程，并返回 `pg_background_handle(pid, cookie)`。
- `pg_background_submit(sql, queue_size, label)` 启动即发即忘的工作，并返回一个句柄。
- `pg_background_result(pid, cookie)` 一次性消费结果行。
- `pg_background_result_info(pid, cookie)` 返回完成状态和行数元数据，但不消费结果行。
- `pg_background_error_info(pid, cookie)` 返回结构化 SQL 错误详情。
- `pg_background_wait(pid, cookie, timeout_ms DEFAULT 0)` 等待完成；`timeout_ms <= 0` 表示无限期阻塞。
- `pg_background_cancel(pid, cookie, grace_ms DEFAULT 0)` 请求协作式取消。
- `pg_background_detach(pid, cookie)` 停止跟踪工作进程，但允许它继续运行。
- `pg_background_outcome(pid, cookie)` 返回合并后的状态快照，即使状态缺失也不会抛错。
- `pg_background_list` 和 `pg_background_activity` 是监控视图；`pg_background_stats()` 返回会话计数器。

便捷辅助函数包括 `pg_background_run_query`、`pg_background_drain`、`pg_background_wait_any`、`pg_background_cancel_by_label` 和 `pg_background_purge`。

### 进度报告

在工作进程 SQL 内报告进度，再由启动方轮询：

```sql
SELECT pid, cookie FROM pg_background_launch($$
  SELECT pg_background_report_progress(0, 'starting');
  SELECT pg_sleep(1);
  SELECT pg_background_report_progress(100, 'done');
$$) \gset bg_

SELECT * FROM pg_background_get_progress(:bg_pid, :bg_cookie);
```

`pg_background_report_progress` 是 2.0 名称；更早的 `pg_background_progress` 名称已经硬重命名。

### GUC 与加载

`pg_background` 不要求配置 `shared_preload_libraries`。预加载是可选的，主要用于希望会话首次加载扩展前就能使用其 GUC 的场景。

```sql
SET pg_background.max_workers = 10;
SET pg_background.default_queue_size = '256KB';
SET pg_background.worker_timeout = '5min';
```

- `pg_background.max_workers` 默认值为 `16`。
- `pg_background.default_queue_size` 默认值为 `65536` 字节。
- `pg_background.worker_timeout` 默认值为 `0`，表示没有执行超时。

### 注意事项

- 上游发行版 2.0.4 保留 SQL/control 版本 2.0，面向 PostgreSQL 14-18 和 PostgreSQL 19 beta。
- 从 1.8 之前的安装升级时，必须先升级到 1.8/1.10 发布线，再更新到 2.0。
- 原始 v1 的仅 PID API 已移除。不带后缀的名称现在具备 cookie 保护语义，并返回/使用 `(pid, cookie)`。
- `pg_background_cancel_v2_grace` 和 `pg_background_wait_v2_timeout` 已合并进 `pg_background_cancel(..., grace_ms)` 和 `pg_background_wait(..., timeout_ms)`。
- `pg_background_status_v2` 已移除；请使用 `pg_background_outcome(pid, cookie)`。

### 2.0.4 更新与权限

2.0.4 修复同一工作进程内多条命令之间的数据可见性、复杂结果类型的传输和命令完成标签。`pg_background.worker_timeout` 非零时限制整个 SQL 字符串的执行时间（不含提交）；为零时由 `statement_timeout` 分别限制每条语句。安装新版库后启动新的工作进程才能使用修复。

安装需要超级用户。默认不向 PUBLIC 开放调用权限；由管理员授予 `pgbackground_role`，或使用 `pg_background_grant_privileges`。后台事务独立提交，调用者回滚不会撤销已经完成的后台操作。
