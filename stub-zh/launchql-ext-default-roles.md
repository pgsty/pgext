## 用法

来源：

- [packages/default-roles/launchql-ext-default-roles.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/default-roles/launchql-ext-default-roles.control)
- [packages/default-roles/sql/launchql-ext-default-roles--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/default-roles/sql/launchql-ext-default-roles--0.4.5.sql)

`launchql-ext-default-roles` 0.4.5 创建 LaunchQL 扩展所需的共享角色。

### 核心工作流

```sql
CREATE EXTENSION "launchql-ext-default-roles" CASCADE;
```

### 角色与权限

SQL 仅在角色尚不存在时创建 `anonymous`、`authenticated` 与 `administrator`。新建管理员角色会继承前两个组。安装前应检查已有角色定义；扩展会复用已有角色，而非统一重设其属性。

控制文件要求 `plpgsql`，并允许非超级用户安装，但创建角色仍需要相应的集群权限。无需共享库或预加载。角色属于整个集群，`DROP EXTENSION` 不会删除它们。所核验组件未声明具体 PostgreSQL 主版本兼容范围。

新建 `administrator` 时还会授予 `BYPASSRLS`，这要求超级用户权限。应谨慎审查该角色的成员关系。
