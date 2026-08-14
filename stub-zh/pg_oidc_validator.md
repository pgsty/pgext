## 用法

来源：

- [pg_oidc_validator 1.1.0 README](https://github.com/percona/pg_oidc_validator/blob/1.1.0/README.md)
- [pg_oidc_validator 1.1.0 Keycloak 示例](https://github.com/percona/pg_oidc_validator/tree/1.1.0/examples/keycloak)
- [pg_oidc_validator 1.1.0 验证器源码](https://github.com/percona/pg_oidc_validator/blob/1.1.0/src/pg_oidc_validator.cpp)
- [PostgreSQL 18 OAuth 认证文档](https://www.postgresql.org/docs/18/auth-oauth.html)
- [PostgreSQL 18 libpq OAuth 文档](https://www.postgresql.org/docs/18/libpq-oauth.html)

`pg_oidc_validator` 1.1.0 是 PostgreSQL 18 的 OAuth 验证模块，用于根据 OpenID Connect 提供者验证 JWT 访问令牌。它是没有 control 文件或 SQL 扩展的服务器动态库，因此不要运行 `CREATE EXTENSION`。

### 配置服务器

在 `postgresql.conf` 中加载模块，然后重启 PostgreSQL：

```ini
oauth_validator_libraries = 'pg_oidc_validator'
```

在 `pg_hba.conf` 中添加 OAuth 规则；发行者与所需 scope 必须和提供者匹配。除严格的本地测试外，应使用 `hostssl`：

```text
hostssl  all  all  127.0.0.1/32  oauth  issuer=https://id.example.com/realms/postgres scope="openid postgres" validator=pg_oidc_validator
```

修改 HBA 或验证器设置后应重新加载 PostgreSQL；把模块加入 `oauth_validator_libraries` 本身则需要重启。

默认使用 `sub` 声明作为认证身份。如需返回另一个稳定的字符串声明用于角色匹配，可配置：

```ini
pg_oidc_validator.authn_field = 'email'
```

1.1.0 还提供 `pg_oidc_validator.discovery_url_override`。它会改变发现元数据与 JWKS 的获取位置，但不会改变用于验证 JWT `iss` 声明的发行者；适用于 OIDC 提供者具有不同内外部 URL 的环境。这两个验证器设置都可以通过 `SIGHUP` 重新加载。

如果 HBA 规则没有设置 `map=`，选中的声明必须与请求的 PostgreSQL 角色完全一致。提供者身份与数据库角色不同时，应使用具名的 `pg_ident.conf` 映射；验证器不会创建角色。

### 使用 libpq 连接

支持 OAuth 的 libpq 客户端可以启动提供者的设备授权流程：

```bash
psql 'host=127.0.0.1 dbname=app user=alice oauth_issuer=https://id.example.com/realms/postgres oauth_client_id=postgres-client'
```

仅在注册客户端要求时使用 `oauth_client_secret`。客户端标识、请求的 scope、发行者与提供者配置必须一致。

### 提供者与安全边界

- Keycloak 必须为命令行客户端启用 OAuth 2 device flow。
- Microsoft Entra ID 要求租户专属的 v2 发行者与自定义 scope；在 `pg_hba.conf` 中使用完整 scope 名称。
- Google 无法通过 libpq 内置 device flow 使用，但自定义客户端可能可用。
- Dex 不会发送 OAuth scope；显式使用空的 `scope=""` 会关闭 scope 校验，从而削弱常规检查。
- 客户端的 `oauth_issuer` 必须与 HBA 发行者及发现文档完全一致。应把发行者与任何 `pg_oidc_validator.discovery_url_override` 端点都视为可信安全边界，并对数据库和提供者连接强制执行经过验证的 TLS。
- 令牌校验不能替代 PostgreSQL 授权、角色成员关系或行级安全。
- Pigsty RPM 软件包仅覆盖 EL10；DEB 软件包覆盖受支持的 Debian 与 Ubuntu 目标。该模块要求 PostgreSQL 18。
