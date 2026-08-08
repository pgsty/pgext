## Usage

Sources:

- [pg_vault_tde 1.7.0 README](https://api.pgxn.org/src/pg_vault_tde/pg_vault_tde-1.7.0/README.md)
- [pg_vault_tde v1.7.0 release](https://github.com/labmiriade/pg_vault_tde/releases/tag/v1.7.0)
- [pg_vault_tde 1.7 control file](https://api.pgxn.org/src/pg_vault_tde/pg_vault_tde-1.7.0/pg_vault_tde.control)
- [pg_vault_tde operator documentation](https://api.pgxn.org/src/pg_vault_tde/pg_vault_tde-1.7.0/doc/pg_vault_tde.md)

`pg_vault_tde` adds transparent tuple encryption for PostgreSQL 17 and 18 through the `encrypted_heap` table access method. It encrypts user-column data with AES-256-GCM before storage and manages per-relation data-encryption keys through HashiCorp Vault/OpenBao, a local PKCS#12 wallet, or—in v1.7—a PKCS#11 HSM. MVCC tuple headers remain plaintext.

### Configure and Install

```conf
shared_preload_libraries = 'pg_vault_tde'
pg_vault_tde.kms_provider = 'vault'
pg_vault_tde.vault_url = 'https://vault.example.com:8200'
pg_vault_tde.vault_transit_mount = 'transit'
pg_vault_tde.vault_key_name = 'pg-tde-dek'
pg_vault_tde.vault_ca_cert = '/etc/ssl/vault/ca.pem'
```

Configure Vault authentication through the documented token, AppRole, or Kubernetes settings without committing secrets to PostgreSQL configuration. Restart PostgreSQL, then create the extension:

```sql
CREATE EXTENSION pg_vault_tde;
SELECT * FROM pg_vault_tde_health_check();
```

`kms_provider` has no usable default and must be set explicitly. The extension requires OpenSSL 3 and libcurl in addition to PostgreSQL server files.

### Create an Encrypted Table

```sql
CREATE TABLE customer_secrets (
  id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
  email text,
  ssn text
) USING encrypted_heap;
```

Encryption is table-level: ordinary `heap` tables are unchanged. Tuple values, TOAST data, and WAL representations are encrypted; tuple headers required for MVCC remain visible.

### Indexes

Use `tde_btree` for equality lookup without storing plaintext keys:

```sql
CREATE UNIQUE INDEX customer_secrets_id_tde_idx
ON customer_secrets USING tde_btree (id);

CREATE INDEX customer_secrets_email_tde_idx
ON customer_secrets USING tde_btree (email);
```

`tde_btree` uses deterministic AES-256-SIV and supports equality, not range ordering or index-only scans. Other access methods on an `encrypted_heap` table are rejected by default because they would write plaintext index keys. `PRIMARY KEY` and `UNIQUE` table constraints still create native btree indexes and produce a warning; decide whether that exposure is acceptable before defining them.

### Integrity and Rotation

```sql
SELECT * FROM pg_vault_tde_verify_integrity('customer_secrets');
SELECT * FROM pg_vault_tde_encrypted_size('customer_secrets');

SELECT pg_vault_tde_rotate_online('customer_secrets', 1000);
SELECT * FROM pg_vault_tde_get_rotation_status('customer_secrets');

SELECT pg_vault_tde_rotate_kek();
```

Online DEK rotation re-encrypts a table in batches and rebuilds its `tde_btree` indexes. KEK rotation re-wraps per-table DEKs without rewriting tuples. Restrict these operations, monitor completion, and avoid concurrent key-catalog restoration.

### Provider and Backup Boundaries

- The local wallet defaults outside `PGDATA`; copy and protect it separately because plain `pg_basebackup` does not include it.
- Version 1.7 adds the `pkcs11` provider and `pg_vault_tde_pkcs11_keygen()`. The standalone `pg_dump_tde` and `pg_restore_tde` tools do not support PKCS#11 in this release.
- Plain `pg_dump` and `COPY ... TO` read decrypted rows and therefore produce plaintext without a warning. Use the supplied encrypted logical-backup tools where supported.
- Physical backups contain encrypted relation bytes and wrapped DEKs, but not the KEK. Provision access to the Vault/HSM or copy the local wallet separately. The key-sealing functions and `pg_basebackup_tde` wrapper can accompany a physical backup with a tamper-evident DEK bundle.

### Critical Caveats

- Never toggle `pg_vault_tde.enabled` while an `encrypted_heap` table contains rows written under the other setting. The extension does not rewrite existing rows and mixed formats can be silently misread as corruption.
- Ordinary indexes, statistics, logs, query results, client traffic, temporary work, and backups can expose plaintext outside the encrypted heap. TDE is one storage-layer control, not end-to-end encryption.
- `tde_btree` disables range semantics, and encrypted tables disable HOT updates in the current design; benchmark update-heavy workloads and index maintenance.
- Keep KMS credentials, wallet passphrases, HSM PINs, KEKs, sealed bundles, and restore procedures under separate access controls. A backup without the matching key path is unrecoverable.
- Package release 1.7.0 installs SQL extension version `1.7`, is not relocatable, requires preloading and a restart, and supports PostgreSQL 17-18 only.
