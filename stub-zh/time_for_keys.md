## 用法

来源：

- [官方上游文档](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/README.md)

`time_for_keys` — 以 C 与 PL/pgSQL 约束触发器在有效时间表之间实现时态外键。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "time_for_keys";
SELECT extversion
FROM pg_extension
WHERE extname = 'time_for_keys';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
