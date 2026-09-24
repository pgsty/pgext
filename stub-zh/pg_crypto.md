## 用法

来源：

- [README](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/README.md)
- [Control file / 控制文件](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/pg_crypto.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-crypto/blob/a07210ac1f98e5a722843c7389afdd0cd26ebe3c/docs/SECURITY.md)

`pg_crypto` 在 PostgreSQL 14–18 中以 Rust 实现常见的 pgcrypto 风格函数和额外密码原语，与内置 `pgcrypto` 扩展属于不同项目。

### 核心工作流

```sql
CREATE EXTENSION pg_crypto;
SELECT encode(digest('hello', 'sha256'), 'hex');
WITH key AS (SELECT gen_random_bytes(32) AS value)
SELECT xchacha20poly1305_decrypt(
  xchacha20poly1305_encrypt('example'::bytea, value, 'context'::bytea),
  value, 'context'::bytea
) FROM key;
```

### 接口

`digest` 与 `hmac` 计算摘要和消息认证码，`crypt` 与 `gen_salt` 提供密码哈希兼容接口。对称和公钥 OpenPGP 接口包括 `pgp_sym_encrypt`、`pgp_sym_decrypt` 以及文本封装辅助函数。现代原语包括 `xchacha20poly1305_encrypt`、`xchacha20poly1305_decrypt`、`argon2id_hash` 和 `secretbox`。

认证加密会验证认证标签与关联数据。兼容接口仍包含传统裸加密函数，但调用者须自行正确处理密钥长度、初始向量与认证设计。

### 安全与安装

扩展可重定位，不需要预加载或 OpenSSL 运行库。由于 SQL 名称重叠，`pg_crypto` 与 `pgcrypto` 不能安装到同一模式。迁移既有密文或哈希前，应阅读固定版本的兼容性与安全文档。

密钥、明文和密码会到达数据库服务端，传输安全、SQL 日志、角色权限和备份都属于安全设计的一部分。API 兼容性并不等同于已经通过独立密码学审计。
