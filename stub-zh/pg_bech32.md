## 用法

来源：

- [官方扩展控制文件](https://github.com/KaiserKarel/pg_bech32/blob/main/pg_bech32.control)
- [官方上游文档](https://github.com/KaiserKarel/pg_bech32/blob/main/README.md)
- [官方 Rust 包清单](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/Cargo.toml)

`pg_bech32` — 为 PostgreSQL 增加 Bech32、Bech32m 与无校验 Bech32 编解码支持。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_bech32";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bech32';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
