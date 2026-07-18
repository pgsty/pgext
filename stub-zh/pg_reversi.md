## 用法

来源：

- [官方上游文档](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/README.md)

`pg_reversi` — 在 PostgreSQL 9.6 中实现 Reversi 黑白棋游戏的示例扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "pg_reversi";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_reversi';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
