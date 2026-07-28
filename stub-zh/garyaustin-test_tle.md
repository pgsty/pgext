## 用法

来源：

- [Official database.dev 包页面](https://database.dev/garyaustin/test_tle)

`garyaustin-test_tle` — 测试自定义属性。当应用程序需要此特定数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中对其进行测试。

### 核心工作流

```sql
CREATE EXTENSION "garyaustin-test_tle";
```

在目标数据库中安装扩展，如果有可用示例，则运行上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `update_to_app_metadata` 是一个扩展函数。
- `user_has_property` 是一个扩展函数。
- `user_properties_match` 是一个扩展函数。
- `user_property_in` 是一个扩展函数。
- `property_names` 是由扩展安装或管理的表。
- `user_properties` 是由扩展安装或管理的表。
- `IF` 是由扩展创建的模式。
- `on_role_change` 是一个扩展定义的触发器。

### 要求与注意事项

- 该目录记录版本 `0.0.2`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
