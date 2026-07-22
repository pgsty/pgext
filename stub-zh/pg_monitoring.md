## 用法

来源：

- [官方 README](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/README)
- [0.2 版 SQL](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/pg_monitoring--0.2.sql)
- [发行说明](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/Changes)
- [扩展控制文件](https://github.com/einhverfr/pgxn-postgres_monitoring/blob/b2b877bcca7ccc129f281c96587248e9ef5398a5/pg_monitoring.control)

`pg_monitoring` 0.2 是一组历史性的纯 SQL 复制延迟与聚合负载函数。它使用 PostgreSQL 10 之前的统计字段和 WAL 函数名，表负载计算也存在数据质量缺陷。它可以作为可审查示例，不能作为当前监控契约。

### 已安装函数

```sql
CREATE EXTENSION pg_monitoring;

SELECT * FROM pg_monitoring_lag_info();
SELECT pg_monitoring_time_since_replay();
SELECT * FROM pg_monitoring_raw_table_load();
SELECT * FROM pg_monitoring_load_across_tables();
SELECT * FROM pg_monitoring_load_across_databases();
```

- `pg_monitoring_lag_info()` 读取 `pg_stat_replication` 并估算字节延迟。
- `pg_monitoring_time_since_replay()` 报告备库距最后一次事务重放的秒数。
- `pg_monitoring_raw_table_load()` 聚合表扫描与修改计数器。
- `pg_monitoring_load_across_tables()` 修改一行基线表，并返回相对上次调用的增量。
- `pg_monitoring_load_across_databases()` 返回当前集群范围的后端、事务、读取与命中总和；它不是增量。

### 版本与正确性缺陷

延迟函数引用 `replay_location` 和 `pg_current_xlog_location()`；PostgreSQL 10 已将其替换为 `replay_lsn` 与 `pg_current_wal_lsn()`。它还用基于 255 的乘数换算 WAL 高位，而 32 位十六进制边界需要 256。不要指望该函数在当前 PostgreSQL 上能执行或报告正确字节数。

`pg_monitoring_raw_table_load()` 从 `n_tup_upd` 而非 `n_tup_del` 取得删除数，因此删除数错误，更新还在总量中被重复计算。增量函数被文档明确标为不支持并发调用：每次轮询都会替换唯一的共享基线，所以两个采集器、一次失败抓取、统计重置或重启都会产生误导区间。安装表 `pg_monitoring_last_run` 没有被这些函数使用。

### 安全迁移

在针对目标大版本重写并测试 SQL 前，不要把这些结果用于告警或服务级指标。应优先直接使用当前统计视图与 `pg_wal_lsn_diff` 等 LSN 辅助函数，让监控系统按目标维护时间戳、基线、重置检测与速率计算。只授予采集器必需的监控权限，并用已知负载验证主库/备库 NULL 行为、故障切换、统计重置、计数器回绕/清零、并发抓取与版本升级。
