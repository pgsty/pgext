## 用法

来源：

- [官方上游 README](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/docs/README.md)
- [官方扩展控制文件 (veil2.control)](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/veil2.control)
- [官方扩展 SQL (veil2--0.9.1.sql)](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/sql/veil2--0.9.1.sql)

`veil2` — 为虚拟私有数据库实现提供基础。在实现相应的安全、审计或访问控制工作流时使用它。经过审核的上游材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION veil2;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `veil2.add_session_privileges(scope_type_id integer, scope_id integer, roles bitmap, privs bitmap)` 是一个扩展函数，返回 `void`。
- `veil2.all_accessor_roles(accessor_id in out integer, session_context_type_id in integer, session_context_id in integer, role_id out integer, context_type_id out integer, context_id out integer)` 是一个扩展函数，返回 `setof`。
- `veil2.always_true(integer)` 是一个扩展函数，返回 `boolean`。
- `veil2.authenticate(accessor_id integer, authent_type text, token text)` 是一个扩展函数，返回 `boolean`。
- `veil2.authenticate_bcrypt(accessor_id integer, token text)` 是一个扩展函数，返回 `boolean`。
- `veil2.authenticate_false(accessor_id integer, token text)` 是一个扩展函数，返回 `boolean`。
- `veil2.authenticate_plaintext(accessor_id integer, token text)` 是一个扩展函数，返回 `boolean`。
- `veil2.base_accessor_roleprivs(accessor_id in out integer, session_context_type_id in integer, session_context_id in integer, mapping_context_type_id in out integer, mapping_context_id in out integer, assignment_context_type_id out integer, assignment_context_id out integer, role_id out int…)` 是一个扩展函数，返回 `setof`。
- `veil2.bcrypt(passwd text)` 是一个扩展函数，返回 `text`。
- `veil2.become_accessor(accessor_id in integer, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer, session_context_id in integer, session_id out bigint, session_token out text, success out boolean, errmsg out text)` 是一个扩展函数，返回 `record`。
- `veil2.become_accessor(accessor_id in integer, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer, session_context_id in integer, session_id out integer, session_token out text, success out boolean, errmsg out text)` 是一个扩展函数，返回 `record`。
- `veil2.become_user(username in text, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer default null, session_context_id in integer default null, session_id out bigint, session_token out text, success out boolean, errmsg out text)` 是一个扩展函数，返回 `record`。
- `veil2.become_user(username in text, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer default null, session_context_id in integer default null, session_id out integer, session_token out text, success out boolean, errmsg out text)` 是一个扩展函数，返回 `record`。
- `veil2.check_accessor_context(label text, accessor_id integer, context_type_id integer, context_id integer)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 经过审核的控制文件声明默认版本为 `0.9.3`。
- 首先安装确认的扩展依赖项：`pgbitmap`, `pgcrypto`。
- 控制文件要求超级用户进行安装。
- 上游材料包含显式弃用边界。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
