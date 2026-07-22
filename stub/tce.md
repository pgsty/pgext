## Usage

Sources:

- [Official project documentation](https://the-cryptic-elephant.readthedocs.io/en/latest/)
- [Official administration notes](https://gitlab.com/dalibo/the_cryptic_elephant/-/blob/latest/docs/admin.md)
- [Official driver notes](https://gitlab.com/dalibo/the_cryptic_elephant/-/blob/latest/docs/drivers.md)

`tce`, The Cryptic Elephant, is an early transparent column-encryption extension. It encrypts values in PostgreSQL memory and storage with per-role envelope keys backed by external key management, but upstream explicitly says not to use this `0.10.0-dev` line in production.

### Load and Create the Extension

The library must be loaded before the extension is created so that its hooks, GUCs, and security-label provider exist.

```sql
ALTER DATABASE demo SET session_preload_libraries TO 'tce';
-- Reconnect to demo before continuing.
CREATE EXTENSION tce;
```

For an isolated test only, allow file-based private-key retrieval, initialize a test superuser's envelope key, and convert a column to `tcetext`.

```sql
CREATE ROLE alice LOGIN SUPERUSER;
ALTER DATABASE demo
  SET tce.allow_unsecured_key_fetch_methods TO true;

\connect demo alice

SECURITY LABEL FOR tce_user_key ON ROLE alice
  IS 'FETCHED FROM FILE /secure/alice-private.pem';

CREATE TABLE secrets (id bigint PRIMARY KEY, secret text);
INSERT INTO secrets VALUES (1, 'classified');

ALTER TABLE secrets
  ALTER COLUMN secret TYPE tcetext
  USING secret::tcetext;

SELECT secret, raw(secret) FROM secrets;
```

The file method is intentionally disabled by default and is documented for testing only. Use a supported KMS design for key-encryption keys, and protect the database server and superuser boundary.

### Main Objects and Settings

- `tcetext` and `tceint` store encrypted text and integer values; `raw` exposes ciphertext for inspection.
- `tce.transparent_column_encryption` enables DDL rewriting so newly created regular columns are converted automatically.
- `tce.data_encryption_method` selects the data cipher; the reviewed source implements ChaCha20-Poly1305.
- `tce.encrypted_data_display_mode` chooses `error`, `raw`, or `empty` when a role cannot decrypt.
- `tce.safe_lock` protects the data-encryption-key envelope from accidental removal.
- `tce.kms.type` and provider-specific settings configure AWS or OpenBao key retrieval.
- `tce.autocast` assists binary-protocol drivers, but its GUC definition warns that it is for testing only.

### Security and Administration Boundary

The design assumes trusted system administrators and PostgreSQL superusers. Security labels, GUCs, and the `tce.envelope` table can appear in dumps. Logical dumps and logical replication decrypt values before they leave the publisher; physical backups retain encrypted storage. Restore targets need the extension loaded, the TCE types available, and usable role keys.

Roles without a key receive the configured display behavior. Key changes apply to new sessions, so terminate old sessions when revocation must be immediate. Test drivers explicitly: binary-protocol clients do not work transparently without `tce.autocast`, while enabling it is itself marked non-production by the current source.
