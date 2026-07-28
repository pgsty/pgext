## Usage

Sources:

- [Official upstream README](https://github.com/khulnasoft/vault/blob/9f472c76a7cd7d9847175a6718853c1510fc911c/README.md)
- [Official extension control file (khulnasoft_vault.control)](https://github.com/khulnasoft/vault/blob/9f472c76a7cd7d9847175a6718853c1510fc911c/khulnasoft_vault.control)
- [Official extension SQL (khulnasoft_vault--0.2.8.sql)](https://github.com/khulnasoft/vault/blob/9f472c76a7cd7d9847175a6718853c1510fc911c/sql/khulnasoft_vault--0.2.8.sql)

`khulnasoft_vault` — Khulnasoft provides a table called vault.secrets that can be used to store sensitive information like API keys. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION khulnasoft_vault;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `vault.create_secret(new_secret text, new_name text = NULL, new_description text = '', new_key_id uuid = NULL)` is an extension function and returns `uuid`.
- `vault.update_secret(secret_id uuid, new_secret text = NULL, new_name text = NULL, new_description text = NULL, new_key_id uuid = NULL)` is an extension function and returns `void`.
- `vault.secrets` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.8`.
- Install the confirmed extension dependencies first: `pgsodium`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
