## 用法

来源：

- [官方扩展控制文件](https://github.com/umitanuki/twitter_fdw/blob/9433a19b1ee7abb1ac08c4471ece2238785d677b/twitter_fdw.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/twitter_fdw/)

`twitter_fdw` — 面向已退役的免认证 Twitter Search API 的旧版外部数据包装器。

已复核目录快照记录的版本为 `1.1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "twitter_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'twitter_fdw';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
