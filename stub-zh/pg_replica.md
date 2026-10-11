## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_replica/pg_replica-0.7.4/README.md)
- [pg_replica.control](https://api.pgxn.org/src/pg_replica/pg_replica-0.7.4/pg_replica.control)
- [Cargo.toml](https://api.pgxn.org/src/pg_replica/pg_replica-0.7.4/Cargo.toml)

`pg_replica` 0.7.4 为 PostgreSQL 物理复制增加内置 Raft 领导者选举与故障转移。它面向 PostgreSQL 18，要求配置集群、预加载并重启；创建扩展本身不会建立复制。

### 核心用法

```ini
shared_preload_libraries = 'pg_replica'
```

```sql
CREATE EXTENSION pg_replica;
SELECT replica.status();
```

### 运行边界

`replica.status()` 显示领导者与复制状态。`replica.rotate_credential()` 在当前主库修改凭据，在备库上不执行变更。进程重启仍由监督程序管理，客户端需要重连与路由逻辑。基于多数派的同步持久性需主动启用，异步模式不能宣称无数据丢失。节点重新加入可能执行 rewind 或完整基础备份，将恢复节点指定为主库可能丢弃原主库的数据。启用后台工作进程前，应核对成员关系、凭据、隔离与恢复流程。安装要求超级用户。
