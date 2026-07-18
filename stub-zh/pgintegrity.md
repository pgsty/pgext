## 用法

来源：

- [官方扩展控制文件](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/pgintegrity.control)
- [官方上游文档](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/README.md)

`pgintegrity` — 基于触发器的 PostgreSQL 表行完整性水印检查扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`dblink`。

```sql
CREATE EXTENSION "pgintegrity";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgintegrity';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
