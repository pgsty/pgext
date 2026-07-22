## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/README.md)
- [版本 1.0 的 SQL 对象](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm--1.0.sql)
- [扩展 control 文件](https://github.com/pasha132/pg_idm/blob/9968b613451322e0c4cc9cb2ecccf0693883d7f3/pg_idm.control)

`pg_idm` 把 PostgreSQL 角色和角色成员关系公开为可写关系视图。`INSERT`、`UPDATE` 和 `DELETE` 操作由 `INSTEAD OF` 触发器转换为角色 DDL，可用于经过严格控制的身份管理集成。这些行代表集群级安全主体，而不是只属于当前数据库的数据。

扩展创建后不可重定位，且使用通用对象名，因此应安装到调用方选择的专用 schema。

```sql
CREATE SCHEMA idm;
CREATE EXTENSION pg_idm WITH SCHEMA idm;

SELECT rolname, rolcanlogin, rolconnlimit
FROM idm.roles
ORDER BY rolname;
```

### 管理角色

```sql
INSERT INTO idm.roles (rolname, rolcanlogin, rolconnlimit)
VALUES ('app_reader', true, 10)
RETURNING rolname, rolcanlogin, rolconnlimit;

UPDATE idm.roles
SET rolcanlogin = false
WHERE rolname = 'app_reader';

DELETE FROM idm.roles
WHERE rolname = 'app_reader';
```

`idm.roles` 投影 `pg_roles` 中的属性，包括 `rolsuper`、`rolinherit`、`rolcreaterole`、`rolcreatedb`、`rolcanlogin`、`rolreplication`、`rolconnlimit`、`rolpassword`、`rolvaliduntil` 和 `rolbypassrls`。插入时省略的属性采用扩展的角色默认值；更新只生成发生变化的角色选项。

### 管理成员关系

```sql
INSERT INTO idm.auth_members (roleid, member, admin_option)
VALUES ('reporting'::regrole, 'app_reader'::regrole, false);

DELETE FROM idm.auth_members
WHERE roleid = 'reporting'::regrole
  AND member = 'app_reader'::regrole;
```

`idm.auth_members` 映射 `pg_auth_members`；插入会执行 `GRANT`，删除会执行 `REVOKE`。该视图不提供更新成员关系的工作流。

### 安全与运维注意事项

- 触发器函数明确声明为 `SECURITY INVOKER`。调用方既要拥有视图 DML 权限，也要拥有生成的角色命令所要求的 PostgreSQL 权限；扩展不会提升调用方权限。
- 角色变更会影响集群中的所有数据库。只应向专用管理角色授予写权限，像直接角色 DDL 一样审计这些操作，并在自动删除前测试所有权与依赖失败。
- 通过 DML 提交的密码值可能进入客户端历史、SQL 日志、审计记录或监控系统。应采用批准的密钥处理与参数化接口；未经这些路径复核，不要把视图当作密码分发通道。
- control 文件说明安装扩展本身不要求超级用户，但 `SUPERUSER`、`REPLICATION` 和 `BYPASSRLS` 等实际角色属性仍受 PostgreSQL 核心权限规则约束。
- 上游版本 `1.0` 声称支持 PostgreSQL 9.6 及以上，但没有维护中的兼容矩阵。应针对实际服务器大版本验证生成的 DDL 和目录列。
