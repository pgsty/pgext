## 用法

来源：

- [README.md](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/README.md)
- [pgraft.control](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/pgraft.control)
- [pgraft--2.0.0.sql](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/pgraft--2.0.0.sql)
- [docs/user-guide/configuration.md](https://github.com/pgElephant/pgraft/blob/bc816afee21018a6836356fcc8d2cbf44753be54/docs/user-guide/configuration.md)

`pgraft` 通过 C 与 Go 库提供 Raft 共识、主节点选举和复制键值状态。固定的 2.0.0 源码继承早期 RAM 子项目，目标版本为 PostgreSQL 14–18；它不会自动复制任意应用表。

### 配置与启用

在每个节点预加载 `pgraft`，配置唯一的 `pgraft.name`、顺序相同的 `pgraft.initial_cluster` 列表和一致的 `pgraft.initial_cluster_token`，以及本地 `pgraft.listen_peer_urls` 和持久目录 `pgraft.data_dir`。配置节点互联并重启服务器，然后在数据库中启用 SQL API。

```sql
CREATE EXTENSION pgraft;
SELECT * FROM pgraft.get_cluster_status();
SELECT * FROM pgraft.get_nodes();
SELECT pgraft.is_leader(), pgraft.get_leader();
```

### SQL 操作

`pgraft.kv_put` 与 `pgraft.kv_delete` 在主节点修改键值状态，`pgraft.kv_get` 读取本地状态。`pgraft.add_node` 与 `pgraft.remove_node` 修改成员关系，也必须在主节点调用。`pgraft.log_get_replication_status` 和 `pgraft.log_get_stats` 检查复制状态。配置的持久目录保存日志、快照和 HardState，成员变更与恢复应遵守法定多数要求。

### 升级边界

2.0.0 将旧的非限定函数名移入 `pgraft` 模式并移除原前缀。例如 `pgraft_get_cluster_status` 改为 `pgraft.get_cluster_status`。扩展升级时应同步调整应用调用，升级不级联删除，因此依赖对象可能阻止升级。

```sql
ALTER EXTENSION pgraft UPDATE TO '2.0.0';
```

安装和升级后需要重启，让两个原生库同时重新加载。修改状态的函数不再向 `PUBLIC` 开放，应仅授予所需操作。本文对应固定的 2.0.0 源码，仓库公开发布列表仍显示 1.0 系列。
