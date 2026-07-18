## 用法

来源：

- [官方扩展控制文件](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/schedule.control)

`schedule` — PostgreSQL 的 cron 格式 schedule 类型及 UTC 时间匹配函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "schedule";
SELECT extversion
FROM pg_extension
WHERE extname = 'schedule';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
