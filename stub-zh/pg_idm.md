## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/README.md)
- [1.0 版本 SQL 对象](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm--1.0.sql)
- [扩展 control 文件](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm.control)
- [GPL-3.0 许可证](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/LICENSE)

`pg_idm` 将 PostgreSQL 角色与成员关系呈现为可写关系视图。应安装到调用者指定的专用 schema，以避免对象名冲突。

```sql
CREATE SCHEMA idm;
CREATE EXTENSION pg_idm WITH SCHEMA idm;
SELECT rolname, rolcanlogin, rolconnlimit
FROM idm.roles
ORDER BY rolname;
```

对 `idm.roles` 执行插入、更新和删除，会分别映射为角色创建、修改和移除。对 `idm.auth_members` 执行插入和删除，则映射为成员关系授予与撤销。这些操作修改的是集群级安全主体，而不仅是当前数据库中的记录。

触发器函数使用调用者权限执行，因此调用者仍需具备对应的 PostgreSQL 角色管理权限。只应把这些视图的写权限授予严格受控的管理角色，并像直接执行 `CREATE ROLE`、`ALTER ROLE`、`GRANT` 和 `DROP ROLE` 一样审计变更。
