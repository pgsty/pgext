## 用法

来源：

- [上游 README 与 SQL 签名](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/README.md)
- [扩展 control 文件](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/pg_ecdsa_verify.control)
- [Cargo 包与 PostgreSQL feature 矩阵](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/Cargo.toml)
- [Rust 实现](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/src/lib.rs)

`pg_ecdsa_verify` 提供一个基于 Rust/pgrx 的 ECDSA 签名验证函数。`ecdsa_verify.ecdsa_verify(bytea, bytea, bytea, text, text)` 会对载荷计算哈希、用公钥验证签名并返回布尔值。它支持 `secp256r1` 与 `secp256k1` 曲线，哈希选择器仅支持 `sha256`。

### 验证已存储的签名数据

```sql
CREATE EXTENSION pg_ecdsa_verify;

SELECT ecdsa_verify.ecdsa_verify(
  public_key,
  payload,
  signature,
  'sha256',
  'secp256r1'
)
FROM signed_messages
WHERE message_id = 42;
```

公钥是 X 与 Y 坐标拼接成的原始 64 字节值，签名是 R 与 S 值拼接成的原始 64 字节值。该扩展有意不生成密钥或创建签名，因此私钥无需进入 PostgreSQL。

实现会在没有长度检查的情况下切分两个字节数组，并在遇到不支持的曲线/哈希名称时 panic。调用前应校验 `octet_length(public_key) = 64` 和 `octet_length(signature) = 64`，并把畸形输入当作错误路径，而不是无效签名的 false 结果。目录与 Cargo 包的版本均为 `1.2.4`；Cargo feature 覆盖 PostgreSQL 13 至 18。安装仅限超级用户，并会创建固定的 `ecdsa_verify` schema。
