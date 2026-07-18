## 用法

来源：

- [官方扩展控制文件](https://github.com/umitanuki/tinyint-postgresql/blob/76495ba27565df2e6758e8f26778e16e963004fb/tinyint.control)
- [官方上游文档](https://pgxn.org/dist/tinyint/doc/tinyint.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/tinyint/)

`tinyint` — 一字节有符号整数类型，提供算术、类型转换、聚合以及 B-tree、哈希和 GIN 操作符类。

已复核目录快照记录的版本为 `0.1.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tinyint";
SELECT extversion
FROM pg_extension
WHERE extname = 'tinyint';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
