## 用法

来源：

- [官方扩展控制文件](https://github.com/eva-ics/eva4/blob/f0e6801d2bb79d6651487766e3a3843fc50e2c2b/eva_pg/eva_pg.control)
- [官方 Rust 包清单](https://github.com/eva-ics/eva4/blob/f0e6801d2bb79d6651487766e3a3843fc50e2c2b/eva_pg/Cargo.toml)
- [官方上游 README](https://github.com/eva-ics/eva4/blob/f0e6801d2bb79d6651487766e3a3843fc50e2c2b/README.md)

`eva_pg` — EVA ICS 工业物联网平台的 pgrx PostgreSQL 扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "eva_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'eva_pg';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
