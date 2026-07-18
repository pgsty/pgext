## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/README.md)
- [1.0 版本 SQL 对象](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots--1.0.sql)
- [启动器实现](https://github.com/vibhorkum/synchronize-logical-slots/blob/aea7531d9516b8c0315a00307757198f2335510b/synchronize_logical_slots_launcher.c)

`synchronize_logical_slots` 是一个已放弃、仅适用于 EPAS 12 的机制，用于把逻辑复制槽状态复制到同步流复制备库。它依赖 `dblink`，安装后台工作进程启动器，并且必须在主库和每个同步备库上构建安装。

```sql
CREATE EXTENSION synchronize_logical_slots CASCADE;
GRANT EXECUTE ON FUNCTION primary_checkpoint() TO replication_agent;
```

在主库与备库的 `shared_preload_libraries` 中配置 `$libdir/synchronize_logical_slots_launcher`，启用 `hot_standby_feedback`，向复制用户授予 `pg_read_all_stats`，添加所需 `pg_hba.conf` 规则，并重启所有节点。应严格限制复制角色和 `primary_checkpoint()` 授权。

该项目只处理同步备库。提升后，上游要求人工删除不用的逻辑槽；遗留槽会无限保留 WAL。它早于 PostgreSQL 持续维护的故障转移槽机制，也不支持当前版本，应优先使用对应版本的原生机制。如果遗留 EPAS 12 被迫使用，应在类生产负载下验证槽位推进、故障转移、网络分区、重复提升、WAL 保留及不一致状态恢复。
