## 用法

来源：

- [AWS AD 安全组访问控制指南](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AD.Security.Groups.html)
- [AWS Aurora PostgreSQL 扩展版本表](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`pg_ad_mapping` 1.0 是 Amazon Aurora PostgreSQL 的托管扩展，用于将 AWS Directory Service for Microsoft Active Directory 安全组映射到 PostgreSQL 登录角色。它不是可移植的社区软件包：认证、二进制文件、升级和运维控制都由 Aurora 提供。

### 核心流程

首先为 Aurora 集群配置 Kerberos 认证与域成员关系。在自定义 DB cluster parameter group 的 `shared_preload_libraries` 中加入 `pg_ad_mapping`，重启 writer，然后以 `rds_superuser` 完成设置：

```sql
SHOW shared_preload_libraries;
CREATE EXTENSION pg_ad_mapping;

ALTER ROLE "accounts-role" LOGIN;
GRANT CONNECT ON DATABASE "accounts-db" TO "accounts-role";

SELECT pgadmap_set_mapping(
  'accounts-group',
  'accounts-role',
  'S-1-5-21-3168537779-1985441202-1799118680-1612',
  10
);

SELECT * FROM pgadmap_read_mapping();
```

SID 必须标识对应的 AD 安全组。用户同时属于多个已映射组时，权重最高的映射胜出；权重相同则最后添加的映射胜出。

### 重要对象

- `pgadmap_set_mapping(ad_group, db_role, ad_group_sid, weight)` 创建映射；它只能由 primary 实例上的 `rds_superuser` 执行。
- `pgadmap_read_mapping()` 列出所有映射的 SID、数据库角色、权重和 AD 组。
- `pgadmap_reset_mapping(ad_group_sid, db_role, weight)` 删除一条精确匹配；无参数调用 `pgadmap_reset_mapping()` 会删除全部映射。重置操作只能由 primary 实例上的 `rds_superuser` 执行。

### 认证与角色边界

映射的数据库角色必须具有 `LOGIN`，并拥有目标数据库的 `CONNECT` 权限。不要向组映射角色授予 `rds_ad`；AWS 将 `rds_ad` 用于单独配置的 AD 用户。组认证通过带域后缀的 principal 触发，而已经用 `rds_ad` 单独配置的用户不会再通过组映射登录。

### 运维说明

AWS 文档说明 AD 组访问从 Aurora PostgreSQL 14.10 和 15.5 开始提供；具体集群可用的引擎版本与扩展版本应以实时扩展矩阵为准。修改预加载项后必须重启 writer。映射按集群管理，管理性写操作必须在 primary 上执行。

该功能不支持 Microsoft Entra ID。AWS 还记录了已启用 Kerberos 的旧集群迁移风险：加入预加载库后，登录可能失败，甚至可能导致引擎崩溃。对于受影响的既有实例，应按 AWS 指引禁用后重新启用 Kerberos；2025 年 4 月之后创建的实例不需要这一变通步骤。
