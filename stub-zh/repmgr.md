## 用法

来源：

- [README v5.5.1](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/README.md)
- [Control](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/repmgr.control)
- [SQL 5.5](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/repmgr--5.5.sql)
- [Quickstart](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/doc/quickstart.xml)
- [repmgrd configuration](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/doc/repmgrd-configuration.xml)
- [Upgrade guide](https://github.com/EnterpriseDB/repmgr/blob/v5.5.1/doc/upgrading-repmgr.xml)
- [5.5.1 release](https://github.com/EnterpriseDB/repmgr/releases/tag/v5.5.1)

`repmgr` 管理 PostgreSQL 物理复制、监控与故障切换。发行版 5.5.1 支持 PostgreSQL 13–18，SQL 扩展版本仍为 5.5；它修复了向外部通知命令替换参数时的转义问题。集群各节点应使用同一 repmgr 发布版本。

### 启用

使用独立的元数据数据库与管理角色，并配置复制连接及所需的 SSH 访问。上游快速入门使用超级用户角色；普通角色需要为高权限操作提供 `--superuser`。以下 SQL 应在元数据数据库内以足够权限执行：

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
# Clone from primary
repmgr -h primary-host -U repmgr -d repmgr -f /etc/repmgr.conf standby clone

# Start the standby
pg_ctl -D /var/lib/postgresql/data start

# Register the standby
repmgr -f /etc/repmgr.conf standby register
```

### 监控

```bash
# Show cluster status
repmgr -f /etc/repmgr.conf cluster show

# Show replication status
repmgr -f /etc/repmgr.conf node status
```

### 手动切换

将备节点提升为主节点（计划性切换）：

```bash
repmgr -f /etc/repmgr.conf standby switchover
```

### 使用 repmgrd 自动故障转移

使用 `repmgrd` 前，在每个节点配置以下参数并重启 PostgreSQL：

```ini
shared_preload_libraries = 'repmgr'
```

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

### 5.5.1 升级

按升级指南暂停或停止自动故障转移，更新各节点的程序，再重启正在运行的 `repmgrd` 并检查集群状态。5.5.0 与 5.5.1 的 SQL 版本均为 `5.5`，不要尝试更新到不存在的 SQL 版本 5.5.1；从更早的 SQL 版本升级时，应遵循对应的扩展升级链。
