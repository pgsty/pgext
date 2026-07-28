## 用法

来源：

- [Official database.dev 包页面](https://database.dev/pointsource/supabase_rbac)

`pointsource@supabase_rbac` — 基于角色的访问控制，适用于您的 Supabase 项目。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "pointsource@supabase_rbac";
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `_build_user_claims` 是一个扩展函数。
- `_check_permission_escalation` 是一个扩展函数。
- `_check_role_escalation` 是一个扩展函数。
- `_get_user_groups` 是一个扩展函数。
- `_jwt_is_expired` 是一个扩展函数。
- `_on_group_created` 是一个扩展函数。
- `_on_role_definition_change` 是一个扩展函数。
- `_set_updated_at` 是一个扩展函数。
- `_sync_member_metadata` 是一个扩展函数。
- `_sync_member_permission` 是一个扩展函数。
- `_validate_grantable_roles` 是一个扩展函数。
- `_validate_permissions` 是一个扩展函数。
- `_validate_roles` 是一个扩展函数。
- `accept_invite` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `5.2.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
