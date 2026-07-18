## 用法

来源：

- [官方扩展控制文件](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/plid.control)
- [官方上游文档](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/README.md)
- [官方 Rust 包清单](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/Cargo.toml)

`plid` — plid 提供 128 位、可按字典序排序的 PostgreSQL PLID 类型，支持短前缀、毫秒时间戳、随机位和可选的单调生成。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "plid";
SELECT extversion
FROM pg_extension
WHERE extname = 'plid';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
