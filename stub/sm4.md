## Usage

Sources:

- [Official upstream README](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm4_c/README.md)
- [Official extension control file (sm4.control)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm4_c/sm4.control)
- [Official extension SQL (sm4--1.0.sql)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm4_c/sm4--1.0.sql)

`sm4` — SM4 encryption/decryption functions (Chinese National Standard). Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION sm4;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `sm4_c_decrypt(ciphertext bytea, key text)` is an extension function and returns `text`.
- `sm4_c_decrypt_cbc(ciphertext bytea, key text, iv text)` is an extension function and returns `text`.
- `sm4_c_decrypt_gcm(ciphertext_with_tag bytea, key text, iv text, aad text DEFAULT NULL)` is an extension function and returns `text`.
- `sm4_c_decrypt_gcm_auto_iv(ciphertext bytea, key text, aad text DEFAULT NULL)` is an extension function and returns `text`.
- `sm4_c_decrypt_gcm_auto_iv_base64(ciphertext_base64 text, key text, aad text DEFAULT NULL)` is an extension function and returns `text`.
- `sm4_c_decrypt_gcm_base64(ciphertext_base64 text, key text, iv text, aad text DEFAULT NULL)` is an extension function and returns `text`.
- `sm4_c_decrypt_hex(ciphertext_hex text, key text)` is an extension function and returns `text`.
- `sm4_c_encrypt(plaintext text, key text)` is an extension function and returns `bytea`.
- `sm4_c_encrypt_cbc(plaintext text, key text, iv text)` is an extension function and returns `bytea`.
- `sm4_c_encrypt_gcm(plaintext text, key text, iv text, aad text DEFAULT NULL)` is an extension function and returns `bytea`.
- `sm4_c_encrypt_gcm_auto_iv(plaintext text, key text, aad text DEFAULT NULL)` is an extension function and returns `bytea`.
- `sm4_c_encrypt_gcm_auto_iv_base64(plaintext text, key text, aad text DEFAULT NULL)` is an extension function and returns `text`.
- `sm4_c_encrypt_gcm_base64(plaintext text, key text, iv text, aad text DEFAULT NULL)` is an extension function and returns `text`.
- `sm4_c_encrypt_hex(plaintext text, key text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
