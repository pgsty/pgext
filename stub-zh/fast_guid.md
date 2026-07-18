## 用法

来源：

- [官方扩展控制文件](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/fast_guid.control)
- [官方上游文档](https://github.com/dyninc/postgresql-fast-guid/blob/1fc31f4f9177be1430411a82da1da23834a83ac7/README.rst)

`fast_guid` — fast_guid：为 PostgreSQL 提供快速 GUID 生成函数。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "fast_guid";
SELECT extversion
FROM pg_extension
WHERE extname = 'fast_guid';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
