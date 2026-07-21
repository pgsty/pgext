## 用法

来源：

- [pgsentinel 1.4.2 README](https://github.com/pgsentinel/pgsentinel/blob/v1.4.2/README.md)
- [pgsentinel 1.4.2 发行版](https://github.com/pgsentinel/pgsentinel/releases/tag/v1.4.2)
- [pgsentinel 1.4.1 到 1.4.2 的更改](https://github.com/pgsentinel/pgsentinel/compare/v1.4.1...v1.4.2)
- [pgsentinel 控制文件](https://github.com/pgsentinel/pgsentinel/blob/v1.4.2/src/pgsentinel.control)

`pgsentinel` 通过在固定时间间隔内采样 `pg_stat_activity` 并将活动与 `pg_stat_statements` 查询统计信息关联来记录活跃会话历史。它最近的样本存储在由后台工作进程管理的共享内存环形缓冲区中。

```ini
shared_preload_libraries = 'pg_stat_statements,pgsentinel'
pg_stat_statements.track = all
pgsentinel.db_name = 'postgres'
```

重启 PostgreSQL，然后在用于工作的数据库中启用这两个扩展：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pgsentinel;
```

### 活跃会话历史

```sql
SELECT ash_time, datname, usename, pid, state,
       wait_event_type, wait_event, query, queryid
FROM pg_active_session_history
ORDER BY ash_time DESC;
```

除了 `pg_stat_activity` 之外的关键列：

| 列 | 描述 |
|--------|-------------|
| `ash_time` | 采样时间戳 |
| `top_level_query` | PL/pgSQL 的顶级语句（如果适用） |
| `query` | 包含实际参数值的语句 |
| `cmdtype` | 语句类型：SELECT, UPDATE, INSERT, DELETE, UTILITY, UNKNOWN, NOTHING |
| `queryid` | 与 `pg_stat_statements` 的链接 |
| `blockers` | 阻塞进程的数量 |
| `blockerpid` | 阻塞进程的 PID |
| `blocker_state` | 阻塞者的状态 |

### 查询统计历史

启用时，pgsentinel 还会并发地采样 `pg_stat_statements`：

```sql
SELECT ash_time, queryid, calls, total_exec_time, rows,
       shared_blks_hit, shared_blks_read
FROM pg_stat_statements_history
ORDER BY ash_time DESC;
```

### 示例：等待分析

```sql
-- Top wait events in the last hour
SELECT wait_event_type, wait_event, count(*)
FROM pg_active_session_history
WHERE ash_time > now() - interval '1 hour'
  AND wait_event IS NOT NULL
GROUP BY 1, 2
ORDER BY 3 DESC;

-- Blocking analysis
SELECT blockerpid, blocker_state, count(*)
FROM pg_active_session_history
WHERE blockers > 0
GROUP BY 1, 2
ORDER BY 3 DESC;
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pgsentinel_ash.sampling_period` | 1 | 每秒采样周期 |
| `pgsentinel_ash.max_entries` | 1000 | ASH 的环形缓冲区大小 |
| `pgsentinel.db_name` | `postgres` | 工作连接的数据库 |
| `pgsentinel_ash.track_idle_trans` | `false` | 跟踪 idle-in-transaction 会话 |
| `pgsentinel_pgssh.max_entries` | 1000 | pg_stat_statements 历史的环形缓冲区大小 |
| `pgsentinel_pgssh.enable` | `false` | 启用 pg_stat_statements 历史 |

### 版本和操作注意事项

- 发行版 1.4.2 修复了 PostgreSQL 17 上的查询统计历史问题，其 `pg_stat_statements` 视图将阻塞 I/O 时间列重命名为 `shared_blk_read_time` 和 `shared_blk_write_time`。
- 同一发行版在不改变 SQL 视图或 GUC 表面的情况下增加了对 PostgreSQL 19 的构建兼容性。
- 环形缓冲区历史是内存驻留且有限的；只有在考虑共享内存使用情况后才能增加条目限制，并在需要更长保留时间时导出样本。
- 查询文本可以包含字面量参数值。当语句可能包括敏感数据时，应限制对历史视图的访问。
