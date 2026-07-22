## Usage

Sources:

- [Official README](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/README)
- [Official extension SQL](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/pg_stat_statements2--1.2.sql)
- [Official C implementation](https://github.com/uptimejp/pg_stat_statements2/blob/db2e94c49c9be23137f6a8925616535950f2ae89/pg_stat_statements2.c)

`pg_stat_statements2` is a PostgreSQL 9.4.4-era fork of the standard statement-statistics extension, modified to reuse query fingerprints from `sql_firewall`. It exists for that historical integration; it is not a drop-in modern replacement for PostgreSQL's bundled module.

### Prerequisite and Startup

Both libraries must be preloaded, in this order, and changing the setting requires a server restart:

```conf
shared_preload_libraries = 'sql_firewall,pg_stat_statements2'
```

After restart, install the SQL objects. A dedicated schema reduces name collisions with other extensions:

```sql
CREATE SCHEMA firewall_stats;
CREATE EXTENSION pg_stat_statements2 WITH SCHEMA firewall_stats;

SELECT query, calls, total_time, rows
FROM firewall_stats.pg_stat_statements
ORDER BY total_time DESC
LIMIT 20;
```

### SQL Surface

- `pg_stat_statements(showtext boolean)` returns the PostgreSQL 9.4-style counters, including calls, total time, rows, block activity, and block I/O time.
- `pg_stat_statements` is a view over that function and is granted to `PUBLIC` by the installation script.
- `pg_stat_statements_reset()` clears the shared statistics and has its default `PUBLIC` privilege revoked.

The view exposes captured query text to any role that can select it. Review and, where required, revoke access to `firewall_stats.pg_stat_statements`; treat the reset function as an administrative and destructive operation.

### Compatibility Boundaries

The extension deliberately removes its own `JumbleQuery` implementation and calls the symbol exported by `sql_firewall`. Therefore library load order and an ABI-compatible firewall build are mandatory; a missing or mismatched symbol prevents startup.

Its SQL object names overlap the bundled statement-statistics extension. Schema separation limits SQL-name collisions, but it does not establish that both shared libraries, hooks, GUCs, or ABIs can coexist. Test the complete preload combination on an isolated server.

The output schema predates planning counters and other fields found in current PostgreSQL releases. Upstream documents the PostgreSQL 9.4.4 and historical `sql_firewall` combination only, and the reviewed repository was last changed in 2015. Do not load it on a modern server without a deliberate port, source review, and workload testing.
