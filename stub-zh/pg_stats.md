## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/README.md)
- [1.0 版本 SQL 对象](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats--1.0.sql)
- [扩展 control 文件](https://github.com/s-hironobu/pg_stats/blob/c7753215e4070a88379775470c888d5eed049d17/pg_stats.control)

`pg_stats` 是一组纯 SQL 的便捷监控视图。它把 PostgreSQL 统计信息与大小函数组合成 `pg_stat_tables`、`pg_stat_indexes`、`pg_stat_users` 和 `pg_stat_queries`。

```sql
CREATE EXTENSION pg_stats;
SELECT schemaname, relname, idx_scan_ratio, hit_ratio, total_size
FROM pg_stat_tables
ORDER BY total_size DESC;
```

`pg_stat_tables` 增加扫描、缓存命中、DML、HOT 更新与大小比例；`pg_stat_indexes` 增加索引命中率和大小；用户与查询视图则汇总活动后端和运行中查询时长。

这些只是派生快照，不是持久测量。统计重置后、分母很小时，或底层视图未覆盖某些活动时，比率可能为空或产生误导。1.0 版本针对较老的统计字段编写，例如历史 waiting 标志。每个 PostgreSQL 大版本都应检查 SQL，并限定视图名称以免和核心 `pg_stats` 系统视图混淆；告警与长期趋势应使用持续维护的监控系统。
