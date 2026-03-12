


## 用法

> [pgsentinel: PostgreSQL 活动会话历史](https://github.com/pgsentinel/pgsentinel)

pgsentinel 通过定期采样 `pg_stat_activity` 来记录活动会话历史，并将活动与 `pg_stat_statements` 查询统计进行关联。

### 活动会话历史

```sql
SELECT ash_time, datname, usename, pid, state,
       wait_event_type, wait_event, query, queryid
FROM pg_active_session_history
ORDER BY ash_time DESC;
```

除 `pg_stat_activity` 外的关键列：

| 列名 | 描述 |
|--------|-------------|
| `ash_time` | 采样时间戳 |
| `top_level_query` | 顶层语句（用于 PL/pgSQL） |
| `query` | 包含实际参数值的语句 |
| `cmdtype` | 语句类型：SELECT、UPDATE、INSERT、DELETE、UTILITY、UNKNOWN、NOTHING |
| `queryid` | 关联到 `pg_stat_statements` |
| `blockers` | 阻塞进程数量 |
| `blockerpid` | 阻塞进程的 PID |
| `blocker_state` | 阻塞进程的状态 |

### 查询统计历史

启用后，pgsentinel 还会同时采样 `pg_stat_statements`：

```sql
SELECT ash_time, queryid, calls, total_exec_time, rows,
       shared_blks_hit, shared_blks_read
FROM pg_stat_statements_history
ORDER BY ash_time DESC;
```

### 示例：等待分析

```sql
-- 最近一小时的热点等待事件
SELECT wait_event_type, wait_event, count(*)
FROM pg_active_session_history
WHERE ash_time > now() - interval '1 hour'
  AND wait_event IS NOT NULL
GROUP BY 1, 2
ORDER BY 3 DESC;

-- 阻塞分析
SELECT blockerpid, blocker_state, count(*)
FROM pg_active_session_history
WHERE blockers > 0
GROUP BY 1, 2
ORDER BY 3 DESC;
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pgsentinel_ash.sampling_period` | 1 | 采样周期（秒） |
| `pgsentinel_ash.max_entries` | 1000 | ASH 环形缓冲区大小 |
| `pgsentinel.db_name` | `postgres` | 工作进程连接的数据库 |
| `pgsentinel_ash.track_idle_trans` | `false` | 追踪空闲事务中的会话 |
| `pgsentinel_pgssh.max_entries` | 1000 | pg_stat_statements 历史的环形缓冲区 |
| `pgsentinel_pgssh.enable` | `false` | 启用 pg_stat_statements 历史 |
