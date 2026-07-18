## 用法

来源：

- [官方上游文档](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/README.md)
- [官方 Rust 包清单](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/Cargo.toml)
- [官方项目或服务商页面](https://github.com/Creatif/creatif-backend/tree/master/pgx_ulid)

`ulid` — 内嵌的 pgrx ULID 类型、生成器与转换函数

已复核目录快照记录的版本为 `0.1.3`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ulid";
SELECT extversion
FROM pg_extension
WHERE extname = 'ulid';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
