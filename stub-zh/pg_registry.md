## 用法

来源：

- [官方扩展控制文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_registry/pg_registry.control)
- [官方 Rust 包清单](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_registry/Cargo.toml)
- [官方上游 README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)

`pg_registry` — 存储并验证 JSON Schema，可从表生成模式，并将模式绑定到 Kafka 主题。

已复核目录快照记录的版本为 `0.2.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_registry";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_registry';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
