## 用法

来源：

- [官方扩展控制文件](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/pg_quota.control)

`pg_quota` — 通过后台工作进程实施按角色统计的软磁盘空间配额

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_quota";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_quota';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
