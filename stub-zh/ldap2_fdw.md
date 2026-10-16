## 用法

来源：

- [README.md](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/README.md)
- [ldap2_fdw--1.0.sql](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/ldap2_fdw--1.0.sql)
- [ldap2_fdw.c](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/ldap2_fdw.c)
- [ldap2_fdw.control](https://github.com/zonkiie/pg_ldap_fdw2/blob/8a28137a7ea30e6bc637d6fd73b590d7162c7836/ldap2_fdw.control)

`ldap2_fdw` 1.0 将 LDAP 条目映射为外部表并支持写入。作者明确警告不要用于生产，测试、内存管理和错误处理仍不完善。

### 核心用法

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION ldap2_fdw;
CREATE SERVER ldap_directory FOREIGN DATA WRAPPER ldap2_fdw
    OPTIONS (uri 'ldap://localhost');
```

### 运行边界

官方安装流程先创建 `uuid-ossp`，导入生成的定义可能使用 `uuid_generate_v4()`。需按实际目录设置服务器 `uri`、用户映射的 `username` 和 `password`，以及外部表的 `basedn`、`filter`、`scope`。凭据保存在用户映射中，并需明确控制授权。LDAP 条目不一定能完整映射为关系结构，服务器大小限制可能截断结果，远程过滤仅支持没有其他条件的单个 DN 谓词。默认仅超级用户可安装。上游未声明支持的 PostgreSQL 大版本范围，也未提供独立许可证。
