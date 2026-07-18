## 用法

来源：

- [固定提交的上游 README](https://github.com/slardiere/streaminglag/blob/8f8b9889f94a1b1685f5078f3540b470fc335e24/README.md)
- [固定提交的扩展 control 文件](https://github.com/slardiere/streaminglag/blob/8f8b9889f94a1b1685f5078f3540b470fc335e24/streaminglag.control)

`streaminglag` 是纯 SQL 辅助扩展，通过 `pg_streaming_lag(text,text)` 或 `pg_streaming_lag(pg_lsn,pg_lsn)` 把 WAL log sequence number 转为整数并计算两个位置之差。它面向旧版 `pg_stat_replication` 列名和 PostgreSQL 9.4 以前的 text LSN 编写。

```sql
CREATE EXTENSION streaminglag;

SELECT pid,
       pg_size_pretty(pg_streaming_lag(sent_lsn, replay_lsn)) AS extension_lag,
       pg_size_pretty(pg_wal_lsn_diff(sent_lsn, replay_lsn)::bigint) AS builtin_lag
FROM pg_stat_replication;
```

不要把扩展结果用于要求正确性的监控。已复核 SQL 使用十六进制 `ff000000` 而不是 `2^32` 乘以 LSN 高半部分，因此高半部分不同时会得到错误差值。现代 PostgreSQL 已提供 `pg_wal_lsn_diff(pg_lsn,pg_lsn)`，当前 replication view 也会直接暴露 LSN position。

该项目是已弃用且无人维护的 2014 年版本 `1.0`，没有 release 或 maintenance history。应优先使用 built-in function；仅为 historical compatibility 保留该扩展，并在移除前逐项比对 migration 期间的结果。
