## 用法

来源：

- [官方扩展控制文件](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/pg_dbwa.control)
- [官方上游文档](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/README.md)

`pg_dbwa` — 通过保存 pg_stat_statements 历史快照分析本地或远程 PostgreSQL 工作负载。

已复核目录快照记录的版本为 `0.3.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`dblink`, `pg_eyes`, `pg_prttn_tools`, `pg_stat_statements`, `plpgsql`。

```sql
CREATE EXTENSION "pg_dbwa";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dbwa';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
