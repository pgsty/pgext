## 用法

来源：

- [官方 README](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/README.md)
- [主后台工作进程源码](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/pg_keeper.c)
- [备库提升源码](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/standby.c)

`pg_keeper` 是一个传统的双节点 PostgreSQL 集群模块。每个节点上预加载的后台工作进程轮询对端：主库在备库连续失败后从同步复制切换为异步复制，备库则在主库连续失败后自行提升。它不是共识系统，也不能单独防止脑裂。

### 配置边界

上游仓库只构建共享模块，没有扩展控制文件或版本化 SQL。应在已经正常运行的主备对上配置并预加载它，而不是执行 `CREATE EXTENSION`：

```conf
shared_preload_libraries = 'pg_keeper'
max_worker_processes = 8

pg_keeper.keepalives_time = 5
pg_keeper.keepalives_count = 3
pg_keeper.my_conninfo = 'host=127.0.0.1 port=5432 dbname=postgres'
pg_keeper.partner_conninfo = 'host=partner port=5432 dbname=postgres'
pg_keeper.after_command = '/usr/local/sbin/verified-fencing-hook'
```

源码把间隔与次数设置拼写为 `keepalives_time` 和 `keepalives_count`，必须保留确切名称。`partner_conninfo` 用于心跳查询，`my_conninfo` 用于本机 `ALTER SYSTEM` 操作；两个连接串都是必填项。

### 故障行为

同步主库的心跳失败次数达到阈值后，`pg_keeper` 通过 `ALTER SYSTEM` 清空 `synchronous_standby_names`，使事务可以异步继续。备库恢复后重新启用同步复制需要运维人员操作。

备库达到同一阈值后，会创建历史提升触发并提升本机。可选的 `after_command` 在提升后运行，用于集成隔离或 STONITH。仅检查数据库可达性无法区分主库故障与网络分区，因此可靠的外部隔离至关重要。

每次重启后都要在两个节点确认工作进程状态与日志状态转换。近似检测时间为轮询间隔乘以失败次数，另加连接超时与调度开销。

### 兼容性与安全性

README 把自身发行版称为 1.0，而目录条目报告 2.0；上游仓库没有控制文件证明存在该扩展版本。应记录这一差异，不能假定目录 2.0 代表更新的上游 API。

项目只在 CentOS 6.5 的 PostgreSQL 9.5 与 9.6 上测试，最后变更于 2019 年。文档使用已经移除的恢复配置，提升代码也针对旧服务器行为。未经完整源码移植与破坏性故障转移测试，不能部署到现代 PostgreSQL。应优先使用带分布式共识和成熟隔离机制的维护中 HA 系统。连接凭据、`ALTER SYSTEM` 权限、提升、同步复制丢失、回切、时间线分叉和双主恢复都需要明确的运维设计。
