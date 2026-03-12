

## 用法

> [pg_squeeze: 从关系中移除未使用空间的工具](https://github.com/cybertec-postgresql/pg_squeeze)

`pg_squeeze` 需要设置 `wal_level = logical`，并且必须添加到 `shared_preload_libraries` 中。它使用逻辑解码而非触发器，在允许并发读写的同时移除表膨胀。

### 注册表进行定期处理

向 `squeeze.tables` 插入数据以启用定期膨胀检查：

```sql
INSERT INTO squeeze.tables (tabschema, tabname, schedule)
VALUES ('public', 'foo', ('{30}', '{22}', NULL, NULL, '{3, 5}'));
```

`schedule` 字段使用类似 crontab 的格式：`(minutes, hours, days_of_month, months, days_of_week)`。上面的配置在每周三和周五的 22:30 检查表 `foo`。

可选列：`free_space_extra`（触发所需的最小额外空闲空间百分比，默认 50）、`min_size`（最小 MB 数，默认 8）、`vacuum_max_age`（距上次 VACUUM 的最大时间，默认 1小时）、`max_retry`（重试次数，默认 0）、`clustering_index`（按此索引排序元组）、`rel_tablespace`、`ind_tablespaces`、`skip_analyze`。

### 临时压缩

```sql
SELECT squeeze.squeeze_table('public', 'pgbench_accounts');
SELECT squeeze.squeeze_table('public', 'mytable', 'my_cluster_idx', 'target_tablespace');
```

### 启动/停止工作进程

```sql
SELECT squeeze.start_worker();   -- 启动调度器和压缩工作进程
SELECT squeeze.stop_worker();    -- 停止当前数据库的所有工作进程
```

通过 `postgresql.conf` 在集群启动时自动启动：

```
squeeze.worker_autostart = 'my_database your_database'
squeeze.worker_role = postgres
```

### 监控

- **`squeeze.log`** -- 每个成功压缩的表对应一条记录（包含 `started`、`finished`、`ins_initial`、`ins`、`upd`、`del`）
- **`squeeze.errors`** -- 压缩过程中的错误
- **`squeeze.get_active_workers()`** -- 显示当前活动的压缩工作进程及其进度

### 配置

- **`squeeze.max_xlock_time`** -- 最大排他锁时间，毫秒（默认不限制）
- **`squeeze.workers_per_database`** -- 每个数据库的并发压缩工作进程数（默认 1）
