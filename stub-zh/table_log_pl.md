## 用法

来源：

- [官方上游文档](https://github.com/glynastill/table_log_pl/blob/564007f69efb810d0dc2123a8080037810283fff/README.md)

`table_log_pl` — 以 PL/pgSQL 实现的 table_log 兼容替代扩展，用触发器记录表级 DML 变更。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "table_log_pl";
SELECT extversion
FROM pg_extension
WHERE extname = 'table_log_pl';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
