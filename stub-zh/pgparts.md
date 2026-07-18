## 用法

来源：

- [官方扩展控制文件](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/pgparts.control)
- [官方上游文档](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/README.md)

`pgparts` — 用于通过触发器创建和维护表分区的 SQL/PLpgSQL 辅助扩展。

已复核目录快照记录的版本为 `0.0.15`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "pgparts";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgparts';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
