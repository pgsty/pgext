## 用法

来源：

- [TencentDB for PostgreSQL 角色与权限文档](https://cloud.tencent.com/document/product/409/61551)
- [TencentDB for PostgreSQL 扩展支持矩阵](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_superuser` 是 TencentDB for PostgreSQL 的服务商扩展，与托管角色 `pg_tencentdb_superuser` 关联。该角色在腾讯云上提供对原生超级用户能力的受控替代；无论扩展还是权限模型都不能移植到自管 PostgreSQL。

### 核心流程

只有当当前内核支持矩阵暴露该扩展时，才通过 TencentDB 扩展管理界面或 `CREATE EXTENSION` 使用它。然后在把成员资格授予专用管理员前，核对版本与托管角色：

```sql
CREATE EXTENSION tencentdb_superuser;

SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_superuser';

SELECT rolname, rolcreatedb, rolcreaterole, rolreplication, rolbypassrls
FROM pg_roles
WHERE rolname = 'pg_tencentdb_superuser';

GRANT pg_tencentdb_superuser TO dba_role;
```

授权必须由 TencentDB 允许管理该托管角色的账号执行。应审计成员资格，并在确切内核版本上测试所需操作，不要假定它等同于原生超级用户。

### 已确认的权限边界

- 托管角色拥有 `CREATEDB`、`CREATEROLE`、`REPLICATION` 和 `BYPASSRLS` 风格能力，并可广泛访问由非超级用户角色拥有的对象。
- 它可创建受支持的发布、订阅、复制槽和扩展。在创建受支持扩展时，TencentDB 会临时提升当前托管角色，使其通过所需超级用户检查。
- 原生 `superuser`、`pg_execute_server_program`、`pg_read_server_files` 和 `pg_write_server_files` 不向客户开放。
- 加载库受限于服务插件目录，且只有另一个托管超级用户才能终止托管超级用户后端。检查点和事件触发器能力也取决于 TencentDB 内核修订版本。

### 服务与版本边界

官方矩阵在不同 PostgreSQL 内核大版本上列出 `tencentdb_superuser` 版本 1.0 与 1.1；可用性并不统一。因此目录版本 `1.1` 必须被视为特定内核的服务版本，而不是可通用下载的包。只应把该角色授予可信人工或自动化管理员，绝不授予应用登录；应将所有权与日常访问分离，记录成员资格变更和特权 DDL，并在内核升级或服务策略变更后重新检查 TencentDB 文档。
