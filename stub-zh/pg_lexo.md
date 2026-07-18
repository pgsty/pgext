## 用法

来源：

- [官方扩展控制文件](https://github.com/Blad3Mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/pg_lexo.control)
- [官方上游文档](https://github.com/Blad3Mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/README.md)
- [官方 Rust 包清单](https://github.com/Blad3Mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/Cargo.toml)

`pg_lexo` — 提供 Base62 字典序位置类型及辅助函数，可在无需批量更新其他行的情况下稳定排序与重排列表。

已复核目录快照记录的版本为 `0.6.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_lexo";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lexo';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
