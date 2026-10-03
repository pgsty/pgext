## 用法

来源：

- [pgsentinel 1.5.1 README](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/README.md)
- [pgsentinel 1.5.1 发行版](https://github.com/pgsentinel/pgsentinel/releases/tag/v1.5.1)
- [1.5.1 upgrade SQL](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/src/pgsentinel--1.5.0--1.5.1.sql)
- [pgsentinel 控制文件](https://github.com/pgsentinel/pgsentinel/blob/v1.5.1/src/pgsentinel.control)

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

### 版本和权限变化

1.5.0 将 `queryid` 与 `nested_queryid` 分开：PostgreSQL 16 及以上的前者来自活动查询标识，后者来自内层解析语句；更早版本中二者一致。升级会重建历史视图，须先检查依赖视图及授权。

1.5.1 撤销 `get_parsedinfo(int)` 的 PUBLIC 执行权限，因为它能返回其他后端的查询文本及字面量。安装新文件后必须应用 SQL 升级，并仅向确有需要的角色授权：

```sql
ALTER EXTENSION pgsentinel UPDATE TO '1.5.1';
```

历史记录位于有限的共享内存环形缓冲区中，需要长期保留时应另行导出。采样查询文本可能包含敏感值，还应检查历史视图的访问权限。上游 CI 覆盖 PostgreSQL 10–19。
