## 用法

来源：

- [Official README](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/README.md)
- [Control file](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/drift_detect.control)
- [Version 1.0 SQL](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/drift_detect--1.0.sql)
- [Planner hook and configuration](https://github.com/kvfi/pg-drift/blob/663a10971ae1bb1549004d2c4d100906b1cde88a/drift_detect.c)

`drift_detect` 检测已登记查询的执行计划结构变化，保存事件并通知监听者，例如报告索引扫描变为顺序扫描。所引用的源码快照声明扩展版本为 1.0，支持 PostgreSQL 14–17；仓库尚未发布正式发行版本。

### 启用并登记查询

将库追加到 `shared_preload_libraries`，保留原有条目，并选择后台工作进程要监控的唯一数据库。重启 PostgreSQL 后，以管理员身份在该数据库中创建扩展。管理函数需要 PL/pgSQL。

```ini
shared_preload_libraries = 'drift_detect'
drift_detect.database = 'appdb'
```

```sql
CREATE EXTENSION drift_detect;
```

对于已经存在的订单表，按应用实际发送的参数化 SQL 登记查询，并提供用于基线计划的代表性参数。返回的整数用于标识被跟踪的查询。

```sql
SELECT drift.track_query(
    'get user orders',
    'SELECT * FROM orders WHERE user_id = $1 ORDER BY created_at DESC LIMIT 20',
    '{"params": [1]}'::jsonb
);
LISTEN plan_drift;
SELECT * FROM drift.drift_summary();
```

### 对象与计划匹配

`drift.track_query` 登记查询及其初始基线，可选的第四个参数用于指定通知频道。`drift.untrack_query` 停用跟踪。`drift.acknowledge_drift` 接收事件 ID 和处理者标签，将事件标记为已处理，不会修改查询。`drift.drift_summary` 列出尚未确认的事件。

`drift.tracked_queries`、`drift.plan_snapshots` 和 `drift.plan_drift_events` 分别保存登记信息、基线与事件。`drift.reload_cache` 重新加载已启用的登记项。`drift.compute_query_id` 和 `drift.fingerprint` 是诊断辅助函数。匹配依据 PostgreSQL 规范化查询 ID；结构指纹包含计划节点、关系、索引和策略，不包含成本及行数估算。

### 配置与边界

`drift_detect.enabled` 默认为 true。`drift_detect.snapshot_interval_secs` 默认为 60 秒，`drift_detect.notify_channel` 默认为 `plan_drift`，`drift_detect.min_table_growth_ratio` 默认为 3.0。这些设置在重新加载配置后生效。`drift_detect.max_tracked` 默认为 500；修改它或 `drift_detect.database` 都需要重启。

后台工作进程只监控一个数据库，部署环境必须允许安装和预加载该库。SQL 管理函数使用调用者权限；为非管理员用户显式安排模式、表和函数的访问权限。通知表示计划结构发生变化，不能直接证明性能退化。应保留持久事件用于排查，并按工作负载管理保留周期；此源码快照未说明自动保留清理机制。
