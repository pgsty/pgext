## Usage

Sources:

- [Official extension control file](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/pg_graphql_server.control)
- [Official Rust package manifest](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/Cargo.toml)

`pg_graphql_server` — Runs an HTTP GraphQL endpoint from a PostgreSQL background worker and delegates queries to pg_graphql.

The reviewed catalog snapshot records version `0.0.1`, kind `preload`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `pg_graphql`, `plpgsql`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_graphql_server";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_graphql_server';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
