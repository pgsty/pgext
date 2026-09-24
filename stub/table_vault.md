## Usage

Sources:

- [README](https://github.com/2ndQuadrant/rls-examples/blob/6a5b6d83e87ac682a0e2c7221ae2c2d31e5d9059/table-vault/README.md)
- [Control](https://github.com/2ndQuadrant/rls-examples/blob/6a5b6d83e87ac682a0e2c7221ae2c2d31e5d9059/table-vault/table_vault.control)
- [SQL](https://github.com/2ndQuadrant/rls-examples/blob/6a5b6d83e87ac682a0e2c7221ae2c2d31e5d9059/table-vault/table_vault--1.0.sql)

`table_vault` is a historical SQL demonstration of application identity context for row-level security. Version 1.0 comes from the 2016 rls-examples source snapshot; it is an example to review and adapt, not a production authentication framework.

### Core Workflow

Install the files, enable `pgcrypto`, then create `table_vault` as an administrator. PL/pgSQL is required. Objects live in the fixed `table_vault` schema; there is no shared library or preload step. In this psql example, supply the `vault_passphrase` variable securely before running it, and initialize the key table only once.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION table_vault;
INSERT INTO table_vault.key_vault
VALUES (crypt(:'vault_passphrase', gen_salt('bf')));
SELECT table_vault.set_username('app_user', :'vault_passphrase');
SELECT table_vault.get_username();
```

The setter checks the shared passphrase and stores a random UUID in `signed_vault.session_id`, resolved through its private `sessions` table. It returns the context token. The getter returns the application username only while the context passes validation and is less than one day old; missing or expired context raises an error.

### Objects and Operational Boundaries

- `set_username(text,text)` and `get_username()` are publicly executable security-definer functions; possession of the shared passphrase authorizes setting an application username.
- `key_vault` holds the protected credential material. The SQL revokes public access and grants schema usage for the API. Restrict administrative table access and protect the credential when sending SQL.
- The functions do not set a fixed `search_path` and call unqualified helper functions. Review owner privileges and schema trust before allowing untrusted SQL callers.
- Session context persists across transactions. A connection pool must establish the correct context for every checkout and must not leak a previous caller's token.
- No current PostgreSQL compatibility matrix or explicit license is supplied. Keep this historical demonstration separate from claims of modern security or support.

Expired rows are rejected when read, but the extension does not schedule deletion from `sessions`; its owner must plan cleanup.
