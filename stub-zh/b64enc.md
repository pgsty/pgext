## 用法

来源：

- [官方扩展控制文件](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/b64enc.control)
- [官方 Rust 包清单](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/Cargo.toml)
- [官方上游 README](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/README.md)

`b64enc` — Rust/C 扩展，定义 b64enc 基础类型，并提供 URL-safe Base64 输入和输出函数。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。

```sql
CREATE EXTENSION "b64enc";
SELECT extversion
FROM pg_extension
WHERE extname = 'b64enc';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
