## 用法

来源：

- [官方扩展控制文件](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/iban.control)
- [官方上游文档](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/README.md)

`iban` — 为 PostgreSQL 提供 IBAN 数据类型和国际银行账号校验函数。

已复核目录快照记录的版本为 `1.1`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `14`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "iban";
SELECT extversion
FROM pg_extension
WHERE extname = 'iban';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
