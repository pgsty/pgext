## 用法

来源：

- [README.md](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/README.md)
- [Cargo.toml](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/Cargo.toml)
- [src/lib.rs](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/src/lib.rs)
- [snout_oauth.control](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/snout_oauth.control)

`snout_oauth` 0.1.0 是 PostgreSQL 18 的 OAuth 验证库，用本地公开 JWKS 文件验证 Snout 数据库令牌，不通过 HTTP 联系签发方，也不安装 SQL 对象。

### 核心用法

```ini
oauth_validator_libraries = 'snout_oauth'
snout_oauth.issuer = 'https://auth.example.com'
snout_oauth.audience = 'project-ref'
snout_oauth.keys_file = '/etc/postgresql/oauth-jwks.json'
```

### 运行边界

启用 OAuth 认证前须设置签发方、项目受众和公钥文件。文件必须包含公开的 EC P-256 密钥，私钥会被拒绝。令牌必须通过签发方、受众、签名、有效期和生命周期检查，`token_use` 必须为 `db`，`db_role` 必须与请求的数据库角色一致。它不适用于任意 OIDC 令牌。

修改后重载服务器配置和认证规则。`oauth_validator_libraries` 按需载入库，共享预加载为可选项，不支持 `CREATE EXTENSION`；控制文件仅供开发工具使用。验证库自行校验数据库角色，因此可以委托角色映射。HBA 规则应仅匹配预期网络和用户。

密钥文件标识或元数据变化时会被重新读取，轮换时应原子替换文件。默认时钟宽限为 30 秒，令牌最长生命周期为一天；必填设置为空时拒绝认证。客户端也须具备 PostgreSQL 18 OAuth 能力。
