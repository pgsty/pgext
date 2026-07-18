## 用法

来源：

- [官方扩展控制文件](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/call_graph.control)
- [官方上游文档](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/README)

`call_graph` — 自动跟踪 PostgreSQL 函数调用，并聚合为调用图及可选的表访问统计。

已复核目录快照记录的版本为 `1.4`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "call_graph";
SELECT extversion
FROM pg_extension
WHERE extname = 'call_graph';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
