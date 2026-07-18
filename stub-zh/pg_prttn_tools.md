## 用法

来源：

- [官方扩展控制文件](https://github.com/alexandersamoylov/pg_prttn_tools/blob/c7a034b56cd8c2327779dbc0d8060fd2220eb010/pg_prttn_tools.control)

`pg_prttn_tools` — 提供用于管理基于继承机制的分区表的 PL/pgSQL 工具。

已复核目录快照记录的版本为 `0.5`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "pg_prttn_tools";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_prttn_tools';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
