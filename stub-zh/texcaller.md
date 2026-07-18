## 用法

来源：

- [官方扩展控制文件](https://github.com/vog/texcaller/blob/fc5fbc9862a6df3e907e2ee409e5d995aa175ef6/postgresql/texcaller.control)
- [官方上游文档](https://vog.github.io/texcaller/group__postgresql.html)

`texcaller` — 在 PostgreSQL 中将 TeX 或 LaTeX 源文档转换为 DVI 或 PDF 的函数。

已复核目录快照记录的版本为 `0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "texcaller";
SELECT extversion
FROM pg_extension
WHERE extname = 'texcaller';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
