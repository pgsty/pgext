## Usage

Sources:

- [README](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/README.md)
- [Control](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgvault/pgvault.control)
- [SQL](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgvault/pgvault--0.1.0.sql)
- [Bootstrap and commands](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/scripts/secrets.sh)

`pgvault` is PGEverything's per-database encrypted secret store. Version 0.1.0 is an unreleased source snapshot using `pgsodium` authenticated encryption and `pgcrypto` password hashing. Vault users are application-level principals, separate from PostgreSQL login roles.

### Enable and Use

The documented deployment is the upstream PostgreSQL 16 container. Configure a private `PGE_VAULT_KEY` before starting its pgsodium-enabled server; the development default is insecure. Keep a recoverable copy outside the database: restoring ciphertext without the matching root key is insufficient.

```sh
make secrets-init
make secret-user-create
make secrets-create
make secret-read
```

These interactive helpers prompt for the database and credentials. Initialization creates `pgcrypto`, `age`, `pgsodium` and `pgvault`, registers a derived key, and bootstraps an administrator once. `age` is initialized to satisfy the container's preloaded AGE hook; the pgvault control dependencies are `pgcrypto` and `pgsodium`, and its SQL also uses PL/pgSQL. The extension itself is pure SQL and adds no preload library.

### API and Roles

| Interface | Purpose |
| --- | --- |
| `vault.create_secret`, `vault.update_secret`, `vault.delete_secret` | Write operations for users with write permission |
| `vault.reveal_secret` | Decrypt by UUID for users with read permission |
| `vault.list_secrets` | List metadata without returning plaintext |
| `vault.create_user`, `vault.deactivate_user`, `vault.change_password` | Manage vault principals and credentials |
| `vault.red_alert`, `vault.stand_down` | Disable and re-enable secret access |

Administrators manage users but have no read/write secret permission in this API. Other users receive read, write or both permissions. Every API call authenticates a username and password; these permissions are not per-secret ownership rules. Configuration, users and ciphertext live in the fixed `vault` schema and are included in extension-aware logical dumps.

### Operational Boundary

Credentials are supplied to SQL calls, so protect transport and statement logging. Restrict PostgreSQL roles and ownership separately; application-level checks do not protect against a database superuser who can change objects or permissions. The security-definer search path includes `public`, which must not be writable by untrusted users. Preserve the root key across restart, restore and replication. Upstream supplies no explicit license declaration or broader PostgreSQL compatibility matrix.
