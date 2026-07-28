## 用法

来源：

- [官方上游 README](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [官方扩展控制文件 (launchql-extension-utils.control)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/utils/launchql-extension-utils.control)
- [官方扩展 SQL (launchql-extension-utils--0.0.1.sql)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/utils/sql/launchql-extension-utils--0.0.1.sql)

`launchql-extension-utils` — PostgreSQL 工具集。使用它来进行相应的 SQL 或数据库工具工作流。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-extension-utils";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `get_entity_from_str(qualified_name text)` 是一个扩展函数，返回 `text`。
- `get_schema_from_str(qualified_name text)` 是一个扩展函数，返回 `text`。
- `list_indexes(_table text, _index text)` 是一个扩展函数，返回 `TABLE`。
- `list_memberships(_user text)` 是一个扩展函数，返回 `TABLE`。
- `tg_update_timestamps()` 是一个扩展函数，返回 `trigger`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
