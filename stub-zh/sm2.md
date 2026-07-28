## 用法

来源：

- [官方上游 README](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm2_c/README.md)
- [官方扩展控制文件 (sm2.control)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm2_c/sm2.control)
- [官方扩展 SQL (sm2--1.0.sql)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm2_c/sm2--1.0.sql)

`sm2` — SM2 椭圆曲线加密函数（中国国家标准 GB/T 32918）。在实现相应的安全、审计或访问控制流程时使用。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION sm2;
```

在目标数据库中安装扩展，如果有可用示例，请运行上游最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `sm2_c_decrypt(ciphertext bytea, private_key text)` 是一个扩展函数，返回 `text`。
- `sm2_c_decrypt_base64(ciphertext_base64 text, private_key text)` 是一个扩展函数，返回 `text`。
- `sm2_c_decrypt_hex(ciphertext_hex text, private_key text)` 是一个扩展函数，返回 `text`。
- `sm2_c_encrypt(plaintext text, public_key text)` 是一个扩展函数，返回 `bytea`。
- `sm2_c_encrypt_base64(plaintext text, public_key text)` 是一个扩展函数，返回 `text`。
- `sm2_c_encrypt_hex(plaintext text, public_key text)` 是一个扩展函数，返回 `text`。
- `sm2_c_generate_key()` 是一个扩展函数，返回 `text[]`。
- `sm2_c_get_pubkey(private_key text)` 是一个扩展函数，返回 `text`。
- `sm2_c_sign(message text, private_key text, id text DEFAULT NULL)` 是一个扩展函数，返回 `bytea`。
- `sm2_c_sign_hex(message text, private_key text, id text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm2_c_verify(message text, public_key text, signature bytea, id text DEFAULT NULL)` 是一个扩展函数，返回 `boolean`。
- `sm2_c_verify_hex(message text, public_key text, signature_hex text, id text DEFAULT NULL)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
