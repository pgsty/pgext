## 用法

来源：

- [官方扩展控制文件](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/superloglog/superloglog_counter.control)
- [官方上游文档](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/superloglog/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/superloglog_estimator/)

`superloglog_counter` — 基于 SuperLogLog 算法的 C 数据类型、函数与近似去重计数聚合扩展。

已复核目录快照记录的版本为 `1.2.3`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "superloglog_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'superloglog_counter';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
