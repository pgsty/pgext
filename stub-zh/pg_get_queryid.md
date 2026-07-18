## 用法

来源：

- [官方上游文档](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/README.md)

`pg_get_queryid` — 按后端进程号取得 pg_stat_statements 最近生成或使用的查询标识符

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_stat_statements`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_get_queryid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_get_queryid';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
