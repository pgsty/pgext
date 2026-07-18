## 用法

来源：

- [官方扩展控制文件](https://github.com/nyurik/postile/blob/a11e6479aa435fc17379a7c23ad1f9501c43aca8/postile.control)

`postile` — postile 是一个基于 pgrx 的 PostgreSQL 扩展，提供地图瓦片生成辅助函数以及 gzip/Brotli bytea 压缩。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postile";
SELECT extversion
FROM pg_extension
WHERE extname = 'postile';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
