## 用法

来源：

- [官方扩展控制文件](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/fhirpath.control)
- [官方上游文档](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/README.md)

`fhirpath` — fhirpath：为 PostgreSQL 提供用于检查 JSONB 的数据类型。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "fhirpath";
SELECT extversion
FROM pg_extension
WHERE extname = 'fhirpath';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
