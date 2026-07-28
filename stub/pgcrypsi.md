## Usage

Sources:

- [Official upstream README](https://github.com/telkomdev/pgcrypsi/blob/ee909d7ad315aa5c5178687db542e5fc73959d84/README.md)
- [Official extension control file (pgcrypsi.control)](https://github.com/telkomdev/pgcrypsi/blob/ee909d7ad315aa5c5178687db542e5fc73959d84/pgcrypsi.control)
- [Official extension SQL (pgcrypsi--0.0.1.sql)](https://github.com/telkomdev/pgcrypsi/blob/ee909d7ad315aa5c5178687db542e5fc73959d84/pgcrypsi--0.0.1.sql)

`pgcrypsi` — C Crypsi PostgreSQL Extension. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgcrypsi;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgcrypsi_aes_128_gcm_decrypt(text, text)` is an extension function and returns `text`.
- `pgcrypsi_aes_128_gcm_encrypt(text, text)` is an extension function and returns `text`.
- `pgcrypsi_aes_192_gcm_decrypt(text, text)` is an extension function and returns `text`.
- `pgcrypsi_aes_192_gcm_encrypt(text, text)` is an extension function and returns `text`.
- `pgcrypsi_aes_256_gcm_decrypt(text, text)` is an extension function and returns `text`.
- `pgcrypsi_aes_256_gcm_encrypt(text, text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
