## 用法

来源：

- [官方扩展控制文件](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/wiki.control)
- [官方项目或服务商页面](https://github.com/lacanoid/pgwiki/wiki)

`wiki` — 提供库内 Wiki 数据模型、SQL/PLPerl API 及可选 MediaWiki 集成的扩展。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plperl`, `plperlu`。
整理后的兼容版本集合为 `10,11,12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "wiki";
SELECT extversion
FROM pg_extension
WHERE extname = 'wiki';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
