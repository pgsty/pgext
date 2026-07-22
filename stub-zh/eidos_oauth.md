## 用法

来源：

- [官方 eidos_oauth README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/README.md)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/eidos_oauth.control)
- [验证器实现](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/src/validator.rs)

`eidos_oauth` 版本 `0.2.0` 为 PostgreSQL 18 连接认证验证 OAuth bearer JWT。它通过远程 JWKS 端点验证签名，检查配置的 issuer 与 audience 声明，返回认证身份，并把声明暴露为会话 GUC，以供授权策略使用。

### 核心流程

在服务器启动时配置库，然后创建 SQL 扩展并检查验证器状态：

```conf
shared_preload_libraries = 'eidos_oauth'
oauth_validator_libraries = 'eidos_oauth'

eidos_oauth.jwks_url = 'https://auth.example.com/.well-known/jwks.json'
eidos_oauth.issuer = 'https://auth.example.com'
eidos_oauth.audience = 'my-app'
eidos_oauth.role_claim = 'sub'
eidos_oauth.role_prefix = 'oauth_'
eidos_oauth.jwks_cache_seconds = 300
eidos_oauth.jwks_timeout_ms = 5000
```

```sql
CREATE EXTENSION eidos_oauth;

SELECT eidos_oauth.oauth_jwks_status();
SELECT eidos_oauth.oauth_refresh_jwks();
SELECT eidos_oauth.oauth_validate_token('signed.jwt.value');
SELECT current_setting('app.user_sub', true);
```

SQL 辅助函数包括 `oauth_validate_token`、`oauth_get_claim`、`oauth_inject_claims`、`oauth_jwks_status` 和 `oauth_refresh_jwks`。连接时验证使用配置的声明（默认为 `sub`）作为认证身份；该身份到数据库角色的映射仍由 PostgreSQL 角色映射负责。

### 运维说明

连接验证器 API 需要 PostgreSQL 18。源码将 `oauth_inject_claims` 描述为测试或 18 之前版本的辅助函数，而不是 PostgreSQL 18 OAuth 认证的替代方案。control 文件不可重定位且仅限超级用户安装。配置变更使用 SIGHUP 上下文 GUC，但加入这些库需要重启。应保护 JWKS 端点，测试密钥轮换与服务中断行为，限制刷新和令牌测试函数，并让 RLS 策略把注入声明视为认证上下文，而不是可信的应用输入。
