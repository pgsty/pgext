## 用法

来源：

- [官方上游 README](https://github.com/asghonim/pgho_permissions/blob/077ab341aade6e33fc7c60eee0297cdb37b67ed2/README.md)
- [官方扩展控制文件 (pgho_permissions.control)](https://github.com/asghonim/pgho_permissions/blob/077ab341aade6e33fc7c60eee0297cdb37b67ed2/pgho_permissions.control)
- [官方扩展 SQL (pgho_permissions--0.0.18.sql)](https://github.com/asghonim/pgho_permissions/blob/077ab341aade6e33fc7c60eee0297cdb37b67ed2/pgho_permissions--0.0.18.sql)

`pgho_permissions` — 一个 PostgreSQL 扩展，为您的数据库带来灵活的分层访问控制。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgho_permissions;

-- Install the extension
SELECT dbdev.install('asghonim@pgho_permissions');

-- (Re)create the schema and extension
DROP EXTENSION  IF EXISTS "asghonim@pgho_permissions";
DROP SCHEMA     IF EXISTS pgho_permissions;
CREATE SCHEMA   IF NOT EXISTS pgho_permissions;
CREATE EXTENSION IF NOT EXISTS "asghonim@pgho_permissions"
  SCHEMA pgho_permissions
  VERSION '0.0.18';
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.18`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
