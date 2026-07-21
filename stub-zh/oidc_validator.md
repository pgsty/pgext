## 用法

来源：

- [官方README](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/b65bbbe288f84fab91d58b8304e8a526d1326af5/README.md)
- [验证器配置源码](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/b65bbbe288f84fab91d58b8304e8a526d1326af5/src/config.rs)
- [PostgreSQL OAuth回调实现](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/b65bbbe288f84fab91d58b8304e8a526d1326af5/src/ffi.rs)
- [PostgreSQL 18 OAuth认证文档](https://www.postgresql.org/docs/18/auth-oauth.html)

`oidc_validator` 是一个用 Rust 编写的 PostgreSQL 18 OAuth 验证模块。它会验证 JWT 访问令牌并将其主题作为已认证的身份返回。这是一个无头的认证库，而不是 SQL 扩展，因此不会创建任何 SQL 对象，也不会使用 `CREATE EXTENSION`。

### 核心工作流程

将 `oidc_validator.so` 安装到 PostgreSQL 的库目录中，然后配置 PostgreSQL 18 验证模块：

```conf
oauth_validator_libraries = 'oidc_validator'
```

在 `pg_hba.conf` 中添加一个 OAuth 规则：

```conf
host all all 0.0.0.0/0 oauth issuer="https://issuer.example" scope="openid profile"
```

向 PostgreSQL 服务器进程提供验证器配置：

```shell
POSTGRES_OIDC_ISSUER=https://issuer.example
POSTGRES_OIDC_CLIENT_ID=postgres
POSTGRES_OIDC_AUDIENCE=postgres
```

更改 `oauth_validator_libraries` 或服务器进程环境后重启 PostgreSQL。然后，OAuth 客户端可以通过匹配的 `pg_hba.conf` 规则进行认证。

### 配置索引

- `POSTGRES_OIDC_ISSUER`: 不带已知发现后缀的发行人 URL。
- `POSTGRES_OIDC_CLIENT_ID`: OIDC 应用客户端 ID。
- `POSTGRES_OIDC_AUDIENCE`: 必须的令牌受众，通常为客户端 ID。
- `oauth_validator_libraries`: PostgreSQL 18 设置项，用于加载受信任验证模块。

### 要求和注意事项

- 上游版本 `0.1.0` 针对 PostgreSQL 18，并要求在编译 PostgreSQL 时启用 OpenSSL 和 libcurl。
- 只接受 JWT 样式的凭据令牌。不接受密文访问令牌。
- 验证会执行 OIDC 发现和 JWKS 获取操作，因此 PostgreSQL 服务器必须能够通过 TLS 访问发行人。
- 当前回调忽略请求的 PostgreSQL 角色，并基于成功的令牌验证进行授权。不要使用此实现设置 `delegate_ident_mapping=1`；保持 PostgreSQL 的标准精确名称映射或 `pg_ident.conf` 映射，以便返回的令牌主题可以与请求的角色进行检查。
