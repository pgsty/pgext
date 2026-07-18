## 用法

来源：

- [TencentDB PostgreSQL 权限文档](https://cloud.tencent.com/document/product/409/61551)
- [TencentDB PostgreSQL 扩展支持矩阵](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_superuser` 是 TencentDB for PostgreSQL 的服务商扩展，与托管角色 `pg_tencentdb_superuser` 配套。该角色替代不可用的原生 superuser 权限，提供广泛管理能力，包括创建数据库与角色、绕过行安全、复制、对象访问和受控扩展管理。

```sql
CREATE EXTENSION tencentdb_superuser;
SELECT rolname, rolcreatedb, rolcreaterole, rolreplication, rolbypassrls
FROM pg_roles
WHERE rolname = 'pg_tencentdb_superuser';
```

这是腾讯云服务组件，并非可移植的开源扩展。能力与版本可用性取决于 TencentDB 引擎版本和服务策略。只应把托管角色授予可信管理员，审计成员变更，避免由应用持有，并根据当前实例文档逐项验证所需操作。
