## 用法

来源：

- [官方扩展控制文件](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/pg_flashback.control)
- [官方上游文档](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/README.md)
- [官方 Rust 包清单](https://github.com/CaghanTU/pg_flashback/blob/7f0a26d763372859826f8e21d9b1594dce8dc414/Cargo.toml)

`pg_flashback` — 通过触发器或 WAL 捕获，为 PostgreSQL 提供表级时间点恢复和历史查询。

已复核目录快照记录的版本为 `0.1.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_flashback";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_flashback';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
