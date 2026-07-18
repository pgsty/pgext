## 用法

来源：

- [官方扩展控制文件](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/postgres_redis.control)
- [官方上游文档](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/README.md)
- [官方 Rust 包清单](https://github.com/systemEng-Learning/postgres-redis/blob/05ec5172157932635c1f773fd49d8b61dd13a948/Cargo.toml)

`postgres_redis` — 通过执行器钩子与后台工作进程将选定 PostgreSQL 表值同步到 Redis。

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postgres_redis";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_redis';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
