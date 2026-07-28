## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/agentic-db/blob/cd818fea7f480ef3ff6099f736d66848a92b907a/extensions/@pgpm/jwt-claims/README.md)
- [官方扩展控制文件 (pgpm-jwt-claims.control)](https://github.com/constructive-io/agentic-db/blob/cd818fea7f480ef3ff6099f736d66848a92b907a/extensions/@pgpm/jwt-claims/pgpm-jwt-claims.control)
- [官方扩展 SQL (pgpm-jwt-claims--0.15.5.sql)](https://github.com/constructive-io/agentic-db/blob/cd818fea7f480ef3ff6099f736d66848a92b907a/extensions/@pgpm/jwt-claims/sql/pgpm-jwt-claims--0.15.5.sql)

`pgpm-jwt-claims` — @pgpm/jwt-claims 提供了用于从 PostgreSQL 会话变量中提取和处理 JWT（JSON Web Token）声明的 PostgreSQL 函数。在实现相应的安全、审计或访问控制工作流时使用它。在集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-jwt-claims";
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `ctx.ip_address()` 是一个扩展函数，返回 `inet`。
- `ctx.is_security_definer()` 是一个扩展函数。
- `ctx.origin()` 是一个扩展函数，返回 `origin`。
- `ctx.security_definer()` 是一个扩展函数，返回 `text`。
- `ctx.uagent()` 是一个扩展函数，返回 `text`。
- `ctx.uid()` 是一个扩展函数，返回 `uuid`。
- `jwt_private.current_database_id()` 是一个扩展函数，返回 `uuid`。
- `jwt_private.current_session_id()` 是一个扩展函数，返回 `uuid`。
- `jwt_private.current_token_id()` 是一个扩展函数，返回 `uuid`。
- `jwt_public.current_ip_address()` 是一个扩展函数，返回 `inet`。
- `jwt_public.current_origin()` 是一个扩展函数，返回 `origin`。
- `jwt_public.current_user_agent()` 是一个扩展函数，返回 `text`。
- `jwt_public.current_user_id()` 是一个扩展函数，返回 `uuid`。
- `ctx` 是由扩展创建的一个模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.15.5`。
- 先安装并验证确认的扩展依赖项：`plpgsql`, `pgpm-types`, `pgpm-verify`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
