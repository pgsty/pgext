## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/psl/)

`psl` — 使用编译进扩展的 Public Suffix List 返回主机名对应的可注册域名。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "psl";
SELECT extversion
FROM pg_extension
WHERE extname = 'psl';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
