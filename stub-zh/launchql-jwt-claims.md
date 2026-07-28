## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jwt-claims/README.md)
- [官方扩展控制文件 (launchql-jwt-claims.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jwt-claims/launchql-jwt-claims.control)

`launchql-jwt-claims` — PostgreSQL 扩展，用于在数据库函数中访问 JWT 声明。此扩展提供模式和函数，使您能够在 PostgreSQL 中直接访问 JWT 令牌声明，从而轻松实现身份验证和授权逻辑。当实现相应的安全、审计或访问控制工作流时，请使用它。在安装此扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-jwt-claims";

-- Access user ID from JWT claims
SELECT ctx.user_id();

-- Access IP address from JWT claims
SELECT ctx.ip_address();

-- Access user agent from JWT claims
SELECT ctx.user_agent();

-- Access origin from JWT claims
SELECT ctx.origin();

-- Access database ID from JWT claims
SELECT jwt_private.current_database_id();

-- Access token ID from JWT claims
SELECT jwt_private.current_token_id();

-- Access IP address (public schema)
SELECT jwt_public.current_ip_address();
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.4.5`。
- 请先安装确认的扩展依赖项：`plpgsql`、`uuid-ossp`、`launchql-ext-types`。
- 控制文件将此扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
