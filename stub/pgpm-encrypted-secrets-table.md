## Usage

Sources:

- [packages/encrypted-secrets-table/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/README.md)
- [packages/encrypted-secrets-table/sql/pgpm-encrypted-secrets-table--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/sql/pgpm-encrypted-secrets-table--0.47.0.sql)
- [packages/encrypted-secrets-table/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/Makefile)
- [packages/encrypted-secrets-table/pgpm-encrypted-secrets-table.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/encrypted-secrets-table/pgpm-encrypted-secrets-table.control)

`pgpm-encrypted-secrets-table` provides `secrets_schema.secrets_table` and triggers used by the higher-level secrets module. It stores owner IDs, names and encoded values.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-encrypted-secrets-table" CASCADE;
SELECT count(*) FROM secrets_schema.secrets_table;
```

### Operational Boundaries

The PGP path uses the owner UUID as its passphrase, not an independently protected key. A reader with the stored identifiers and ciphertext can therefore decrypt it. Crypt values are one-way password hashes. This design is not a vault or protection against a database reader; restrict table/function access and avoid treating identifiers as secrets.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `pgcrypto`, `plpgsql`, `pgpm-verify`. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
