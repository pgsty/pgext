## 用法

来源：

- [pg_oidc_validator 0.2 README](https://github.com/percona/pg_oidc_validator/blob/0.2/README.md)
- [Keycloak example for 0.2](https://github.com/percona/pg_oidc_validator/tree/0.2/examples/keycloak)

pg_oidc_validator 是一个用于 PostgreSQL 18 的 OAuth 验证模块，可以验证 libpq OAuth 令牌与 OpenID Connect 发行商。当 PostgreSQL 客户端通过 OIDC 提供者进行身份验证时，请使用此模块；它由服务器加载，并不定义 SQL 扩展，因此无需运行 CREATE EXTENSION。

该项目将该模块描述为实验性且尚未准备好生产环境使用。在依赖之前，请先测试具体的认证提供者、客户端和 PostgreSQL 版本。

### 配置服务器

加载验证器并重启 PostgreSQL：

    oauth_validator_libraries = 'pg_oidc_validator'

向 pg_hba.conf 添加一个 OAuth 规则。发行商和范围必须与提供商匹配：

    host  all  all  127.0.0.1/32  oauth  issuer=https://id.example.com/realms/postgres scope="openid postgres"

编辑完 pg_hba.conf 后重新加载它。验证器会根据从发行商发现的提供者元数据检查令牌的发行商、受众、范围、签名和过期时间。

默认情况下，PostgreSQL 角色与 JWT 的 sub 声明进行匹配。要通过其他声明（如 email）进行身份验证，请设置：

    pg_oidc_validator.authn_field = 'email'

此设置更改了用于角色匹配的身份声明；它不会创建或配置数据库角色。

### 使用 libpq 连接

支持 OAuth 的 libpq 客户端可以启动设备授权流程：

    psql "host=127.0.0.1 dbname=app user=alice +      oauth_issuer=https://id.example.com/realms/postgres +      oauth_client_id=postgres-client"

仅在注册客户端需要时使用 oauth_client_secret。客户端标识符、重定向/设备流设置、受众和请求的范围必须与身份提供者配置一致。

### 配置索引

- oauth_validator_libraries: 服务器级别的 OAuth 验证模块列表；添加 pg_oidc_validator 需要重启。
- pg_oidc_validator.authn_field: 用于角色匹配的 JWT 声明，默认为 sub。
- pg_hba.conf oauth 方法: 选择 OAuth 认证并提供接受的发行商和范围。
- oauth_issuer, oauth_client_id, oauth_client_secret: libpq 连接参数，用于获取令牌。

### 提供者与安全边界

- 上游 0.2 文档针对 PostgreSQL 18，并要求使用支持 OAuth 的 libpq 客户端。
- 验证器支持常见的 OIDC 提供者，但 README 特别指出 Google 不受支持，并描述了 Microsoft Entra ID 的特定设置。
- 令牌验证只是授权的一部分。PostgreSQL 角色成员资格和对象权限仍然控制数据库访问。
- 尽可能在外连接字符串之外保护客户端密钥和提供者凭据，并验证对发行商的信任 TLS。
