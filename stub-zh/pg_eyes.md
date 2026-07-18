## 用法

来源：

- [官方扩展控制文件](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/pg_eyes.control)
- [官方上游文档](https://github.com/alexandersamoylov/pg_eyes/blob/31c21bb1a307235bfda250727254491d505e0ee3/README.md)

`pg_eyes` — 用于监控 PostgreSQL 数据库活动的 SQL 函数与视图

已复核目录快照记录的版本为 `1.4`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pg_stat_statements`, `plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_eyes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_eyes';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
