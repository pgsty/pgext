## 用法

来源：

- [官方上游文档](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/README.md)

`pg_microbench` — pg_microbench：使用 Linux perf_event_open 统计查询规划、执行、工具命令和 SPI 查询阶段的硬件性能计数器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_microbench";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_microbench';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
