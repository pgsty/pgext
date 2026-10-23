## 用法

来源：

- [官方资料](https://github.com/irfanfaraaz/pg-oauth-rs/blob/6c9e9a79eb1854dcac5bd03536def51b644e2352/README.md)
- [官方资料](https://github.com/irfanfaraaz/pg-oauth-rs/blob/6c9e9a79eb1854dcac5bd03536def51b644e2352/pg_oauth_rs.control)
- [官方资料](https://github.com/irfanfaraaz/pg-oauth-rs/blob/6c9e9a79eb1854dcac5bd03536def51b644e2352/src/entry.rs)

`pg_oauth_rs` 0.2.0 是用于 PostgreSQL 18 身份验证的早期 OAuth 校验模块。它通过服务器的 OAuth 校验器接口验证签名 JWT，不提供 SQL 扩展接口。本地库依赖 libcurl。

### 启用方式

配置模块与实际使用的身份提供方。以下示例 URL 和受众值必须替换为身份提供方给出的真实配置：

```ini
oauth_validator_libraries = 'pg_oauth_rs'
pg_oauth_rs.provider = 'generic'
pg_oauth_rs.audience = 'postgres'
pg_oauth_rs.jwks_uri = 'https://identity.example.com/.well-known/jwks.json'
```

首次启用按上游流程重启服务器。在 `pg_hba.conf` 中配置 OAuth 身份验证、签发者与所需作用域，并通过 `pg_ident.conf` 将已验证身份映射到数据库角色。提供方、受众与身份映射必须一致；仅持有一个签名有效的令牌，并不代表获得了预期的数据库权限。

管理员会话可使用以下诊断操作：

```sql
LOAD 'pg_oauth_rs';
SHOW pg_oauth_rs.provider;
```

### 校验与身份

JWT 签名支持 RS256 和 ES256，校验签发者、受众、过期时间与生效时间。提供方预设涵盖通用签发者、Microsoft Entra ID、Okta 与 Keycloak。通用模式使用主体声明作为身份；Entra 的稳定对象标识避免依赖可能重命名的用户名。使用邮箱身份时需要已验证邮箱声明。不支持不透明访问令牌。

`pg_oauth_rs.identity_claim` 选择映射声明。`pg_oauth_rs.clock_skew_leeway`、`pg_oauth_rs.jwks_timeout`、`pg_oauth_rs.jwks_refetch_interval`、`pg_oauth_rs.jwks_max_age` 和 `pg_oauth_rs.jwks_max_stale` 控制时间容差、密钥拉取与缓存有效期。

### 运行边界

签名密钥在后端进程之间共享缓存。制定紧急密钥撤销方案时，应考虑允许使用过期缓存的时长；身份提供方暂时不可用与密钥被撤销是两种不同情况。密钥端点应使用经过验证的 HTTPS。客户端必须支持 PostgreSQL OAuth 身份验证，不能假定旧客户端或 PostgreSQL 19 路线图中的功能适用于本版本。
