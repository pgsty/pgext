

## 用法

> [pgautofailover: PostgreSQL 自动故障转移](https://github.com/hapostgres/pg_auto_failover)

pg_auto_failover 是一个用于 PostgreSQL 自动故障转移的扩展和服务。它由一个监控节点（运行 `pgautofailover` 扩展）和每个数据节点上由 `pg_autoctl` CLI 管理的 keeper 进程组成。

### 架构

- **监控节点**：安装了 `pgautofailover` 扩展的 PostgreSQL 实例，实现了用于故障转移决策的状态机
- **Keeper**（`pg_autoctl run`）：运行在每个数据节点上，向监控节点报告健康状态并执行状态转换
- 支持 2 个以上节点的设置，默认使用同步复制
- 支持 Citus HA（v2.0 起）

### 关键 CLI 命令

```bash
# 创建监控节点
pg_autoctl create monitor --pgdata /path/to/monitor --pgport 5000

# 创建数据节点（自动分配为主节点或从节点）
pg_autoctl create postgres --pgdata /path/to/data --pgport 5001 --monitor postgres://monitor:5000/pg_auto_failover

# 运行 keeper（前台模式）
pg_autoctl run --pgdata /path/to/data

# 检查集群状态
pg_autoctl show state --pgdata /path/to/monitor

# 执行手动切换
pg_autoctl perform switchover --pgdata /path/to/monitor

# 执行手动故障转移
pg_autoctl perform failover --pgdata /path/to/monitor
```

### 故障转移行为

监控节点跟踪节点健康状态。当从节点不可用或其延迟过高时，会从主节点的 `synchronous_standby_names` 中移除。在从节点恢复健康之前，故障转移/切换操作会被阻止，以防止数据丢失。

### 文档

完整文档：[pg-auto-failover.readthedocs.io](https://pg-auto-failover.readthedocs.io/en/main/)
