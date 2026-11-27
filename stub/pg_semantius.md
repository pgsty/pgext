## Usage

Sources:

- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/README.md)
- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/pg_semantius.control)
- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/pg_semantius--0.5.0.sql)
- [Official source](https://github.com/semantius/semantius/blob/v0.5.0/extension/META.json)
- [Official source](https://github.com/semantius/semantius/releases/tag/v0.5.0)

`pg_semantius` 0.5.0 supplies an application data dictionary, RBAC, RLS and a message queue on PostgreSQL 18. This is the versioned PostgreSQL extension channel; the separate CLI deploys the schema over a connection on managed hosts.

### Install in Two Steps

Use a UTF8 database and a superuser with BYPASSRLS. Run each statement separately, without an enclosing transaction and without adding CASCADE:

```sql
CREATE EXTENSION pg_semantius;
CALL semantius.migrate();
SELECT semantius.version();
SELECT * FROM semantius.status();
SELECT * FROM semantius.pending();
```

The first statement creates the cluster roles, installer schema and functions. `semantius.migrate()` installs the application objects as ordinary database objects and commits each migration file separately. It creates `pgcrypto` in `public`; installation is refused if that extension already occupies another schema. A failed later file leaves earlier files committed, and the next call resumes.

### Identity and Application Access

Installation creates `authenticated`, `semantius_user`, `semantius_authenticator` and `semantius_owner` as passwordless NOLOGIN roles. Keep the owner without LOGIN or members. Existing roles with incompatible attributes are refused. Configure application login separately and keep ordinary requests under RLS.

The session model trusts the application tier to authenticate callers and set `request.jwt.claims` within the request transaction. Direct access to a role allowed to forge these settings bypasses that trust boundary. Set the `jwt_aud` row in `public._settings` when audience checking is required; without it the audience is not checked. Inspect `semantius.status()` and role grants after deployment.

### Upgrades and Recovery

After installing newer extension files, run `ALTER EXTENSION pg_semantius UPDATE`, then call the migration procedure outside a transaction. The extension update replaces the installer; the migration call changes the application schema. `public._versions` records the migration checksums. Newer or incompatible run-once histories can cause refusal before any file is applied; do not bypass that protection by editing the ledger.

Back up the ordinary application objects and roles, not just extension membership. Restore into an empty UTF8 database with matching or newer extension files and compatible owner roles. Restore login credentials separately on a new cluster. Preserve database settings `timezone = 'UTC'` and `intervalstyle = 'iso_8601'`; ordinary dumps without database creation do not carry them automatically.

`DROP EXTENSION pg_semantius` removes the installer, while application data and cluster roles remain. There is no shared-preload or restart requirement for the SQL-only installer. Review the upstream migration and ownership boundaries before using destructive full-uninstall instructions.
