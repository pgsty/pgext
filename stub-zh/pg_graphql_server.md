## 用法

来源：

- [官方扩展控制文件](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/pg_graphql_server.control)
- [官方 Rust 包清单](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/Cargo.toml)

`pg_graphql_server` — 由 PostgreSQL 后台工作进程运行 HTTP GraphQL 服务，并将查询交给 pg_graphql 解析。

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`pg_graphql`, `plpgsql`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_graphql_server";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_graphql_server';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
