## 用法

来源：

- [官方扩展控制文件](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/json_fdw.control)
- [官方上游文档](https://github.com/nkhorman/json_fdw/blob/65b8e7d4dc39bb2844879bb3c0199588e9cfa8a2/README.md)

`json_fdw` — 用于访问本地及远程 JSON 文档的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "json_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'json_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
