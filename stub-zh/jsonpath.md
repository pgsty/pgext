## 用法

来源：

- [官方扩展控制文件](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/jsonpath.control)
- [官方上游 README](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/README.md)

`jsonpath` — 通过 PLV8 中内嵌的 JavaScript 实现执行 JSONPath 查询。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plv8`。

```sql
CREATE EXTENSION "jsonpath";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonpath';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
