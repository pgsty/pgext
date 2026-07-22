## Usage

Sources:

- [Official README](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/README.md)
- [Official SQL API implementation](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/src/lib.rs)
- [Official Vault integration](https://github.com/g0ddest/pg_pii_vault/blob/34256629bf8f1af5c4a2437452387ed9d156ee3e/src/vault.rs)

`pg_pii_vault` is an early-stage `piitext` type that encrypts text with AES-256-GCM using per-record key identifiers and obtains exportable encryption keys from HashiCorp Vault Transit. Upstream explicitly says it is not ready for production use.

### Core Workflow

The extension is untrusted and superuser-only. Configure a Vault Transit endpoint and narrowly scoped token, then encrypt explicitly with a stable record key identifier:

```sql
CREATE EXTENSION pg_pii_vault;

SET pii_vault.url = 'https://vault.internal';
SET pii_vault.token = '<vault-token>';
SET pii_vault.mount = 'transit';

CREATE TABLE private_users (
  id integer PRIMARY KEY,
  secret piitext NOT NULL
);

INSERT INTO private_users(id, secret)
VALUES (
  123,
  piitext_encrypt('sensitive value', decode('0000007b', 'hex'))
);

SELECT id, piitext_out_text(secret)
FROM private_users;
```

`piitext_debug` exposes the stored structure and `piitext_raw` exposes raw CBOR bytes; keep execution on both restricted. `piitext_encrypt_piitext` decrypts and re-encrypts an existing value under a new key identifier.

### Critical Security Boundaries

- The implicit `text` to `piitext` cast calls `piitext_in_text` and stores a plaintext staging value; encryption is not automatic. Only rows passed through `piitext_encrypt` or `piitext_encrypt_piitext` are sealed.
- The code asks Vault Transit to create keys as `exportable`, exports the raw encryption key, and performs AES-GCM in the PostgreSQL backend. The database process and its memory therefore join the key trust boundary; this is not normal non-exporting Transit encrypt/decrypt behavior.
- `pii_vault.url`, `pii_vault.token`, `pii_vault.mount`, and `pii_vault.cache_ttl_sec` are `USERSET`. Unless function/GUC use is tightly controlled, callers can select an endpoint and credential and trigger blocking server-side HTTP requests. Revoke public execution and outbound network access by default.
- The key cache is process-local backend memory, not PostgreSQL shared memory. Deleting a Vault key does not take effect in a backend until its cached key expires, and an outage, auth failure, deleted key, corrupt ciphertext, or tag failure all cause `piitext_out_text` to return the same `****` marker.
- `piitext_encrypt` and `piitext_out_text` are declared `IMMUTABLE` even though they depend on random IVs, GUCs, network/Vault state, and cache state. Planner constant folding and expression-index/generated-column assumptions are unsafe.

### Operational Limits

- Vault calls are synchronous inside the SQL backend. Set network and statement timeouts externally, isolate failures, and measure connection-pool multiplication and cache behavior.
- AAD is derived from the type name and key identifier, not from an independently verified table primary key. The application must keep the key identifier bound to the intended record and prevent ciphertext swapping within the same identifier.
- Automatic triggers, key rotation, background cache cleanup, and production integration tests are roadmap items. Mock mode uses a fixed all-zero key and is test-only.
- “Crypto shredding” is not proof of regulatory deletion by itself. Account for caches, Vault versions/backups, database backups, replicas, WAL, exports, logs, and plaintext staging rows in the deletion design.
