## 用法

来源：

- [Official database.dev 包页面](https://database.dev/pmnzt/custom_roles_patch)

`pmnzt@custom_roles_patch` — custom_roles。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "pmnzt@custom_roles_patch";
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `custom_roles_update_to_app_metadata()` 是一个扩展函数，返回 `trigger`。
- `get_user_roles()` 是一个扩展函数，返回 `text[]`。
- `user_has_role(_role text)` 是一个扩展函数，返回 `boolean`。
- `user_role_in(_roles text[])` 是一个扩展函数，返回 `boolean`。
- `user_roles_match(_roles text[])` 是一个扩展函数，返回 `boolean`。
- `custom_role_names` 是一个由扩展安装或管理的表。
- `custom_user_roles` 是一个由扩展安装或管理的表。
- `on_custom_role_change` 是一个扩展定义的触发器。

### 要求与注意事项

- 该目录记录版本 `0.0.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
