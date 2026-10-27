## 用法

来源：

- [README.md](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/README.md)
- [jwt_context.control](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/jwt_context.control)
- [sql/jwt_context--1.0.0.sql](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/sql/jwt_context--1.0.0.sql)
- [jwt_context.c](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/jwt_context.c)
- [.github/workflows/ci.yml](https://github.com/tvondra/jwt_context/blob/10be23c0651f10993e2df2a316af926792e0caa0/.github/workflows/ci.yml)

`jwt_context` 1.0.0 为 SQL 和行安全策略提供实验性签名会话上下文，接受 JSON 上下文之前验证 HS256 或 ES256 签名。所核对的 CI 覆盖 PostgreSQL 14–17，应由超级用户安装，扩展允许迁移模式。

### 读取已验证的上下文

```sql
CREATE EXTENSION jwt_context;
SELECT jwt();
SELECT jwt('TENANT_ID');
RESET jwt.context;
```

可信管理员为 HS256 设置 `jwt.secret`，或为 ES256 设置 `jwt.pubkey`；两者都是仅超级用户可改的配置，保存 base64url 编码的密钥材料。应用将 `jwt.context` 设为可信签发者生成的有效令牌。`jwt()` 返回 JSONB 声明，`jwt(key)` 以文本读取单个声明。`jwt_sign`、`jwt_verify`、`jwt_generate_keys`、`jwt_generate_secret` 用于实验，必须保护签名密钥，不能把签名权限暴露给应用用户。

### 安全与连接池

应用应以普通角色连接：超级用户或可替换验证密钥的用户都能伪造上下文。签名有效不等于业务授权成立。原型没有强制验证令牌过期或撤销状态，依赖令牌前应在外部实现这些检查，泄漏的签名令牌可以被重放。

上下文属于会话状态。连接池把连接交给另一用户之前必须重置，按需使用 RESET 或 RESET ALL，并测试连接池的实际重置行为。将这一实验性模块用于行安全前，应审阅上游威胁模型。
