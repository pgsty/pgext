

## Usage

Sources:

- [PGXN release 0.3.0](https://pgxn.org/dist/hyperion_vault/0.3.0/)
- [REST API reference](https://pgxn.org/dist/hyperion_vault/0.3.0/docs/API.html)
- [Security guidance](https://pgxn.org/dist/hyperion_vault/0.3.0/docs/SECURITY.html)

`hyperion_vault` is a PostgreSQL 18 secrets vault that keeps only ciphertext, a KMS-wrapped data-encryption key, and a nonce in PostgreSQL and WAL. The packaged `0.3.0` release is marked unstable upstream and combines a pgrx extension with a co-located REST API.

### Enable the extension

The rotation supervisor is a background worker, so add `hyperion_vault` to the existing `shared_preload_libraries` list and restart PostgreSQL before creating the extension:

```sql
SHOW shared_preload_libraries;

CREATE EXTENSION hyperion_vault;
SELECT extversion
FROM pg_extension
WHERE extname = 'hyperion_vault';
```

The extension creates the `vault` schema. For the API's dedicated database login, grant the service role through the provided helper:

```sql
CREATE ROLE vault_service LOGIN PASSWORD 'use-a-managed-secret';
SELECT vault.grant_service_role('vault_service');
```

The extension GUCs include `hyperion_vault.rotation_enabled`, `hyperion_vault.scan_interval_secs`, and `hyperion_vault.database`. The REST service exposes create, read, update, delete, rotate, and verify operations under `/v1/secrets`.

### Security boundary

Production mode requires AWS KMS. The local master-key mode is only for development and testing. Secret reads are controlled by the `VAULT_ALLOWED_IPS` IPv4/CIDR allowlist and fail closed when it is empty; management operations require bearer tokens. Review the upstream security and threat-model documentation, terminate the API with TLS, and use verified TLS for remote PostgreSQL connections before deploying this preview release.
