## 用法

来源：

- [官方扩展控制文件](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/fhirbase.control)
- [官方上游文档](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/README.md)

`fhirbase` — 面向 FHIR 数据库的未完成 PostgreSQL 扩展脚手架。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "fhirbase";
SELECT extversion
FROM pg_extension
WHERE extname = 'fhirbase';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
