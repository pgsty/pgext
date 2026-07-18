## 用法

来源：

- [官方扩展控制文件](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/pg_audit.control)
- [官方上游文档](https://github.com/disqus/pg_audit/blob/9c4c21c1a4bd94554fdcbc474043eaba68fb44cd/README.md)

`pg_audit` — pg_audit：安全类 PostgreSQL 扩展；上游说明为“用于创建审计表的扩展”。

已复核目录快照记录的版本为 `0.1.3`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_audit";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_audit';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
