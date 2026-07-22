## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/README.md)
- [1.0 版本 SQL 对象](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots--1.0.sql)
- [启动器实现](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots_launcher.c)

`synchronize_logical_slots` 是一个已放弃、仅适用于 EPAS 12 的机制，用于把逻辑复制槽状态复制到同步流复制备库。它依赖 `dblink`，安装后台工作进程启动器，并且必须在主库和每个同步备库上构建安装。

### 安装边界

```sql
CREATE EXTENSION synchronize_logical_slots CASCADE;
GRANT EXECUTE ON FUNCTION primary_checkpoint() TO replication_agent;
```

在主库和备库的 `shared_preload_libraries` 中配置 `$libdir/synchronize_logical_slots_launcher`。启动器默认连接 `postgres`；若扩展与逻辑槽位于其他数据库，应把 `sync_logical_slot.database` 设置为该数据库。需要在 `max_worker_processes` 中预留容量，在备库启用 `hot_standby_feedback`，向复制用户授予 `pg_read_all_stats`，添加所需 `pg_hba.conf` 规则，并重启所有节点。工作进程大约每 60 秒轮询一次。

### 破坏性与安全边界

在同步备库上，对账会创建或推进缺失槽，删除并重建插件不同的槽，还会删除主库列表中不存在的本地逻辑槽。不要在受管备库上混放无关逻辑槽，启用工作进程前必须测试精确的槽清单。提升后，上游仍要求手工删除不再使用的槽；陈旧槽可能无限保留 WAL。

`primary_checkpoint()` 是 `SECURITY DEFINER` 函数，并调用未限定模式的扩展函数。只应向严格受控的复制角色授权，并在使用前固定函数搜索路径。同步函数会捕获错误并以文本返回，因此仅看到启动器活动并不能证明槽已成功推进；必须监控槽 LSN 和服务器日志。

该项目只支持同步备库，并且早于 PostgreSQL 持续维护的故障转移槽机制。应优先使用对应版本的原生机制。如果遗留 EPAS 12 被迫使用，应在类生产负载下验证槽推进、故障转移、网络分区、无关槽保留、重复提升、WAL 保留及不一致状态恢复。
