## 用法

来源：

- [官方扩展控制文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/eidos_oauth.control)
- [官方上游文档](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/README.md)
- [官方 Rust 包清单](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/Cargo.toml)

`eidos_oauth` — 面向 PostgreSQL 18 OAuth 验证器 API 的 OAuth 2.0 JWT 校验扩展。

已复核目录快照记录的版本为 `0.2.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "eidos_oauth";
SELECT extversion
FROM pg_extension
WHERE extname = 'eidos_oauth';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
