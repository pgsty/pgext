


## Usage

> [wrappers: Foreign data wrappers developed by Supabase](https://github.com/supabase/wrappers)

Supabase Wrappers is a framework for building PostgreSQL Foreign Data Wrappers (FDW), with 30+ pre-built connectors for cloud services, databases, and APIs. It supports WHERE, ORDER BY, and LIMIT pushdown, with some wrappers also supporting data modification (INSERT/UPDATE/DELETE).

### Available Wrappers

| Category | Wrappers |
|----------|----------|
| **Databases** | ClickHouse, BigQuery, Snowflake, DuckDB, SQL Server, Redis, PostgreSQL |
| **Storage** | AWS S3, Cloudflare D1, Apache Iceberg |
| **SaaS/APIs** | Stripe, Firebase, Airtable, Auth0, Notion, Slack, HubSpot, Shopify |
| **Auth** | AWS Cognito, Clerk, Auth0 |
| **Other** | OpenAPI, Logflare, Calendly, Cal.com, Paddle, Orb, Infura, Gravatar |

### Example (Stripe)

```sql
CREATE EXTENSION wrappers;

CREATE SERVER stripe_server
  FOREIGN DATA WRAPPER stripe_wrapper
  OPTIONS (
    api_key_id '<key_ID>',
    api_url 'https://api.stripe.com/v1/',
    api_version '2024-06-20'
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

SELECT id, email, name FROM stripe_customers WHERE email LIKE '%@example.com';
```

The `rowid_column` option is required for INSERT/UPDATE/DELETE support. The `attrs` column provides access to additional metadata as JSON.

Each wrapper uses its own `FOREIGN DATA WRAPPER` name (e.g., `stripe_wrapper`, `firebase_wrapper`, `clickhouse_wrapper`), but they are all installed via the single `wrappers` extension. Refer to the [documentation](https://fdw.dev) for wrapper-specific options and supported objects.
