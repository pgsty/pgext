## Usage

Sources:

- [Official extension control file](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/pg_graphql_server.control)
- [Official background-worker implementation](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/src/lib.rs)
- [Official HTTP-server implementation](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/src/server.rs)

`pg_graphql_server` 0.0.1 is an experimental pgrx background worker that starts HTTP endpoints inside the PostgreSQL postmaster and delegates GraphQL query strings to the separate `pg_graphql` extension. The pinned source has no authentication or TLS layer and should not be exposed as a production API without substantial hardening.

### Core Workflow

Install the library, add `pg_graphql_server` to `shared_preload_libraries`, and restart PostgreSQL. The worker connects to the `postgres` database and creates its administration objects there. Install both extensions and add a listener configuration:

```sql
SHOW shared_preload_libraries;

CREATE EXTENSION pg_graphql;
CREATE EXTENSION pg_graphql_server;

INSERT INTO http_server.servers (
  listen_port,
  postgres_user,
  postgres_pass,
  database_name
)
VALUES (8080, 'graphql_app', 'change-this-password', 'appdb');
```

After commit, the table trigger requests a configuration reload. The worker starts a listener and connects back to the local PostgreSQL port using the stored credentials. It exposes `GET /health` and both `GET /graphql` and `POST /graphql`; GraphQL requests are evaluated with `graphql.resolve($1)` in the configured target database.

### Important Objects

- `http_server.servers` stores one listener per `listen_port`, along with PostgreSQL username, password, and target database.
- `http_server.on_servers_changed()` and its statement trigger call `pg_reload_conf()` after configuration changes.
- `/health` reports connection-pool size and idle connections.
- `/graphql` accepts a GraphQL request and returns the JSON value produced by `graphql.resolve($1)`.

### Preload and Compatibility

The worker is registered by `_PG_init`, starts after recovery finishes, and requires SPI access, so preload and a server restart are mandatory. The Cargo manifest has build features for PostgreSQL 14, 15, and 16. Every configured target database must have `pg_graphql` available and installed or its listener fails and is retried.

### Security and Operational Notes

The pinned implementation binds each HTTP server to all IPv4 interfaces, provides no authentication, authorization, TLS, origin filtering, request-size limit, or query allowlist, and stores database passwords as plain table values. It also bootstraps only from the database named `postgres` and connects back over `::1`. Restrict network reachability, use a narrowly privileged database role, protect the administration table, and place a hardened proxy in front if evaluating the code. Treat 0.0.1 as a prototype and audit the exact source before any non-lab deployment.
