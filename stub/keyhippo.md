## Usage

Sources:

- [Project README](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/README.md)
- [Extension-specific README](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/extension/README.md)
- [Extension control file](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/extension/keyhippo.control)
- [Version 1.2.5 installation SQL](https://github.com/integrated-reasoning/KeyHippo/blob/f335bc455d5248fc3b597c460bada32747c83366/extension/keyhippo--1.2.5.sql)

`keyhippo` 1.2.5 adds API-key authentication and role-based authorization to a Supabase/PostgREST database while preserving row-level security. It stores only a hash of each secret, supports key scopes, expiry, revocation, rotation, and claims, and provides groups, roles, and built-in permission types.

### Issue and use a key

```sql
CREATE EXTENSION keyhippo;

SELECT * FROM keyhippo.create_api_key('Primary API Key', 'default');
SELECT keyhippo.key_data();
SELECT keyhippo.authorize('manage_api_keys');
```

The main APIs include `keyhippo.create_api_key(text,text)`, `keyhippo.verify_api_key(text)`, `keyhippo.revoke_api_key(uuid)`, `keyhippo.rotate_api_key(uuid)`, and `keyhippo.authorize(keyhippo.app_permission)`. API keys arrive through the PostgREST `x-api-key` request header; the creation call returns the plaintext key, so the caller must capture it at issuance.

### Supabase dependencies and installation effects

This is not a generic PostgreSQL authentication extension. Its SQL expects `auth.users`, `auth.uid()`, and the `anon`, `authenticated`, `authenticator`, and `service_role` roles. Installation creates `pgcrypto`, `pg_net`, and `pg_cron` if absent, four fixed schemas, RLS policies, grants, an `auth.users` trigger, and default RBAC records. It also changes `authenticator`'s `pgrst.db_pre_request` setting and notifies PostgREST to reload.

Review the installation SQL before applying it. Initial setup enables an installation notification and sends an HTTP request through `pg_net` to `https://app.keyhippo.com/api/ingest`; audit HTTP logging is disabled by default but the endpoint remains configured. The control file says relocatable even though the SQL creates fixed schemas. Treat all security-definer functions, grants, impersonation procedures, outbound HTTP behavior, and PostgREST role changes as security-sensitive migration content.
