## 用法

来源：

- [官方扩展控制文件](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/pg_bleve.control)
- [官方上游文档](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/README.md)

`pg_bleve` — 将 Bleve 搜索引擎集成到 PostgreSQL 的全文搜索扩展。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `Go`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_bleve";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bleve';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
