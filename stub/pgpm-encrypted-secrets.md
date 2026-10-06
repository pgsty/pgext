## Usage

Sources:

- [packages/encrypted-secrets/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/README.md)
- [packages/encrypted-secrets/sql/pgpm-encrypted-secrets--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/sql/pgpm-encrypted-secrets--0.47.0.sql)
- [packages/encrypted-secrets/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/Makefile)
- [packages/encrypted-secrets/pgpm-encrypted-secrets.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets/pgpm-encrypted-secrets.control)

`pgpm-encrypted-secrets` adds get, upsert, verify and delete operations over the separate secrets-table extension, with PGP encoding and crypt hashing modes.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-encrypted-secrets" CASCADE;
SELECT encrypted_secrets.secrets_getter('00000000-0000-0000-0000-000000000001'::uuid, 'demo', NULL);
```

### Operational Boundaries

`encrypted_secrets.secrets_upsert`, `encrypted_secrets.secrets_verify` and `encrypted_secrets.secrets_delete` take an owner UUID and secret name. The PGP implementation derives its passphrase from the owner UUID stored with the row; it does not protect ciphertext from someone who can read that identifier. Enforce application authorization and SQL grants independently, and never log real secret arguments.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `pgcrypto`, `plpgsql`, `pgpm-encrypted-secrets-table`, `pgpm-verify`. Its SQL expects the platform roles `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
