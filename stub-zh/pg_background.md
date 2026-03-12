

## 用法

> [pg_background: 在后台工作进程中执行 SQL](https://github.com/vibhorkum/pg_background)

在 PostgreSQL 的**后台工作进程**中执行任意 SQL 命令。与 `dblink`（创建独立连接）不同，`pg_background` 的工作进程运行在数据库服务器**内部**，使用**独立事务**。

**使用场景：**
- 后台维护（VACUUM、ANALYZE、REINDEX）
- 异步审计日志记录
- 长时间运行的 ETL 管道
- 独立的通知发送
- 并行查询模式

### 快速开始（V2 API）

```sql
CREATE EXTENSION pg_background;

-- 启动后台任务
SELECT * FROM pg_background_launch_v2(
  'SELECT count(*) FROM large_table'
) AS handle;
--   pid  |      cookie
-- -------+-------------------
--  12345 | 1234567890123456

-- 获取结果（一次性消费）
SELECT * FROM pg_background_result_v2(12345, 1234567890123456) AS (count BIGINT);

-- 即发即忘（不需要结果）
SELECT * FROM pg_background_submit_v2(
  'INSERT INTO audit_log (ts, event) VALUES (now(), ''system_check'')'
) AS handle;
```


## V2 API 参考

| 函数 | 返回值 | 说明 |
|------|--------|------|
| `pg_background_launch_v2(sql, queue_size)` | `pg_background_handle` | 启动工作进程，返回 cookie 保护的句柄 |
| `pg_background_submit_v2(sql, queue_size)` | `pg_background_handle` | 即发即忘（不消费结果） |
| `pg_background_result_v2(pid, cookie)` | `SETOF record` | 获取结果（一次性消费） |
| `pg_background_detach_v2(pid, cookie)` | `void` | 停止跟踪工作进程（进程继续运行） |
| `pg_background_cancel_v2(pid, cookie)` | `void` | 请求取消 |
| `pg_background_cancel_v2_grace(pid, cookie, grace_ms)` | `void` | 带宽限期的取消 |
| `pg_background_wait_v2(pid, cookie)` | `void` | 阻塞等待工作进程完成 |
| `pg_background_wait_v2_timeout(pid, cookie, timeout_ms)` | `bool` | 带超时的等待 |
| `pg_background_list_v2()` | `SETOF record` | 列出当前会话中已知的工作进程 |
| `pg_background_stats_v2()` | `pg_background_stats` | 会话统计信息（v1.8+） |
| `pg_background_progress(pct, msg)` | `void` | 从工作进程中报告进度（v1.8+） |
| `pg_background_get_progress_v2(pid, cookie)` | `pg_background_progress` | 获取工作进程进度（v1.8+） |

### 取消与分离

| 操作 | 停止执行 | 移除跟踪 |
|------|----------|----------|
| `cancel_v2()` | 是（尽力） | 否 |
| `detach_v2()` | 否 | 是 |

- 使用 **`cancel_v2()`** 来**停止任务**（终止执行，阻止提交）
- 使用 **`detach_v2()`** 来**停止跟踪**（释放簿记资源，工作进程继续运行）

### 工作进程生命周期

```sql
-- 取消运行中的任务
SELECT pg_background_cancel_v2(pid, cookie);

-- 等待完成
SELECT pg_background_wait_v2(pid, cookie);

-- 带超时等待（完成则返回 true）
SELECT pg_background_wait_v2_timeout(pid, cookie, 5000);

-- 列出活跃的工作进程
SELECT * FROM pg_background_list_v2() AS (
  pid int4, cookie int8, launched_at timestamptz,
  user_id oid, queue_size int4, state text,
  sql_preview text, last_error text, consumed bool
);
```

工作进程状态：`running`、`stopped`、`canceled`、`error`

### 进度报告（v1.8+）

```sql
-- 在工作进程 SQL 内部调用
SELECT pg_background_progress(50, 'Halfway done');

-- 从启动方查看进度
SELECT * FROM pg_background_get_progress_v2(pid, cookie);
```


## GUC 配置（v1.8+）

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `pg_background.max_workers` | 16 | 每个会话的最大并发工作进程数 |
| `pg_background.default_queue_size` | 65536 | 默认共享内存队列大小 |
| `pg_background.worker_timeout` | 0 | 工作进程执行超时（0 = 无限制） |


## V1 API（遗留）

V1 API 保留用于向后兼容，但缺少基于 cookie 的 PID 重用保护：

```sql
SELECT pg_background_launch('VACUUM VERBOSE my_table') AS pid \gset
SELECT * FROM pg_background_result(:pid) AS (result TEXT);
SELECT pg_background_detach(:pid);
```

生产环境建议使用 V2 API，因其具备基于 cookie 的 PID 重用保护。


## 安全模型

- 扩展由超级用户安装，默认**不授予 PUBLIC 权限**
- 会创建一个专用的 `pg_background_worker` NOLOGIN 角色
- 辅助函数管理权限：`grant_pg_background_privileges(role, include_v1)`
- 工作进程以**启动用户**（而非超级用户）身份执行
