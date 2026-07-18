## 用法

来源：

- [官方扩展控制文件](https://github.com/katsuragi-y/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/pg_reject_truncate.control)
- [官方上游 README](https://github.com/katsuragi-y/pg_reject_truncate/blob/4ed44aa2bde9ff52e7569dce90c3cc0352062fe5/README.md)

`pg_reject_truncate` — 通过 ProcessUtility 钩子拒绝 PostgreSQL TRUNCATE 操作。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_reject_truncate";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_reject_truncate';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
