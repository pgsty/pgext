## 用法

来源：

- [官方扩展控制文件](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/pgrgraphic.control)
- [官方上游文档](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/README.md)

`pgrgraphic` — pgrgraphic 是一个基于 PL/R 的 PostgreSQL 扩展，可将数据库查询结果渲染为 PNG 或 JPEG 图表文件。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `R`。
应先安装并验证声明的扩展依赖：`plr`。

```sql
CREATE EXTENSION "pgrgraphic";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrgraphic';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
