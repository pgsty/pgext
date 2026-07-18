## 用法

来源：

- [官方扩展控制文件](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/pg_ivfflat.control)
- [官方上游文档](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/README.md)
- [官方 Rust 包清单](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/Cargo.toml)

`pg_ivfflat` — 基于 pgrx 的实验性扩展，提供 Vector 类型和内存 IVFFlat 索引访问方法。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_ivfflat";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ivfflat';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
