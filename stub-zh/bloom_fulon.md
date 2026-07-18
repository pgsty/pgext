## 用法

来源：

- [官方扩展控制文件](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/bloom_fulon.control)
- [官方 Rust 包清单](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/Cargo.toml)

`bloom_fulon` — 用于学习目的的 pgrx PostgreSQL 索引访问方法示例。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "bloom_fulon";
SELECT extversion
FROM pg_extension
WHERE extname = 'bloom_fulon';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
