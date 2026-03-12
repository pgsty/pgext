

## 用法

> [pgpool_recovery: pgpool-II 的恢复函数](https://pgpool.net/)

`pgpool_recovery` 扩展提供 Pgpool-II 用于后端 PostgreSQL 节点在线恢复的恢复相关函数。

### 函数

```sql
-- 触发后端节点的在线恢复
-- 在主节点上执行恢复脚本
SELECT pgpool_recovery(
    'recovery_1st_stage_script',   -- $PGDATA 中的脚本名称
    'target_hostname',             -- 要恢复节点的主机名
    'target_pgdata',               -- 目标数据目录
    'target_port'                  -- 目标端口号
);

-- 第二阶段恢复（可选，用于流复制）
SELECT pgpool_remote_start(
    'target_hostname',             -- 已恢复节点的主机名
    'target_pgdata'                -- 目标数据目录
);

-- 检查目标节点是否就绪
SELECT pgpool_pgctl('status', 'target_pgdata');
```

### 工作原理

1. Pgpool-II 在主节点上调用 `pgpool_recovery()`，执行用户定义的 Shell 脚本来执行基础备份和设置
2. 恢复脚本将数据复制到目标节点并配置复制
3. `pgpool_remote_start()` 启动已恢复的 PostgreSQL 实例
4. Pgpool-II 将已恢复的节点重新附加到连接池中

### 恢复脚本

恢复脚本（例如 `recovery_1st_stage`）必须放置在主节点的 `$PGDATA` 目录中。它通常执行：

- `pg_basebackup` 将数据复制到目标
- 为流复制配置 `primary_conninfo`
- 在目标上创建 `standby.signal`

该扩展必须安装在 Pgpool-II 管理的所有 PostgreSQL 后端节点上。
