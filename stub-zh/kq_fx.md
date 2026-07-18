## 用法

来源：

- [官方扩展控制文件](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/kq_fx.control)
- [官方上游文档](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/README.md)
- [官方 Rust 包清单](https://github.com/ketteq-neon/kq_fx/blob/d62ad40bf4dc61d568801f6ed768c77128256e4b/Cargo.toml)

`kq_fx` — kq_fx：通用类 PostgreSQL 扩展；上游说明为“ketteQ, Inc”。

已复核目录快照记录的版本为 `1.0.1`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "kq_fx";
SELECT extversion
FROM pg_extension
WHERE extname = 'kq_fx';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
