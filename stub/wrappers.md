## Usage

Sources:

- [Wrappers v0.6.4 README](https://github.com/supabase/wrappers/blob/v0.6.4/README.md)
- [Official FDW documentation](https://fdw.dev/)
- [v0.6.4 release](https://github.com/supabase/wrappers/releases/tag/v0.6.4)
- [MongoDB FDW documentation](https://fdw.dev/catalog/mongodb/)
- [Security guidance](https://fdw.dev/guides/security/)

`wrappers` is both a Rust framework for writing PostgreSQL foreign data wrappers and a packaged collection of Supabase-maintained FDWs. A single extension installs many wrapper implementations, then each foreign server chooses the specific wrapper type it needs.

```sql
CREATE EXTENSION wrappers;
```

### Typical Workflow

Create a server for one wrapper, then expose remote data through foreign tables:

```sql
CREATE SERVER stripe_server
  FOREIGN DATA WRAPPER stripe_wrapper
  OPTIONS (
    api_key_id 'stripe_api_key',
    api_url 'https://api.stripe.com/v1/'
  );

CREATE FOREIGN TABLE stripe_customers (
  id text,
  email text,
  name text,
  description text,
  created timestamp,
  attrs jsonb
)
  SERVER stripe_server
  OPTIONS (
    object 'customers',
    rowid_column 'id'
  );
```

### What It Covers

Upstream ships wrappers for databases and services such as BigQuery, ClickHouse, DuckDB, DynamoDB, MySQL/Doris, Redis, S3, S3 Vectors, Stripe, Snowflake, Slack, Notion, OpenAPI, Infura, and many others. Read and write support varies by wrapper, but pushdown for `WHERE`, `ORDER BY`, and `LIMIT` is a core framework feature.

### Version 0.6.2

The `v0.6.2` release keeps the same extension model and adds:

- a MongoDB FDW with read and write support
- session-variable credentials for per-request authentication in WASM wrappers
- RFC 8288 `Link` header pagination for the OpenAPI FDW
- runtime, dependency, and wrapper-specific fixes documented in the release notes

Wrapper-specific pages remain the authority for server options, foreign-table columns, pushdown, and write support.

### Caveats

- Wrapper-specific options, supported objects, and write support differ widely; check the official catalog page for the exact FDW you use.
- The docs warn that logical restores can fail when materialized views depend on foreign tables, so avoid that pattern or rely on physical backups.
- Foreign tables do not provide a security boundary by themselves. Keep them in private schemas, grant access deliberately, use the least-privileged remote credentials available, and apply Row Level Security to local tables that expose or cache remote data.
- Keep API keys and tokens in the supported secret store or per-request credential mechanism instead of embedding them in SQL checked into source control.

### 0.6.4 Update

The 0.6.4 release updates embedded DuckDB to 1.5.6 and repairs the Shopify/Paddle WASM interface definitions. The upstream Cargo manifest exposes PostgreSQL 13–18 targets and pgrx 0.19.2; packaged build choices can differ. The available wrappers depend on enabled build features. Upgrade the installed extension with matching library/SQL files and review the specific wrapper documentation before changing server options.
