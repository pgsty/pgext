## 用法

来源：

- [官方扩展控制文件](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/postgreChain.control)
- [官方上游文档](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/README.md)
- [官方 Rust 包清单](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/Cargo.toml)

`postgreChain` — 直接从 PostgreSQL 查询 Solana RPC 数据，并执行钱包、空投与转账操作

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postgreChain";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgreChain';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
