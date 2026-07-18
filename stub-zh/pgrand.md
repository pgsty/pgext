## 用法

来源：

- [官方扩展控制文件](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/pgrand.control)
- [官方 Rust 包清单](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/Cargo.toml)

`pgrand` — 通过 PostgreSQL FDW 生成随机测试数据，并包含实验性的表访问方法。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgrand";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrand';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
