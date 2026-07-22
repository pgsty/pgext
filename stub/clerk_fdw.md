## Usage

Sources:

- [Pinned official README](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/README.md)
- [Pinned FDW implementation](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/src/lib.rs)

`clerk_fdw` is a read-only foreign data wrapper for the Clerk user-management API. It exposes users, organizations, and organization memberships as foreign tables; rows are fetched from Clerk when a scan runs rather than stored in PostgreSQL.

### Core Workflow

Version 0.3.3 installs the handler and validator functions. Define the wrapper, a server containing the Clerk secret key, and one or more foreign tables:

```sql
CREATE EXTENSION clerk_fdw;

CREATE FOREIGN DATA WRAPPER clerk_wrapper
  HANDLER clerk_fdw_handler
  VALIDATOR clerk_fdw_validator;

CREATE SERVER my_clerk_server
  FOREIGN DATA WRAPPER clerk_wrapper
  OPTIONS (api_key '<clerk-secret-key>');

CREATE FOREIGN TABLE clerk_users (
  user_id text,
  first_name text,
  last_name text,
  email text,
  gender text,
  created_at bigint,
  updated_at bigint,
  last_sign_in_at bigint,
  phone_numbers bigint,
  username text,
  attrs jsonb
)
SERVER my_clerk_server
OPTIONS (object 'users');

SELECT user_id, email, attrs FROM clerk_users;
```

If `api_key` is absent, the implementation reads `CLERK_API_KEY` from the PostgreSQL server process environment. Protect the foreign server definition and catalog access because server options can contain credentials.

### Supported Foreign Objects

- `object 'users'`: user identifiers, names, primary email, timestamps, username, and the complete response in `attrs`.
- `object 'organizations'`: organization identifier, name, slug, timestamps, creator, and `attrs`.
- `object 'organization_memberships'`: `user_id`, `organization_id`, and `role`.

Column names and types must match the mappings implemented by the FDW. Include an `attrs jsonb` column when the unflattened API object is needed.

### Query and Performance Boundaries

The FDW ignores PostgreSQL quals, sort keys, and limits during the remote fetch. A predicate such as `WHERE organization_id = 'org_123'` is evaluated after the scan, not pushed down to Clerk.

Membership scans first enumerate every organization and then request its memberships. The implementation retries rate-limited membership calls with exponential backoff and can take a long time; materialize results into a local table for repeated analytics. The implementation exposes scan callbacks only, so do not expect `INSERT`, `UPDATE`, or `DELETE` support.

Version 0.3.3 is built for PostgreSQL 14–17, is non-relocatable, and requires superuser privileges to create. It does not require `shared_preload_libraries`.
