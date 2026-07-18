## 用法

来源：

- [官方扩展控制文件](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/pgopenssltypes.control)
- [官方上游文档](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/README.md)

`pgopenssltypes` — pgopenssltypes 为 PostgreSQL 提供可持久化的 OpenSSL 密码学类型及相关功能。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgopenssltypes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgopenssltypes';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
