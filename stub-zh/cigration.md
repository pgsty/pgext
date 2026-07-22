## 用法

来源：

- [Official README and operating guide](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/README.md)
- [Version 1.1 control file](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/cigration.control)
- [Version 1.1 SQL implementation](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/cigration--1.1.sql)

`cigration` 1.1 用于编排 Citus 共置分片的在线移动，覆盖扩容、缩容、再均衡和 Worker 替换。它先复制数据，再用逻辑复制追平写入，只在切换元数据时短暂阻塞写操作；旧分片会保留在回收站模式中，直到管理员明确删除。它是高权限集群维护工具，并非应用层迁移框架。

### 前置条件与安装

上游要求 PostgreSQL 10 及以上、Citus 9.2 及以上，并在协调节点安装 `dblink`。所有参与节点都必须设置 `wal_level = logical`；执行角色还必须能在协调节点、Worker、备库及本机之间进行无需交互密码的连接。使用前必须审查这一较强的信任边界。

在协调节点安装时关闭 Citus DDL 传播：

```sql
SET citus.enable_ddl_propagation TO off;
CREATE EXTENSION IF NOT EXISTS dblink;
CREATE EXTENSION cigration;
RESET citus.enable_ddl_propagation;
```

安装后应使用 `cigration.cigration_create_distributed_table`，不能再调用普通的两参数 `create_distributed_table`；扩展会把后者替换为保护函数。

### 再均衡工作流

先创建作业、检查其共置分片任务，再执行：

```sql
SELECT * FROM cigration.cigration_create_rebalance_job();

SELECT jobid, taskid, status, source_nodename, target_nodename,
       all_colocated_logicalrels, total_shard_size
FROM cigration.pg_citus_shard_migration
ORDER BY jobid, taskid;

SELECT cigration.cigration_run_shard_migration_job();
```

旧源分片会被重命名到每个作业专属的回收站模式中。执行破坏性清理前先检查：

```sql
SELECT * FROM cigration.cigration_get_recyclebin_metadata();
SELECT cigration.cigration_cleanup_recyclebin();
```

其他入口包括 `cigration_create_drain_node_job`、`cigration_create_move_node_job`、`cigration_create_general_shard_migration_job`、`cigration_cancel_shard_migration_job` 和 `cigration_cleanup_error_env`。

### 支持的数据与安全边界

上游仅支持单副本、哈希分布表。参考表、多副本分片、append 分布表和 unlogged 表均不支持。一个共置组中的所有表会作为原子任务一起迁移，因此不合理的共置设计可能产生极大的单次迁移。

扩展会安装事件触发器：禁止删除任务涉及的表，并在任意迁移任务运行时禁止 `ALTER TABLE`、`CREATE INDEX`、`ALTER INDEX` 和 `DROP INDEX`。它还会通过 `dblink` 创建和删除 publication、subscription、slot、分片表及 Citus 元数据；应在专用维护窗口使用，并提前演练恢复流程。

没有主键或 replica identity 的表在逻辑传输期间不能并发执行 UPDATE 或 DELETE。PostgreSQL 10 的逻辑复制不传播 TRUNCATE，因此并发 TRUNCATE 可能造成数据分歧。长事务和 HA 切换也会中断作业。修复错误后，仅当所有迁移作业均停止时才能调用 `cigration_cleanup_error_env`；只有验证新分片位置和备份后，才能删除回收站数据。
