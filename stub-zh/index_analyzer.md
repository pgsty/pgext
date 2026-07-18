## 用法

来源：

- [官方扩展控制文件](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/index_analyzer.control)
- [官方上游文档](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/README.md)

`index_analyzer` — 用于评估现有或候选索引并检查外键索引覆盖情况的 PL/pgSQL 函数。

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "index_analyzer";
SELECT extversion
FROM pg_extension
WHERE extname = 'index_analyzer';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
