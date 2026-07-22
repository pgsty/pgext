## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/README.md)
- [1.0 版本 SQL 对象](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats--1.0.sql)
- [扩展 control 文件](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats.control)

`pg_stats` 是一组纯 SQL 的便捷监控视图。1.0 版本会创建 `pg_stat_tables`、`pg_stat_indexes`、`pg_stat_users`、`pg_stat_queries`、`pg_stat_long_trx` 和 `pg_stat_waiting_locks`。

### 核心流程

```sql
CREATE EXTENSION pg_stats;
SELECT schemaname, relname, idx_scan_ratio, hit_ratio, total_size
FROM pg_stat_tables
ORDER BY total_size DESC;
```

### 视图索引与限制

`pg_stat_tables` 增加扫描、缓存命中、DML、HOT 更新与大小比例；`pg_stat_indexes` 增加索引命中率和大小；`pg_stat_users` 汇总后端登录时长，`pg_stat_queries` 与 `pg_stat_long_trx` 显示事务时长，`pg_stat_waiting_locks` 报告尚未授予的锁。

这些只是派生快照，不是持久测量。若干百分比会先做整数除法再舍入，`pg_stat_indexes` 还按表 OID 而非索引 OID 关联 I/O 行，因此拥有多个索引的表可能产生重复或错配的索引指标。1.0 版本还使用已移除的 `pg_stat_activity.waiting` 列。每个 PostgreSQL 大版本都应检查并修正 SQL，并限定视图名称以免和核心 `pg_stats` 系统视图混淆；告警与长期趋势应使用持续维护的监控系统。
