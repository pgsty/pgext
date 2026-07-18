## 用法

来源：

- [官方扩展控制文件](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/douban_fdw.control)
- [官方上游文档](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/README.md)

`douban_fdw` — 用于查询豆瓣公开电影排行榜 API 的 PostgreSQL 外部数据包装器

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `Go`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "douban_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'douban_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
