## 用法

来源：

- [官方扩展控制文件](https://github.com/quentusrex/pg_sip/blob/3fb695cc3e10748e683db2ad8b60b8037fbb6c45/pg_sip.control)

`pg_sip` — 早期 SIP 扩展，目前仅提供底层库版本查询函数

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_sip";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sip';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
