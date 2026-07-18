## 用法

来源：

- [官方扩展控制文件](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/pgmer1.control)
- [官方上游文档](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/README.md)
- [官方 Rust 包清单](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/Cargo.toml)

`pgmer1` — MeritRank HTTP 连接器，在 SQL 中调用图排名服务。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgmer1";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmer1';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
