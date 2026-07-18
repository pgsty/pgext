## 用法

来源：

- [官方扩展控制文件](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/postgreth.control)
- [官方上游文档](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/README.md)
- [官方 Rust 包清单](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/Cargo.toml)

`postgreth` — postgreth 是一个 pgrx 扩展，提供以太坊 Bloom、地址和哈希类型，并可将 Solidity ABI 项与事件日志解码为 JSONB。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postgreth";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgreth';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
