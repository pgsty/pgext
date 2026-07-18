## 用法

来源：

- [官方扩展控制文件](https://github.com/2ndquadrant-it/brin4postgis/blob/fa3259493a7223bb7ffe0daee9d9cce7abc35d6c/brin4postgis.control)

`brin4postgis` — brin4postgis：地理空间类 PostgreSQL 扩展；上游说明为“用于 PostGIS 的 BRIN 索引”。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`postgis`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "brin4postgis";
SELECT extversion
FROM pg_extension
WHERE extname = 'brin4postgis';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
