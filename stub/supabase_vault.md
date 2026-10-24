## Usage

Sources:

- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/README.md)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/supabase_vault.control.in)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/Makefile)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/sql/supabase_vault--0.3.1--0.3.2.sql)

`supabase_vault` 0.3.2 stores encrypted secrets in PostgreSQL and exposes decrypted values through a privileged view. Its fixed schema is `vault`; the current release no longer requires `pgsodium`. Upstream tests PostgreSQL 14–19. Install as a superuser; no preload or restart is required.

### Core Workflow

The following values are demonstration strings, not real credentials:

```sql
CREATE EXTENSION supabase_vault;
SELECT vault.create_secret('example-secret', 'api_token', 'Demo only');
SELECT name, decrypted_secret FROM vault.decrypted_secrets WHERE name = 'api_token';
SELECT vault.update_secret(id, 'replacement-example', 'api_token', 'Rotated demo')
FROM vault.secrets WHERE name = 'api_token';
```

`vault.create_secret` returns the secret UUID and accepts optional name and description. `vault.update_secret` updates a selected UUID. `vault.secrets` holds encrypted values and metadata; `vault.decrypted_secrets` exposes plaintext to roles allowed to read it. Restrict both helper and view privileges deliberately.

### Security and Upgrade

Version 0.3.2 moves `pg_temp` to the end of the search path used by secret creation/update helpers. After installing its files, run `ALTER EXTENSION supabase_vault UPDATE TO '0.3.2'` in each database.

Encryption at rest does not prevent authorized queries from returning plaintext. Secret literals and returned values can appear in database logs, client histories or application telemetry; use the documented secret workflow and configure those systems appropriately. Protect the encryption-key environment and verify backup/restore behavior. Do not treat access to the decrypted view as safe merely because the stored table is encrypted.
