## 用法

来源：

- [官方扩展控制文件](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/pgotel.control)
- [官方上游文档](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/README.md)

`pgotel` — pgotel 为 PostgreSQL 扩展提供 OpenTelemetry 指标导出能力，并以 C 接口封装 OpenTelemetry C++ SDK。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "pgotel";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgotel';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
