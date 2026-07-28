## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jwt-claims/README.md)
- [Official extension control file (launchql-jwt-claims.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/jwt-claims/launchql-jwt-claims.control)

`launchql-jwt-claims` — PostgreSQL extension for accessing JWT claims in database functions. This extension provides schemas and functions to access JWT token claims from within PostgreSQL, making it easy to implement authentication and authorization logic directly in your database. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-jwt-claims";

-- Access user ID from JWT claims
SELECT ctx.user_id();

-- Access IP address from JWT claims
SELECT ctx.ip_address();

-- Access user agent from JWT claims
SELECT ctx.user_agent();

-- Access origin from JWT claims
SELECT ctx.origin();

-- Access database ID from JWT claims
SELECT jwt_private.current_database_id();

-- Access token ID from JWT claims
SELECT jwt_private.current_token_id();

-- Access IP address (public schema)
SELECT jwt_public.current_ip_address();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.4.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `uuid-ossp`, `launchql-ext-types`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
