## Usage

Sources:

- [Pinned official README](https://github.com/nuko-yokohama/pg_sulog/blob/e07a87f76261ccfd884da42bff84419b395cba81/README.md)
- [Pinned hook implementation](https://github.com/nuko-yokohama/pg_sulog/blob/e07a87f76261ccfd884da42bff84419b395cba81/pg_sulog.c)

`pg_sulog` is a preload-only hook module that writes superuser statements to the PostgreSQL server log and can attempt to block them. It creates no SQL objects, so there is no application-facing `CREATE EXTENSION` workflow; its behavior begins when the shared library is preloaded at postmaster startup.

### Configuration

Install the matching shared library, add it to `postgresql.conf`, choose a mode, and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_sulog'
pg_sulog.mode = 'LOGGING'
```

The documented modes are:

- `LOGGING`: log every statement issued by a superuser and allow it to run; this is the default.
- `MAINTENANCE`: allow and log `VACUUM`, `REINDEX`, `ANALYZE`, and `CLUSTER`, while treating other superuser statements as blocked.
- `BLOCK`: treat all superuser statements as blocked.

Messages are emitted as PostgreSQL warnings containing a timestamp, `logging` or `blocked`, the login role, and the original query text. Account for sensitive literals in server-log retention and access controls.

### Operational Boundaries

The implementation hooks executor and utility processing and builds an in-memory list of superuser role names. The list is limited to 64 entries. It identifies the login role from the connection and does not provide a SQL audit table, structured event schema, per-role policy, redaction, or durable tamper protection.

This abandoned module was validated upstream only on CentOS 7 and PostgreSQL 9.3–9.5, with 9.6 development noted. PostgreSQL hook signatures have changed since then. More importantly, the pinned blocking code contains inconsistent mode comparisons and commented-out execution paths; do not rely on `MAINTENANCE` or `BLOCK` as a security control without source review and destructive isolation tests on the exact server build.

Use `LOGGING` only after verifying compatibility and log volume. A preload failure can prevent PostgreSQL from starting, so retain a way to edit the configuration offline.
