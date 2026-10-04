## Usage

Sources:

- [pg_vault_tde.control](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/pg_vault_tde.control)
- [README.md](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/README.md)
- [doc/pg_vault_tde.md](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/doc/pg_vault_tde.md)

`pg_vault_tde` distribution 1.7.2 encrypts table values with AES-256-GCM through `encrypted_heap`, using Vault/OpenBao, a local PKCS#12 wallet or PKCS#11. SQL/control version remains 1.7. PostgreSQL 17–18, OpenSSL 3 and libcurl are required. Configure the key provider and authentication, preload the library, and restart before creating it as a superuser.

### Core Workflow

```ini
shared_preload_libraries = 'pg_vault_tde'
```

```sql
CREATE EXTENSION pg_vault_tde;
SELECT * FROM pg_vault_tde_health_check();
CREATE TABLE customer_secrets (id bigint, secret text) USING encrypted_heap;
CREATE INDEX customer_secrets_id_idx ON customer_secrets USING tde_btree (id);
```

### Operational Boundaries

This example assumes a configured provider. `pg_vault_tde_health_check`, `pg_vault_tde_verify_integrity`, `pg_vault_tde_get_rotation_status` and encrypted-size helpers expose operational state. `tde_btree` supports encrypted equality lookup, not ordering/range or index-only scans. Numeric and nondeterministic-collation encrypted indexes are unsupported; audit old unique/exclusion constraints because their checks may be ineffective. Plain indexes can expose keys, and pg_dump/COPY produce decrypted output. Protect external wallets/KMS and use the matching encrypted-backup tools where supported. Tuple headers, statistics, logs and query results are outside this protection.

Before replacing an older library or restarting, follow the upstream recovery checklist: earlier rotated tables may depend on keys held only in memory; concurrent rotation may require copying data while still readable, and rotated out-of-line values need the specified rewrite. Version 1.7.0 TOAST upgrades also have a separate export requirement. Check rotated/partial indexes and wallet ownership. After upgrading to 1.7.2, every pre-existing encrypted table needs VACUUM FULL to migrate its tuple layout; reserve a maintenance window and extra disk space. When `pg_vault_tde.toast_custom_rmgr` is enabled, the WAL ID changes from 128 to 161: clean shutdown, coordinated primary/standby upgrade, slot drainage and a new base backup are required. There is no rolling upgrade across those WAL formats.

Run key-management operations one at a time. Table rotation keeps reads available but blocks writes for one transaction; logical slots must decode the rotation before a restart or another rotation. Version 1.7.2 tightens caller privileges, but upstream documents remaining SECURITY DEFINER wrapper limitations; follow the listed grants/revokes. Do not toggle `pg_vault_tde.enabled` on a populated encrypted table. These catalog updates do not perform an operational upgrade or change the Pigsty package baseline.
