## Usage

Sources:

- [AlohaDB README](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [Control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vault/alohadb_vault.control)
- [SQL](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vault/alohadb_vault--1.0.sql)
- [Source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vault/alohadb_vault.c)

`alohadb_vault` stores encrypted named secrets in the PostgreSQL 18-based AlohaDB distribution. Extension version 1.0 uses pgcrypto PGP symmetric encryption, a role-based table policy, and an audit table. Its documented platform is AlohaDB; a standalone community-PostgreSQL package is not established here.

### Core Workflow

Create `pgcrypto` and the extension as a superuser. Configure the superuser-only `alohadb.vault_passphrase` setting through the protected server/session configuration before use; an empty passphrase raises an error. The module registers this setting when loaded and does not request shared preloading.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION alohadb_vault;
SET alohadb.vault_passphrase = 'example-only-change-me';
SELECT alohadb_vault_store('example_key', 'example_value', 'demonstration');
SELECT alohadb_vault_fetch('example_key');
SELECT * FROM alohadb_vault_list();
SELECT alohadb_vault_delete('example_key');
```

This example uses disposable values. Protect real passphrases from SQL history/logging and back them up separately from encrypted data. PL/pgSQL is needed by the installation block that creates the administration role.

### Objects and Permissions

| Object | Purpose |
| --- | --- |
| `alohadb_vault_store(text,text,text)` | Insert or replace a named encrypted secret; description is optional |
| `alohadb_vault_fetch(text)` | Return decrypted plaintext by key |
| `alohadb_vault_delete(text)` | Remove a secret |
| `alohadb_vault_list()` | Return key, description and timestamps |
| `alohadb_vault_rotate_key(text,text)` | Re-encrypt rows using old/new passphrases and return the count |
| `alohadb_vault_secrets`, `alohadb_vault_audit` | Ciphertext and operation records |
| `alohadb_vault_admin` | Role granted access to the backing tables and RLS policy |

### Maintenance Boundary

The C functions execute SQL with the caller's privileges; grant administration membership only to intended operators. RLS and audit records do not isolate data from a superuser or object owner. The implementation uses unqualified table and pgcrypto function names, so control the caller's search path and schema write permissions.

Key rotation rewrites ciphertext but does not change `alohadb.vault_passphrase`; coordinate the configured key with the successful rotation transaction and retain recoverable backups. The extension is not relocatable, and its installation creates a cluster-level role. Do not assume removing a database extension removes that role or constitutes a key-management procedure.
