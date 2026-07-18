## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/pg_credereum/blob/master/pg_credereum.control)
- [官方上游文档](https://github.com/postgrespro/pg_credereum/blob/master/README.md)

`pg_credereum` — 通过区块链式事务块提供可进行密码学验证的数据库审计能力。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_credereum";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_credereum';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
