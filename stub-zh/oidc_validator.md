## 用法

来源：

- [上游 README](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/master/README.md)
- [Cargo 包元数据](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/master/Cargo.toml)
- [PostgreSQL OAuth 验证器回调实现](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/master/src/ffi.rs)

`oidc_validator` 是使用 Rust 编写的 PostgreSQL 18 OAuth 持有者令牌验证器。它通过 OIDC 发现机制与颁发者的 JWKS 端点校验 JWT 访问令牌，并把令牌主体作为认证身份返回。上游版本为 `0.1.0`，已在 Ubuntu 与 Arch Linux x86-64 上测试。

这是不提供 SQL 接口的认证模块，并非 SQL 扩展。编译并安装 `oidc_validator.so`，再通过 PostgreSQL 18 OAuth 验证器参数加载：

```conf
oauth_validator_libraries = 'oidc_validator'
```

在 `pg_hba.conf` 中加入 OAuth 规则，并为 PostgreSQL 服务器进程提供该验证器所需的环境变量：

```conf
host all all 0.0.0.0/0 oauth issuer="https://issuer.example" scope="openid profile"
```

```shell
POSTGRES_OIDC_ISSUER=https://issuer.example
POSTGRES_OIDC_CLIENT_ID=postgres
POSTGRES_OIDC_AUDIENCE=postgres
```

修改验证器库参数后必须重启 PostgreSQL。上游构建要求 PostgreSQL 18 启用 OpenSSL 与 libcurl。认证时会通过网络获取 OIDC 发现信息和 JWKS 数据，因此服务器必须能通过 TLS 访问颁发者。该验证器只接受 JWT 形式的令牌，不创建任何 SQL 对象。
