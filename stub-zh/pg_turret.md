## 用法

来源：

- [官方扩展控制文件](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/pg_turret.control)
- [官方上游文档](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/README.md)
- [官方 Rust 包清单](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/Cargo.toml)

`pg_turret` — 基于 pgrx 的 PostgreSQL 日志捕获与实时导出扩展，可通过后台工作进程发送到 HTTP、Kafka、Sentry 等目标。

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_turret";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_turret';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
