## Usage

Sources:

- [PGXN 0.0.1 README](https://pgxn.org/dist/pg_dctp/0.0.1/README.html)
- [PGXS Makefile](https://api.pgxn.org/src/pg_dctp/pg_dctp-0.0.1/Makefile)
- [C module source](https://api.pgxn.org/src/pg_dctp/pg_dctp-0.0.1/pg_dctp.c)
- [PostgreSQL license](https://api.pgxn.org/src/pg_dctp/pg_dctp-0.0.1/LICENSE)

`pg_dctp` is a headless PostgreSQL C module that rejects clear-text password literals in `CREATE ROLE`, `CREATE USER`, `ALTER ROLE`, and `ALTER USER`. It installs no SQL objects and has no control file, so activation is entirely a server preload decision.

### Enable the Module

```conf
shared_preload_libraries = 'pg_dctp'
```

Restart PostgreSQL after changing `shared_preload_libraries`. Once loaded, a command that embeds a clear-text password is rejected. Use `psql`'s `\password`, `createuser -P`, or another client that sends a precomputed SCRAM verifier instead of placing a password literal in SQL.

### Scope and Boundaries

The module addresses one logging hazard: a clear-text password embedded in a role DDL statement can appear in statement logs. While reporting the rejection, `pg_dctp` temporarily changes message handling to avoid echoing that password in the error log.

It does not enforce password length, complexity, reuse, expiration, authentication method, or `pg_hba.conf` policy. It also cannot protect secrets sent through unrelated SQL, client logs, shell history, monitoring, or network capture. Test role provisioning, password rotation, backup/restore tooling, and automation before enabling it cluster-wide.

Upstream explicitly calls `pg_dctp` a module rather than a SQL extension and reports validation on PostgreSQL 14 through 18. Version 0.0.1 is the renamed successor of the short-lived `disable_set_password` distribution; use the current module name and do not preload both names.
