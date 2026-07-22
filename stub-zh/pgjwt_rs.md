## 用法

来源：

- [官方 README](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/README.md)
- [Rust 实现](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/src/lib.rs)
- [扩展 SQL](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/sql/pgjwt_rs--0.1.2.sql)
- [Rust 包清单](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/Cargo.toml)
- [扩展控制文件](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/pgjwt_rs.control)

`pgjwt_rs` 0.1.2 在 PostgreSQL 内使用 RS256 或 Ed25519 验证非对称 JWT 签名。它有意只做密码学验证：过期时间、生效时间、受众、签发者、主体、nonce 与应用授权仍由调用者负责。

### 先验证再使用声明

应把可信公钥和服务端选定的算法与不可信令牌分开存放。通过可信的密钥标识或签发者映射把令牌关联到密钥，要求 `valid` 为真，再应用程序所需的全部声明检查。例如，假定应用自有 `auth.incoming_token` 与 `auth.jwt_key` 表：

```sql
CREATE EXTENSION pgjwt_rs;

SELECT v.header, v.payload
FROM auth.incoming_token AS i
JOIN auth.jwt_key AS k ON k.key_id = i.key_id
CROSS JOIN LATERAL jwt_verify(i.token, k.public_key_pem, k.algorithm) AS v
WHERE i.id = 42
  AND v.valid
  AND (v.payload->>'iss') = 'https://issuer.example'
  AND (v.payload->>'aud') = 'orders-api'
  AND now() < to_timestamp((v.payload->>'exp')::bigint);
```

不要只根据令牌头选择 `k.algorithm`。验证器会检查已签名头部是否与预期算法一致，但应用必须从可信配置提供这个预期值。

### 用户函数

- `jwt_verify_rs256(text, text)`：用 SPKI PEM 公钥验证 RS256 签名，返回 `header`、`payload` 与 `valid`。
- `jwt_verify_ed25519(text, text)`：用 SPKI PEM 公钥验证 Ed25519/EdDSA 签名。
- `jwt_verify(text, text, text)`：按可信代码提供的 `RS256`、`EdDSA` 或 `ED25519` 分派。
- `jwt_decode_payload(text)`：不检查签名而解码载荷 JSON；返回对象带有 `_unverified` 标记。

公钥必须使用 `PUBLIC KEY` PEM 标签，不能使用 `RSA PUBLIC KEY`、`RSA PRIVATE KEY` 或其他容器。超过 16 KiB 的令牌和超过 8 KiB 的密钥会被拒绝。

### 安全边界

`jwt_decode_payload` 只能作为查找候选可信密钥的解析辅助。绝不能根据其结果直接授权、选择租户，或构造不受约束的密钥查询。即使 `valid` 为假，验证函数仍会解码并返回载荷，因此调用者必须先检查 `valid`，再读取任何声明。

0.1.2 版不强制检查 `exp`、`nbf` 或 `aud`，也不要求存在 `exp`。请显式添加时间声明、签发者、受众、主体、令牌类型、时钟偏差、密钥轮换与吊销策略。若用 `SECURITY DEFINER` 包装这些函数，应固定安全的 `search_path`、使用带模式限定的对象，并撤销不可信角色的默认执行权限。该包宣称支持 PostgreSQL 13 至 18；请在目标服务器验证准确的 pgrx 制品。
