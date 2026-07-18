## 用法

来源：

- [官方扩展控制文件](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/pg_force_unlogged_create_table.control)
- [官方上游文档](https://github.com/gurjeet/pg_force_unlogged_create_table/blob/b143432f1f05bf0a16cb7480dd23acd4bf3e43f8/README.md)

`pg_force_unlogged_create_table` — 通过 ProcessUtility 钩子强制 CREATE TABLE 与 CREATE TABLE AS 创建 UNLOGGED 表。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_force_unlogged_create_table";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_force_unlogged_create_table';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
