## Usage

Sources:

- [pg_pii_vault.control](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/pg_pii_vault.control)
- [Cargo.toml](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/Cargo.toml)
- [README.md](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/README.md)
- [UPGRADING.md](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/UPGRADING.md)
- [sql/pg_pii_vault--0.0.0--0.1.0.sql](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/sql/pg_pii_vault--0.0.0--0.1.0.sql)
- [sql/pg_pii_vault--0.1.0--0.1.1.sql](https://github.com/g0ddest/pg_pii_vault/blob/306482c1134317c3e503ba9c9cf8ec1f4b332fe1/sql/pg_pii_vault--0.1.0--0.1.1.sql)

`pg_pii_vault` 0.1.1 provides the `piitext` type for explicit column encryption using per-subject keys in HashiCorp Vault Transit. A SELECT-only role sees ciphertext; decryption needs a separately granted function. PostgreSQL 14–18 and superuser installation are supported.

### Core Workflow

```ini
shared_preload_libraries = 'pg_pii_vault'
pii_vault.url = 'https://vault.example.internal:8200'
pii_vault.token_file = '/run/vault-agent/pg_pii_vault.token'
pii_vault.mount = 'transit'
```

```sql
CREATE EXTENSION pg_pii_vault;
SELECT check_name, ok, required FROM piitext_vault_check();
SELECT piitext_encrypt('test value', int4send(123))::text;
```

### Operational Boundaries

Configure a least-privilege Vault token file and TLS, preload `pg_pii_vault`, and restart before use. Preload supplies cluster-wide key-cache invalidation and protects the token setting. The example assumes an initialized Transit mount and uses a test subject key. All `pii_vault.*` settings are superuser-only. Default export mode brings exportable keys into backend memory; transit mode delegates encryption to Vault with non-exportable keys. `piitext_shred` deletes a subject key irreversibly; `piitext_cache_invalidate`, `piitext_stats`, `piitext_vault_check`, `piitext_key_version` and `piitext_reencrypt` support operations. Network or permission failures raise distinct errors; only a confirmed missing key/version produces the deletion marker. Set timeouts and preserve Vault backups along with encrypted database data.

Version 0.1 is incompatible with old application assumptions: implicit plaintext writes and implicit decryption are removed, encryption/decryption/admin functions lose PUBLIC execution, and stored decrypting expressions can retain plaintext. Before upgrading 0.0.0, follow the full upstream inventory and backup procedure, adapt queries, move secrets out of role/database settings, configure preload and restart, then run ALTER EXTENSION in every affected database. Replace potentially exposed old tokens, explicitly encrypt existing staging values, and remove or redesign plaintext-derived indexes, generated columns, materialized views and statistics as directed. Historical WAL, dumps, replicas and Vault snapshots remain separate retention concerns. Once staging data is migrated, `pii_vault.allow_staging = off` can reject it. Upgrade 0.1.0 to 0.1.1 requires matching files, restart and SQL update, but no value rewrite. There is no automatic downgrade; old libraries cannot read new stored formats.
