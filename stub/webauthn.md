## Usage

Sources:

- [Official extension control file](https://github.com/truthly/pg-webauthn/blob/c1f8099db1cbed3a89cf12a2668ed67bb10ee354/webauthn.control)

`webauthn` — Pure-SQL PostgreSQL implementation of WebAuthn credential registration and assertion verification.

The reviewed catalog snapshot records version `1.6`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `cbor`, `pg_ecdsa_verify`, `pgcrypto`, `plpgsql`.

```sql
CREATE EXTENSION "webauthn";
SELECT extversion
FROM pg_extension
WHERE extname = 'webauthn';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
