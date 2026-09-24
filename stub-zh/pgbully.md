## 用法

来源：

- [README](https://github.com/pgElephant/pgbully/blob/259558c49216ea011b7444fe763bebf3bf0ca2ab/README.md)
- [Control file / 控制文件](https://github.com/pgElephant/pgbully/blob/259558c49216ea011b7444fe763bebf3bf0ca2ab/pgbully.control)
- [sql/pgbully--1.0.sql](https://github.com/pgElephant/pgbully/blob/259558c49216ea011b7444fe763bebf3bf0ca2ab/sql/pgbully--1.0.sql)

`pgbully` 使用 Bully 算法，从可达节点中选举编号最大的节点为主节点。它协调独立的 PostgreSQL 节点，不复制应用表，也不实施故障切换或隔离。

### 配置集群

在各 PostgreSQL 15–18 节点上配置唯一节点编号和相同的节点列表，并确保节点间 libpq 认证可用。将库加入预加载列表后重启。节点 1 的配置示例如下：

```conf
shared_preload_libraries = 'pgbully'
pgbully.node_id = 1
pgbully.nodes = '1: host=10.0.0.1 port=5432 dbname=postgres, 2: host=10.0.0.2 port=5432 dbname=postgres, 3: host=10.0.0.3 port=5432 dbname=postgres'
pgbully.heartbeat_interval = '1s'
pgbully.election_timeout = '5s'
```

### SQL 工作流

```sql
CREATE EXTENSION pgbully;
SELECT * FROM pgbully.status();
SELECT node_id, is_leader, reachable, last_seen FROM pgbully.cluster;
SELECT pgbully.is_leader(), pgbully.get_leader(), pgbully.get_term();
```

### 故障边界

安装需要超级用户，状态与成员对象位于 `pgbully` 模式。`pgbully.kv_put` 只能在主节点调用，`pgbully.kv_get` 读取本地状态；这个可选存储不提供法定多数持久性保证。

网络分区后，各分区都可能选出自己的主节点。需要全局唯一写入者的应用必须提供独立的法定多数或隔离机制；主节点检查成功本身不足以授权数据库提升或不可逆的外部操作。
