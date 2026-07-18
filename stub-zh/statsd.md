## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/statsd/)

`statsd` — 通过 UDP 向 StatsD 发送计数器、仪表与计时指标的 C 客户端函数。

已复核目录快照记录的版本为 `0.2.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "statsd";
SELECT extversion
FROM pg_extension
WHERE extname = 'statsd';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
