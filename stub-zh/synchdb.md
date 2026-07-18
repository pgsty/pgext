## 用法

来源：

- [官方上游文档](https://docs.synchdb.com)

`synchdb` — 通过 C/Java 混合模块，将 MySQL、SQL Server、Oracle 与 OpenLog Replicator 数据直接以 CDC 方式复制到 PostgreSQL。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pgcrypto`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "synchdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'synchdb';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
