## 用法

来源：

- [官方 README](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/README.md)
- [扩展 SQL：有效访问权限](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/sql/check_access.sql)
- [扩展 SQL：原始授权](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/sql/check_grants.sql)

`check_access` 0.1 是 Crunchy Data 提供的数据库级权限审计扩展。它展开 PostgreSQL 角色成员关系，分别报告角色实际可用的权限或底层授权。应由数据库超级用户安装；它只检查当前数据库，不会修改权限。

### 核心流程

创建扩展后，先使用对当前调用者开放的视图：

```sql
CREATE EXTENSION check_access;

SELECT *
FROM my_privs
WHERE object_schema = 'app';

SELECT *
FROM my_privs_sys
WHERE object_schema = 'pg_catalog';
```

`my_privs` 排除系统对象，`my_privs_sys` 则包含它们。两者都由安全定义者函数支撑，也是安装脚本唯一保留给 `PUBLIC` 执行的报告接口。如果普通用户需要审计其他角色，管理员必须显式授予更广泛函数的执行权限。

具有相应权限时，可对照检查有效权限与授权记录：

```sql
SELECT * FROM all_access() WHERE base_role = 'app_user';
SELECT * FROM all_access(true) WHERE base_role = 'app_user';

SELECT * FROM check_access('app_user', false);
SELECT * FROM check_grants('app_user', false);
```

`check_access` 和 `all_access` 要求角色拥有容器模式的 `USAGE`，因此反映实际可用权限。`check_grants` 和 `all_grants` 会保留原始对象授权，即使模式权限使其暂时不可用。它们的布尔参数控制是否包含系统对象。

### 报告字段与角色展开

重要字段包括 `base_role`（被审计角色）、`as_role`（提供权限的角色）、`role_path`（成员关系路径）、对象标识、权限名称以及是否带授权选项。扩展也会沿 `NOINHERIT` 链跟踪成员关系；`role_path` 用 `(true)` 或 `(false)` 标记每个祖先关系是否可继承。

该 SQL 覆盖数据库、表空间、外部数据包装器、外部服务器、语言、模式、函数、表、视图、序列及列权限。它是较早期的系统目录实现，早于若干新 PostgreSQL 对象类型；并未声称完整覆盖过程、物化视图、分区关系或未来的目录变化。将结果用作完整安全审计之前，应先在实际 PostgreSQL 版本上验证覆盖范围。
