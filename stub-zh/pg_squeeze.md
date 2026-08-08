## 用法

来源：

- [pg_squeeze REL1_9_4 发行版](https://github.com/cybertec-postgresql/pg_squeeze/releases/tag/REL1_9_4)
- [pg_squeeze REL1_9_4 README](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_4/README.md)
- [pg_squeeze 发行说明](https://github.com/cybertec-postgresql/pg_squeeze/blob/REL1_9_4/NEWS)

`pg_squeeze` 可以在允许并发读写的同时清除表及其索引中的膨胀。它将存活元组复制到新存储，并通过逻辑解码应用并发变更，从而避免 `VACUUM FULL` 的长时间排他锁。只有在规划好复制槽、磁盘空间和表的副本标识后才应使用。

### 配置与安装

```conf
max_replication_slots = 1  # or add one to the existing requirement
shared_preload_libraries = 'pg_squeeze'
wal_level = logical       # required on PostgreSQL versions before 19
```

重启 PostgreSQL，然后创建扩展：

```sql
CREATE EXTENSION pg_squeeze;
```

表必须有标识索引。主键可配合默认副本标识使用；否则，请使用 `ALTER TABLE ... REPLICA IDENTITY USING INDEX` 选择合适的唯一索引。

### 执行临时压缩

```sql
SELECT squeeze.squeeze_table('public', 'pgbench_accounts');

SELECT squeeze.squeeze_table(
  'public',
  'large_table',
  'large_table_cluster_idx',
  'target_tablespace'
);
```

该函数会启动后台任务，并不具备普通 SQL 函数意义上的事务性。请监控操作，不要假定外围的 `ROLLBACK` 会取消它。

### 调度表并监控任务

```sql
INSERT INTO squeeze.tables (tabschema, tabname, schedule)
VALUES ('public', 'events', ('{30}', '{22}', NULL, NULL, '{3,5}'));

SELECT * FROM squeeze.get_active_workers();
SELECT * FROM squeeze.log ORDER BY finished DESC;
SELECT * FROM squeeze.errors;
```

调度元组依次包含分钟、小时、月中日期、月份和星期。注册还支持阈值与放置选项，例如 `free_space_extra`、`min_size`、`vacuum_max_age`、`max_retry`、`clustering_index`、关系/索引表空间，以及 `skip_analyze`。

如需自动启动：

```conf
squeeze.worker_autostart = 'my_database'
squeeze.worker_role = 'postgres'
```

### 版本 1.9.4 与运维注意事项

- 版本 1.9.4 修复了动态构造的 `ANALYZE`、日志和错误语句中的不安全引用问题，其中包括一条超级用户 SQL 注入路径。应尽快升级早期 1.9 版本。
- 全表压缩所需的空闲磁盘空间约为目标表及其索引总大小的两倍。
- 破坏性 DDL、`VACUUM FULL`、`CLUSTER` 或 `TRUNCATE` 可能使正在进行的压缩中止。请协调模式变更并审慎设置 `max_retry`。
- 与其他在线重写工具类似，`pg_squeeze` 会改变行可见性；对于持有旧快照的并发会话，它存在文档明确说明的 MVCC 注意事项。
- 在对包含该扩展的数据库执行 `pg_upgrade` 或转储/恢复之前，请先在新集群的 `shared_preload_libraries` 中配置 `pg_squeeze`。
- 当前 Pigsty 软件包覆盖 PostgreSQL 14-18。对于这些版本，请保持 `wal_level = logical`；上游针对 PostgreSQL 19 放宽的规则尚不适用于该软件包矩阵。
