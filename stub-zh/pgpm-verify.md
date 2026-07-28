## 用法

来源：

- [官方上游 README](https://github.com/babdulhakim2/pgpm-test/blob/8ce14c436c1ace0eb846844ed78b22b777931036/extensions/@pgpm/verify/README.md)
- [官方扩展控制文件 (pgpm-verify.control)](https://github.com/babdulhakim2/pgpm-test/blob/8ce14c436c1ace0eb846844ed78b22b777931036/extensions/@pgpm/verify/pgpm-verify.control)
- [官方扩展 SQL (pgpm-verify--0.15.3.sql)](https://github.com/babdulhakim2/pgpm-test/blob/8ce14c436c1ace0eb846844ed78b22b777931036/extensions/@pgpm/verify/sql/pgpm-verify--0.15.3.sql)

`pgpm-verify` — PostgreSQL 模块的验证工具。当应用程序需要此特定数据库功能时使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-verify";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `get_entity_from_str(qualified_name text)` 是一个扩展函数，返回 `text`。
- `get_schema_from_str(qualified_name text)` 是一个扩展函数，返回 `text`。
- `list_indexes(_table text, _index text)` 是一个扩展函数，返回 `TABLE`。
- `list_memberships(_user text)` 是一个扩展函数，返回 `TABLE`。
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

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.15.3`。
- 先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
