## 用法

来源：

- [官方扩展控制文件](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/postgres_ical.control)
- [官方上游文档](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/README.md)
- [官方 Rust 包清单](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/Cargo.toml)

`postgres_ical` — 在 PostgreSQL 内解析 RFC 5545 iCalendar 文本与远程 URL。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postgres_ical";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_ical';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
