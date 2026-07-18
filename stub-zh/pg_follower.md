## 用法

来源：

- [官方扩展控制文件](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/pg_follower.control)
- [官方上游文档](https://github.com/HUUTFJ/pg_follower/blob/c533664508bf7de6e3d9dfb0fd601f90350c0cbf/README.md)

`pg_follower` — 用于捕获并应用数据库变更的教学型逻辑复制扩展

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_follower";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_follower';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
