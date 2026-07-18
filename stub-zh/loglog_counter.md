## 用法

来源：

- [官方扩展控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/loglog/loglog_counter.control)
- [官方上游文档](https://pgxn.org/dist/loglog_estimator/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/loglog_estimator/)

`loglog_counter` — loglog_counter：为 PostgreSQL 提供基于 LogLog 的去重计数估算聚合函数和数据类型。

已复核目录快照记录的版本为 `1.2.4`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "loglog_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'loglog_counter';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
