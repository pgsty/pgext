## Usage

Sources:

- [Official extension control file](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/pg_pii_vault.control)

`pg_pii_vault` — PostgreSQL extension for GDPR-compliant column-level encryption using HashiCorp Vault Transit Engine.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_pii_vault";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_pii_vault';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
