## 用法

来源：

- [Official repository](https://github.com/guedes/ldap_fdw/tree/b277b085981af0dd014850d5e72d5f43cc2f00cf)
- [Extension control file](https://github.com/guedes/ldap_fdw/blob/b277b085981af0dd014850d5e72d5f43cc2f00cf/ldap_fdw.control)
- [PGXN usage documentation](https://pgxn.org/dist/ldap_fdw/0.1.1/)

ldap_fdw 是一个实验性的只读外部数据包装器，把 LDAP 条目暴露为只有两列的 PostgreSQL 行。0.1.1 文档面向 PostgreSQL 9.2，并明确要求不要用于生产。它不支持 TLS、分页和写入，也没有现代安全模型。

### 核心流程

定义 LDAP 服务器和绑定凭据，然后创建一个恰好包含 `dn` 和 `object_body` 两个文本列的外部表：

```sql
CREATE EXTENSION ldap_fdw;

CREATE SERVER directory
  FOREIGN DATA WRAPPER ldap_fdw
  OPTIONS (address 'ldap.internal', port '389');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER directory
  OPTIONS (user_dn 'cn=reader,dc=example,dc=org', password 'secret');

CREATE FOREIGN TABLE directory_entries (
  dn text,
  object_body text
)
SERVER directory
OPTIONS (
  base_dn 'ou=people,dc=example,dc=org',
  attributes 'uid,cn,mail'
);

SELECT dn, object_body
FROM directory_entries
WHERE dn = 'uid=alice,ou=people,dc=example,dc=org';
```

### 重要对象

- `ldap_fdw` 是使用同步 LDAP 客户端的包装器。
- `address` 和 `port` 是服务器选项；默认端口为 389。
- `user_dn` 和 `password` 是用户映射选项。
- `base_dn` 是外部表必需选项，`attributes` 用于限制返回属性。
- `dn` 是条目的可分辨名称，`object_body` 是与 Hstore 兼容的文本。

### 安全与查询限制

只有 `dn` 等值条件会下推。其他扫描会在基础 DN 下执行单层搜索，既无分页也不下推 SQL LIMIT，因此 LDAP 服务器的结果上限可能静默截断结果。实现会在未加密连接上执行简单绑定，并把凭据存入 PostgreSQL 系统目录。不要用于真实凭据或敏感目录属性；应改用支持 TLS 和分页且仍受维护的集成方案。
