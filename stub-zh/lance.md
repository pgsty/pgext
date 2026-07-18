## 用法

来源：

- [官方扩展控制文件](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/lance.control)
- [官方上游文档](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/README.md)
- [官方 Rust 包清单](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/Cargo.toml)

`lance` — 用于读取和查询 Lance 格式表的 PostgreSQL FDW 扩展。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "lance";
SELECT extversion
FROM pg_extension
WHERE extname = 'lance';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
