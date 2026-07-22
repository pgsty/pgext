## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/README.md)
- [Extension schema and bootstrap data](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/pg_pgbouncer.sql)
- [Process manager implementation](https://github.com/pgbouncer/pg_pgbouncer/blob/e71619b01d391316c98eb8456e5b0b37c76d96d9/src/child.rs)

`pg_pgbouncer` is an alpha-quality pgrx extension that runs a PostgreSQL background worker to generate configuration and supervise external `pgbouncer` processes. Upstream explicitly does not recommend it for production and warns that its API may change substantially.

```conf
shared_preload_libraries = 'pg_pgbouncer'
pg_pgbouncer.database = 'pg_pgbouncer'
```

Restart PostgreSQL, create `pg_pgbouncer` in that exact database, and ensure the operating-system account can find and execute the `pgbouncer` binary. Configuration is stored in `pgbouncer.groups`, `pgbouncer.databases`, `pgbouncer.users`, `pgbouncer.peers`, `pgbouncer.auth`, `pgbouncer.hba`, and `pgbouncer.settings`; the manager writes INI files and starts or signals child processes.

The bootstrap SQL copies password values from `pg_shadow`, stores credentials in ordinary tables, enables `trust` authentication, and includes example users. Treat it as unsafe sample state: isolate a disposable cluster, revoke schema/table access from `PUBLIC`, replace every default, protect generated files, and audit process ownership and signal handling.

The worker is cluster-level and connects through SPI as a superuser. A database outage, malformed table row, missing executable, stale PID/socket, or rapidly crashing child can disrupt supervision. For production use, manage PgBouncer with an operating-system service manager and its supported configuration workflow.
