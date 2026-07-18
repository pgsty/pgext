## 用法

来源：

- [官方上游文档](https://github.com/hollobon/postgresql_track_renames/blob/dcc35b6031c81f304e11f5c533a1a1a773e5bc2b/README.md)

`track_renames` — 将 PostgreSQL 对象重命名详情转发给可配置回调函数的事件触发器函数。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "track_renames";
SELECT extversion
FROM pg_extension
WHERE extname = 'track_renames';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
