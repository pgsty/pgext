## 用法

来源：

- [官方上游文档](https://pgxn.org/dist/hyperloglog_estimator/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/hyperloglog_estimator/)
- [官方上游 README](https://github.com/tvondra/distinct_estimators/blob/248ffd3eb601785b5c6752707f4e054839bccfba/README.md)

`hyperloglog_counter` — 基于 HyperLogLog 的去重估算聚合函数和数据类型。

已复核目录快照记录的版本为 `1.2.6`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "hyperloglog_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'hyperloglog_counter';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
