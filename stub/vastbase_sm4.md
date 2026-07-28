## Usage

Sources:

- [Official upstream README](https://github.com/pangpang20/sm4prj/blob/b55e4567fd7b64cca2da238f10c9bd518021237a/vastbase-sm4-extension/README.md)
- [Official extension control file (vastbase_sm4.control)](https://github.com/pangpang20/sm4prj/blob/b55e4567fd7b64cca2da238f10c9bd518021237a/vastbase-sm4-extension/vastbase_sm4.control)
- [Official extension SQL (vastbase_sm4--1.0.sql)](https://github.com/pangpang20/sm4prj/blob/b55e4567fd7b64cca2da238f10c9bd518021237a/vastbase-sm4-extension/vastbase_sm4--1.0.sql)

`vastbase_sm4` — SM4 encryption/decryption functions for VastBase using Java implementation via JNI. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION vastbase_sm4;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `sm4_decrypt(cipher_text text, hex_key text)` is an extension function and returns `text`.
- `sm4_decrypt_base64(cipher_text text, base64_key text)` is an extension function and returns `text`.
- `sm4_encrypt(plain_text text, hex_key text)` is an extension function and returns `text`.
- `sm4_encrypt_base64(plain_text text, base64_key text)` is an extension function and returns `text`.
- `sm4_generate_key()` is an extension function and returns `text`.
- `sm4_extension_info` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
