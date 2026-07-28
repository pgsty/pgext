## Usage

Sources:

- [Official upstream README](https://github.com/masahikosawada/pg_promoter/blob/9e70f65508e2bfe8c1e30631dc709353127c9ce0/README.md)
- [Official extension control file (pg_promoter.control)](https://github.com/masahikosawada/pg_promoter/blob/9e70f65508e2bfe8c1e30631dc709353127c9ce0/pg_promoter.control)
- [Official implementation source](https://github.com/masahikosawada/pg_promoter/blob/9e70f65508e2bfe8c1e30631dc709353127c9ce0/pg_promoter.c)

`pg_promoter` is an early background-worker prototype that runs only on a standby, polls the primary, and promotes locally after a configured number of failed probes. It does not provide a complete consensus, quorum, or fencing system.

### Core Workflow

Install the library on a disposable standby, configure it for preloading, and restart:

```ini
shared_preload_libraries = 'pg_promoter'
pg_promoter.keepalives_time = 5
pg_promoter.keepalives_count = 3
pg_promoter.primary_conninfo = 'host=primary port=5432 dbname=postgres'
```

The source-code spellings are `keepalives_time` and `keepalives_count`; use those plural forms even though the README examples use singular names. The worker promotes after consecutive connection failures reach the configured count.

### Important Settings

- `pg_promoter.primary_conninfo` is the libpq connection string used to probe the primary.
- `pg_promoter.keepalives_time` is the polling interval in seconds.
- `pg_promoter.keepalives_count` is the failed-probe threshold before promotion.

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `1.0`.
- The control file marks the extension as relocatable.
- This is a preload-only worker; the reviewed repository does not provide extension SQL.
- Network failure is not proof that the old primary is down. Without external fencing this design can create split brain and data loss.
- The code uses an old trigger-file promotion path and PostgreSQL internals. Treat it as historical prototype source and test only against the server generation it was written for.
