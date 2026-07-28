## 用法

来源：

- [官方上游 README](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm4_c/README.md)
- [官方扩展控制文件 (sm4.control)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm4_c/sm4.control)
- [官方扩展 SQL (sm4--1.0.sql)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm4_c/sm4--1.0.sql)

`sm4` — SM4 加密/解密函数（中国国家标准）。在实现相应的安全、审计或访问控制流程时使用。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION sm4;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `sm4_c_decrypt(ciphertext bytea, key text)` 是一个扩展函数，返回 `text`。
- `sm4_c_decrypt_cbc(ciphertext bytea, key text, iv text)` 是一个扩展函数，返回 `text`。
- `sm4_c_decrypt_gcm(ciphertext_with_tag bytea, key text, iv text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm4_c_decrypt_gcm_auto_iv(ciphertext bytea, key text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm4_c_decrypt_gcm_auto_iv_base64(ciphertext_base64 text, key text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm4_c_decrypt_gcm_base64(ciphertext_base64 text, key text, iv text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm4_c_decrypt_hex(ciphertext_hex text, key text)` 是一个扩展函数，返回 `text`。
- `sm4_c_encrypt(plaintext text, key text)` 是一个扩展函数，返回 `bytea`。
- `sm4_c_encrypt_cbc(plaintext text, key text, iv text)` 是一个扩展函数，返回 `bytea`。
- `sm4_c_encrypt_gcm(plaintext text, key text, iv text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `bytea`。
- `sm4_c_encrypt_gcm_auto_iv(plaintext text, key text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `bytea`。
- `sm4_c_encrypt_gcm_auto_iv_base64(plaintext text, key text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm4_c_encrypt_gcm_base64(plaintext text, key text, iv text, aad text DEFAULT NULL)` 是一个扩展函数，返回 `text`。
- `sm4_c_encrypt_hex(plaintext text, key text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
