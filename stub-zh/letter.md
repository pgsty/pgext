## 用法

来源：

- [官方上游 README](https://github.com/paulharter/letter/blob/0ca81e61bb444dc30b46e0ee6b3c415f6e696a39/README.md)
- [官方扩展控制文件 (letter.control)](https://github.com/paulharter/letter/blob/0ca81e61bb444dc30b46e0ee6b3c415f6e696a39/letter.control)
- [官方扩展 SQL (letter--0.1.sql)](https://github.com/paulharter/letter/blob/0ca81e61bb444dc30b46e0ee6b3c415f6e696a39/sql/letter--0.1.sql)

`letter` — letter: PostgreSQL 基于角色的访问控制。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION letter;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `letter.assign(source_table text, user_column text, scope_table text DEFAULT NULL, role_name text DEFAULT NULL, role_column text DEFAULT NULL, if_fn text DEFAULT NULL)` 是一个扩展函数，返回 `boolean`。
- `letter.cache_inval()` 是一个扩展函数，返回 `trigger`。
- `letter.enforce_delete()` 是一个扩展函数，返回 `trigger`。
- `letter.enforce_insert()` 是一个扩展函数，返回 `trigger`。
- `letter.enforce_update()` 是一个扩展函数，返回 `trigger`。
- `letter.grant(privilege text, on_table text, role text, columns text[], scope text, using_path text[] DEFAULT NULL, check_fn text DEFAULT NULL)` 是一个扩展函数，返回 `boolean`。
- `letter.list_grants(filter_role text DEFAULT NULL)` 是一个扩展函数，返回 `TABLE`。
- `letter.read(table_name text, condition text DEFAULT NULL)` 是一个扩展函数，返回 `SETOF`。
- `letter.revoke(privilege text, on_table text, role text, columns text[], scope text)` 是一个扩展函数，返回 `boolean`。
- `letter.role_cleanup()` 是一个扩展函数，返回 `trigger`。
- `letter.unassign(source_table text, user_column text, scope_table text DEFAULT NULL, role_name text DEFAULT NULL, role_column text DEFAULT NULL)` 是一个扩展函数，返回 `boolean`。
- `letter.user_permissions(p_user_id text)` 是一个扩展函数，返回 `TABLE`。
- `letter.assignments` 是一个由扩展安装或管理的表。
- `letter.grants` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
