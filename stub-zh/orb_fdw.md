## 用法

来源：

- [官方扩展控制文件](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/orb_fdw.control)
- [官方上游文档](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/README.md)
- [官方 Rust 包清单](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/Cargo.toml)

`orb_fdw` — orb_fdw：用于访问或集成 Orb 的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `0.13.3`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "orb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'orb_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
