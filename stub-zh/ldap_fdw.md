## 用法

来源：

- [官方扩展控制文件](https://github.com/guedes/ldap_fdw/blob/b277b085981af0dd014850d5e72d5f43cc2f00cf/ldap_fdw.control)
- [官方上游文档](https://pgxn.org/dist/ldap_fdw/0.1.1/doc/ldap_fdw.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/ldap_fdw/)

`ldap_fdw` — 用于查询 LDAP 服务器的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `0.1.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "ldap_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'ldap_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
