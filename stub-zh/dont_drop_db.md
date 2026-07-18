## 用法

来源：

- [官方扩展控制文件](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/dont_drop_db.control)
- [官方上游文档](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/README.md)

`dont_drop_db` — 阻止删除 dont_drop_db.list 参数中列出的数据库。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "dont_drop_db";
SELECT extversion
FROM pg_extension
WHERE extname = 'dont_drop_db';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
