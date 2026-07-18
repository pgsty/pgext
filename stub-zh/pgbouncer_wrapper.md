## 用法

来源：

- [官方扩展控制文件](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/pgbouncer_wrapper.control)
- [官方上游文档](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/README.md)

`pgbouncer_wrapper` — 通过 dblink_fdw 将 PgBouncer 管理控制台封装为 SQL 监控视图与控制函数

已复核目录快照记录的版本为 `1.2.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`dblink`。

```sql
CREATE EXTENSION "pgbouncer_wrapper";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgbouncer_wrapper';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
