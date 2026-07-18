## 用法

来源：

- [官方扩展控制文件](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/explanation.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/explanation/)
- [官方上游 README](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/README.md)

`explanation` — explanation：将 EXPLAIN 计划解析为按邻近树组织的关系型行，便于查询和分析。

已复核目录快照记录的版本为 `0.3.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "explanation";
SELECT extversion
FROM pg_extension
WHERE extname = 'explanation';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
