## 用法

来源：

- [官方扩展控制文件](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/pg_timers.control)
- [官方上游文档](https://github.com/proventuslabs/pg_timers/blob/21160c1b2d19a274279b02b1f8fed6170f6d2984/README.md)

`pg_timers` — 在 PostgreSQL 中按指定时间或间隔调度 SQL 执行

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_timers";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_timers';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
