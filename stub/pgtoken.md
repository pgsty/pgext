## Usage

Sources:

- [Official upstream README](https://github.com/ajayr4j/pgtoken/blob/323ed299774f92d73cbdaa6b73a7158a0b76ca09/README.md)
- [Official extension control file (pgtoken.control)](https://github.com/ajayr4j/pgtoken/blob/323ed299774f92d73cbdaa6b73a7158a0b76ca09/pgtoken.control)
- [Official extension SQL (pgtoken--1.0.sql)](https://github.com/ajayr4j/pgtoken/blob/323ed299774f92d73cbdaa6b73a7158a0b76ca09/pgtoken--1.0.sql)

`pgtoken` — PostgreSQL extension for rank-varint token storage. Store LLM token IDs compactly as bytea. No re-tokenization on read. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgtoken;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgtoken_count(encoded bytea)` is an extension function and returns `integer`.
- `pgtoken_decode(encoded bytea, codebook text DEFAULT 'cl100k_base')` is an extension function and returns `integer[]`.
- `pgtoken_encode(token_ids integer[], codebook text DEFAULT 'cl100k_base')` is an extension function and returns `bytea`.
- `pgtoken_reload_codebooks()` is an extension function and returns `void`.
- `pgtoken_codebooks` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
