## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/query_recorder/)

`query_recorder` — query_recorder：将已执行 SQL 查询记录到轮转文件，供后续分析或重放

已复核目录快照记录的版本为 `1.0.1`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "query_recorder";
SELECT extversion
FROM pg_extension
WHERE extname = 'query_recorder';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
