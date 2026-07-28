## 用法

来源：

- [官方上游 README](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [官方扩展控制文件 (launchql-extension-verify.control)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/verify/launchql-extension-verify.control)
- [官方扩展 SQL (launchql-extension-verify--0.0.1.sql)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/verify/sql/launchql-extension-verify--0.0.1.sql)

`launchql-extension-verify` — PostgreSQL 验证工具集。当应用程序需要此特定数据库功能时使用它。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-extension-verify";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `verify_constraint(_table text, _name text)` 是一个扩展函数，返回 `boolean`。
- `verify_domain(_type text)` 是一个扩展函数，返回 `boolean`。
- `verify_extension(_extname text)` 是一个扩展函数，返回 `boolean`。
- `verify_function(_name text, _user text DEFAULT NULL)` 是一个扩展函数，返回 `boolean`。
- `verify_index(_table text, _index text)` 是一个扩展函数，返回 `boolean`。
- `verify_membership(_user text, _role text)` 是一个扩展函数，返回 `boolean`。
- `verify_policy(_policy text, _table text)` 是一个扩展函数，返回 `boolean`。
- `verify_role(_user text)` 是一个扩展函数，返回 `boolean`。
- `verify_schema(_schema text)` 是一个扩展函数，返回 `boolean`。
- `verify_security(_table text)` 是一个扩展函数，返回 `boolean`。
- `verify_table(_table text)` 是一个扩展函数，返回 `boolean`。
- `verify_table_grant(_table text, _privilege text, _role text)` 是一个扩展函数，返回 `boolean`。
- `verify_trigger(_trigger text)` 是一个扩展函数，返回 `boolean`。
- `verify_type(_type text)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.0.1`。
- 先安装并验证确认的扩展依赖项：`plpgsql`, `uuid-ossp`, `launchql-extension-utils`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
