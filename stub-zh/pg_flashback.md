## 用法

来源：

- [Official extension control file](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/pg_flashback.control)
- [Official upstream documentation](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/README.md)
- [Official Rust package manifest](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/Cargo.toml)

`pg_flashback` 0.1.0 是一个 Rust 与 pgrx 扩展，用于表级时间点还原和只读历史查询。它通过触发器或异步逻辑 WAL 消费捕获 DML，通过 utility hook 记录 DDL，再从快照与增量事件重建目标表。它不能替代集群备份或 PostgreSQL PITR。

### 核心流程

预加载模块并重启 PostgreSQL。WAL 捕获还要求逻辑 WAL；触发器模式无需该设置也能工作。

```conf
wal_level = logical
shared_preload_libraries = 'pg_flashback'
```

```sql
CREATE EXTENSION pg_flashback;

SELECT flashback_track('public.orders');
SELECT flashback_checkpoint('public.orders');

SELECT flashback_restore(
  'public.orders',
  now() - interval '30 seconds'
);
```

`flashback_track(table)` 创建捕获对象、基础快照和模式元数据。`flashback_untrack(table)` 停止跟踪并删除这段历史。`flashback_restore(table, timestamptz)` 把表替换成指定时刻的状态；数组重载可按外键顺序还原多张表。

### 查询与恢复 API

`flashback_query(table, timestamptz, filter_clause)` 在临时表中重建历史状态并返回 `SETOF record`，因此调用者必须提供列定义列表。过滤条件只能是 `WHERE` 谓词。`flashback_recover_deleted(table, timestamptz)` 重新插入当前已缺失的行，并要求表有主键。`flashback_restore_parallel(table, timestamptz, num_workers)` 提供并行查询提示，并不并行重放增量。

用 `flashback_retention_status()`、`flashback.pg_stat_flashback`、`flashback_history(table, interval)` 和 `flashback.restore_log` 查看保留状态、待处理工作、变更历史与还原审计。

### 配置与安全

`pg_flashback.capture_mode` 选择 `auto`、`wal` 或 `trigger`；`pg_flashback.enabled` 是全局捕获开关。其余行为由工作数据库、间隔、批量、行大小、保留期、复制槽和还原内存参数控制。WAL 模式依赖异步消费的逻辑复制槽，触发器模式则增加同步 DML 开销。还原会重建影子表并原子交换，期间短暂持有排他锁；依赖对象、RLS、序列、分区、继承、崩溃恢复、工作进程延迟、保留期和磁盘增长都必须经过测试。

官方仓库和固定版本文件目前返回 HTTP 404。本文反映的是已审阅的上游快照；采用前必须重新确认源码可用性与维护状态。
