## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/agentic-db/blob/cd818fea7f480ef3ff6099f736d66848a92b907a/extensions/@pgpm/jwt-claims/README.md)
- [Official extension control file (pgpm-jwt-claims.control)](https://github.com/constructive-io/agentic-db/blob/cd818fea7f480ef3ff6099f736d66848a92b907a/extensions/@pgpm/jwt-claims/pgpm-jwt-claims.control)
- [Official extension SQL (pgpm-jwt-claims--0.15.5.sql)](https://github.com/constructive-io/agentic-db/blob/cd818fea7f480ef3ff6099f736d66848a92b907a/extensions/@pgpm/jwt-claims/sql/pgpm-jwt-claims--0.15.5.sql)

`pgpm-jwt-claims` — @pgpm/jwt-claims provides PostgreSQL functions for extracting and working with JWT (JSON Web Token) claims stored in PostgreSQL session variables. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-jwt-claims";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ctx.ip_address()` is an extension function and returns `inet`.
- `ctx.is_security_definer()` is an extension function.
- `ctx.origin()` is an extension function and returns `origin`.
- `ctx.security_definer()` is an extension function and returns `text`.
- `ctx.uagent()` is an extension function and returns `text`.
- `ctx.uid()` is an extension function and returns `uuid`.
- `jwt_private.current_database_id()` is an extension function and returns `uuid`.
- `jwt_private.current_session_id()` is an extension function and returns `uuid`.
- `jwt_private.current_token_id()` is an extension function and returns `uuid`.
- `jwt_public.current_ip_address()` is an extension function and returns `inet`.
- `jwt_public.current_origin()` is an extension function and returns `origin`.
- `jwt_public.current_user_agent()` is an extension function and returns `text`.
- `jwt_public.current_user_id()` is an extension function and returns `uuid`.
- `ctx` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgpm-types`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
