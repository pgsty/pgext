## Usage

Sources:

- [Official upstream README](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm2_c/README.md)
- [Official extension control file (sm2.control)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm2_c/sm2.control)
- [Official extension SQL (sm2--1.0.sql)](https://github.com/pangpang20/vastbase_sm4/blob/de11c211179848af08530559ecadd1b722f918a4/sm2_c/sm2--1.0.sql)

`sm2` — SM2 elliptic curve cryptography functions (Chinese National Standard GB/T 32918). Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION sm2;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `sm2_c_decrypt(ciphertext bytea, private_key text)` is an extension function and returns `text`.
- `sm2_c_decrypt_base64(ciphertext_base64 text, private_key text)` is an extension function and returns `text`.
- `sm2_c_decrypt_hex(ciphertext_hex text, private_key text)` is an extension function and returns `text`.
- `sm2_c_encrypt(plaintext text, public_key text)` is an extension function and returns `bytea`.
- `sm2_c_encrypt_base64(plaintext text, public_key text)` is an extension function and returns `text`.
- `sm2_c_encrypt_hex(plaintext text, public_key text)` is an extension function and returns `text`.
- `sm2_c_generate_key()` is an extension function and returns `text[]`.
- `sm2_c_get_pubkey(private_key text)` is an extension function and returns `text`.
- `sm2_c_sign(message text, private_key text, id text DEFAULT NULL)` is an extension function and returns `bytea`.
- `sm2_c_sign_hex(message text, private_key text, id text DEFAULT NULL)` is an extension function and returns `text`.
- `sm2_c_verify(message text, public_key text, signature bytea, id text DEFAULT NULL)` is an extension function and returns `boolean`.
- `sm2_c_verify_hex(message text, public_key text, signature_hex text, id text DEFAULT NULL)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
