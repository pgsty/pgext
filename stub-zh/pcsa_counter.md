## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/pcsa_estimator/pcsa_estimator-1.3.3/pcsa_counter.control)
- [官方上游文档](https://pgxn.org/dist/pcsa_estimator/1.3.3/README.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/pcsa_estimator/)

`pcsa_counter` — 基于 PCSA 算法的去重计数估算数据类型、函数与聚合。

已复核目录快照记录的版本为 `1.3.3`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pcsa_counter";
SELECT extversion
FROM pg_extension
WHERE extname = 'pcsa_counter';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
