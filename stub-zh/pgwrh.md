## 用法

来源：

- [Release README](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/README.md)
- [Cluster and rollout guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/overview.md)
- [Core control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh/pgwrh.control)
- [Core SQL API](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh/src/master/api-management.sql)
- [Alpha release boundaries](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/releases/1.0.0-alpha1.md)

`pgwrh` 1.0.0-alpha1 通过将分区叶节点分配给逻辑副本，扩展 PostgreSQL 18 的读取能力。控制节点保存源数据并接收写入；副本查询本地分片，通过外部表访问其余分片。此版本为测试用 alpha，仅支持全新安装，不承诺能够原地升级至后续版本。

### 启用与配置

以管理员身份在控制节点和每个副本中安装。核心扩展依赖 1.6 或更高版本的 `pg_background`、随项目提供的 `pgwrh_fdw` 以及 PL/pgSQL，不依赖原生 `postgres_fdw`。SQL 核心本身无需预加载；可选的复制等待组件有独立的启动要求。

```sql
CREATE EXTENSION pgwrh CASCADE;
```

在控制节点创建复制组并登记已准备好的副本。示例假定已经创建具备复制权限的非超级用户登录角色，授予源表访问权限，并配置网络认证。启动分片放置变更前，还需准备源表的分区层次，并在 `pgwrh.sharded_table` 中登记策略。

```sql
SELECT pgwrh.create_replica_cluster('readers');
SELECT pgwrh.add_replica('readers', 'replica_a', 'replica-a', 5432,
                        _dbname := 'read_a');
```

在每个副本上，通过 `pgwrh.configure_controller` 配置控制节点地址、复制登录角色、认证凭据和控制节点数据库。控制节点与副本可使用不同的数据库名。副本之间的托管连接要求 SCRAM 认证；凭据轮换与拓扑变更分别管理。

### 审查并应用分片放置

复制因子表示符合条件的副本中保存某个分片的百分比，还可设置最少副本数。放置策略也支持考虑可用区。以下调用创建或复用可编辑的待定配置，并预览分片分配：

```sql
SELECT * FROM pgwrh.preview_shard_placement(
    'readers', pgwrh.next_pending_version('readers')
);
SELECT pgwrh.start_rollout('readers');
```

通过 `pgwrh.missing_subscribed_shard`、`pgwrh.missing_connected_local_shard` 和 `pgwrh.missing_ready_remote_shard` 检查阻碍就绪的问题。准备完成后调用 `pgwrh.commit_rollout`，该函数会再次检查就绪条件。放弃变更时使用 `pgwrh.rollback_rollout`，并在下一轮变更前等待副本确认。准备替代路由期间会保留现有副本。

### 一致性与运维

逻辑复制是异步的。读取必须观察到某次已提交写入时使用 `pgwrh_wait`；可选的控制节点控制台由 `pgwrh_ui` 提供。放置状态就绪不能证明副本当前可达或复制延迟为零。跨副本查询没有统一的集群级快照。

模式变更需要人工协调。本版本不会选举替代控制节点，也无法使跨节点的模式变更保持原子性。需要单独维护控制节点备份和 PostgreSQL 高可用方案，分片冗余不能代替这些措施。将主机标记为离线只改变路由，不会删除其数据或停止分配给它的复制。
