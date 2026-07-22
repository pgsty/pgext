## 用法

来源：

- [官方 PGXN README](https://pgxn.org/dist/passwd-fdw/0.4.0/README.html)
- [官方 PGXN 控制文件](https://api.pgxn.org/src/passwd-fdw/passwd-fdw-0.4.0/passwd-fdw.control)
- [官方 PGXN 发行页](https://pgxn.org/dist/passwd-fdw/)

`passwd-fdw` 通过外部表暴露主机的 Unix 用户与用户组数据库。数据通常来自 `/etc/passwd` 与 `/etc/group`，但系统名称服务切换也可能从 LDAP 或其他提供商解析。它可能泄露主机账号与服务信息，因此访问权属于管理员权限边界。

### 核心流程

```sql
CREATE EXTENSION "passwd-fdw";

CREATE SERVER passwd_svr
    FOREIGN DATA WRAPPER passwd_fdw;

CREATE FOREIGN TABLE passwd_users (
    name text,
    passwd text,
    uid integer,
    gid integer,
    gecos text,
    dir text,
    shell text
) SERVER passwd_svr;

CREATE FOREIGN TABLE passwd_groups (
    name text,
    gid integer,
    members text[]
) SERVER passwd_svr;

SELECT name, uid, gid, dir, shell
FROM passwd_users
ORDER BY uid;
```

扩展名包含连字符，因此必须加引号；安装后的包装器名称是 `passwd_fdw`。

### 数据与查询边界

用户表对应密码数据库字段，用户组表则暴露组名、ID 与成员数组。结果表示数据库服务器主机当前的名称服务视图，不是客户端机器信息，也不是事务稳定的 PostgreSQL 目录。

上游文档没有说明写入回调、下推、服务器选项或用户映射。应把它视为只读的本地系统清单，并在打包构建上验证行为，不要增加未记录的选项。

### 安全与兼容性

撤销公众对外部服务器与表的访问，只向受控管理流程授予所需列和行。即使没有密码哈希，账号名、UID、主目录、Shell、组成员，以及 NSS 延迟和故障也属于敏感运维信息。

版本 `0.4.0` 是文档很少的旧实现。应测试目标 PostgreSQL 兼容性、并发 NSS 查询、LDAP 故障、异常字段内容、大型组成员列表及转储恢复。外部表定义可以恢复，但底层主机身份随环境变化，必须在每台服务器上重新验证。
