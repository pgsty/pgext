## 用法

来源：

- [官方扩展控制文件](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/audit.control)
- [官方上游文档](https://github.com/ttfkam/pg_audit/blob/c7f3069058215873ae8e497e1476c1e195309479/README.md)

`audit` — 为 PostgreSQL 表的 INSERT、UPDATE、DELETE 操作记录审计数据。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "audit";
SELECT extversion
FROM pg_extension
WHERE extname = 'audit';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
