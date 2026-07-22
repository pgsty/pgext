## Usage

Sources:

- [Upstream README](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/README.md)
- [Background-worker implementation](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/src/pg_web.c)
- [HTTP handlers](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/src/pg_web_handler.c)

`pg_web` version `0.1.0` is a 2013 demonstration that runs a tiny embedded HTTP server in a PostgreSQL background worker. Despite the repository description, it exposes no database query interface and should not be deployed as a web administration tool.

### Core Workflow

The worker is registered when the library is loaded. Configure it in `shared_preload_libraries`, choose the postmaster-time port, and restart PostgreSQL.

```conf
shared_preload_libraries = 'pg_web'
pg_web.port = 8080
```

```sql
CREATE EXTENSION pg_web;
```

The versioned SQL script creates no SQL objects. The worker connects to the hard-coded `postgres` database after recovery and serves GET requests for `/`, `/date`, `/count`, and `/ip`. The counter is process-local; the IP endpoint returns the peer address.

### Security and Caveats

The implementation documents no bind-address setting, authentication, authorization, TLS, request limits, or database-backed endpoint. Unknown paths still receive an HTTP 200 response containing an error string. Treat the configured port as remotely reachable unless the host firewall proves otherwise.

This old prototype uses historical background-worker APIs and a bundled networking library. Never expose it to an untrusted network or use it in production. If evaluating it in an isolated lab, verify port binding, process restart behavior, shutdown, malformed requests, and compatibility with the exact PostgreSQL build.
