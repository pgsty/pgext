

## 用法

> [repmgr: PostgreSQL 复制管理器](https://github.com/EnterpriseDB/repmgr)

一套用于管理 PostgreSQL 集群复制和故障转移的工具。支持搭建备库服务器、监控复制状态以及执行故障转移/计划切换。

### 启用

```sql
CREATE EXTENSION repmgr;
```

### 配置

在每个节点上创建 `repmgr.conf`：

```ini
node_id=1
node_name='node1'
conninfo='host=node1 dbname=repmgr user=repmgr'
data_directory='/var/lib/postgresql/data'
```

### 注册主节点

```bash
repmgr -f /etc/repmgr.conf primary register
```

### 克隆并注册备节点

```bash
# 从主节点克隆
repmgr -h primary-host -U repmgr -d repmgr -f /etc/repmgr.conf standby clone

# 启动备节点
pg_ctl -D /var/lib/postgresql/data start

# 注册备节点
repmgr -f /etc/repmgr.conf standby register
```

### 监控

```bash
# 显示集群状态
repmgr -f /etc/repmgr.conf cluster show

# 显示复制状态
repmgr -f /etc/repmgr.conf node status
```

### 手动切换

将备节点提升为主节点（计划性切换）：

```bash
repmgr -f /etc/repmgr.conf standby switchover
```

### 使用 repmgrd 自动故障转移

在每个节点上启动 repmgr 守护进程：

```bash
repmgrd -f /etc/repmgr.conf
```

在 `repmgr.conf` 中配置故障转移：

```ini
failover='automatic'
promote_command='repmgr standby promote -f /etc/repmgr.conf'
follow_command='repmgr standby follow -f /etc/repmgr.conf --upstream-node-id=%n'
```

### 关键命令

- `repmgr primary register` - 注册主节点
- `repmgr standby clone` - 从主节点克隆备节点
- `repmgr standby register` - 注册备节点
- `repmgr standby promote` - 将备节点提升为主节点
- `repmgr standby follow` - 跟随新的主节点
- `repmgr standby switchover` - 计划性切换
- `repmgr cluster show` - 显示集群状态
- `repmgr cluster event` - 显示集群事件
- `repmgr node check` - 节点健康检查
