## Usage

Sources:

- [Official upstream documentation](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/README.md)
- [Required PostgreSQL core patch](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/patches/enable-hooks.patch)
- [Official extension SQL](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/pg_stat_usage--1.0.sql)

`pg_stat_usage` 1.0 is an abandoned prototype that correlates stored-function calls with relation and index activity inside one PostgreSQL backend. It cannot run on a stock server: the extension requires the included, old PostgreSQL core patch adding private statistics hooks and relation-close callbacks.

### Patched-Server Workflow

Only after building a PostgreSQL server with the exact matching patch, install the module at the path expected by local preload and create its SQL objects:

```conf
local_preload_libraries = 'pg_stat_usage'
```

```sql
CREATE EXTENSION pg_stat_usage;

SELECT * FROM pg_stat_usage;
SELECT pg_stat_usage_reset();
```

`local_preload_libraries` affects newly started sessions. Confirm the shared-library location: the control file references `$libdir/plugins/pg_stat_usage`, while the README and default PGXS build use a shorter load name.

### Reported Columns

The `pg_stat_usage` view wraps `pg_stat_usage()` and reports object OID, parent function OID, object type, schema and name, call and scan counts, total and self time, tuple fetch/return/change counters, and block fetch/hit counters. The parent OID can be used to approximate which function was active when a table or index was touched. `pg_stat_usage_reset()` clears the current backend's collected counters.

### Prototype Limitations

Collection state is process-local; there is no shared or external collector aggregating all backends. The source itself says transaction commit/rollback attribution is only occasionally correct and always assumes commit for some relation counters. Timing uses wall-clock measurements, and relation context can be approximate. Although the SQL view is granted to `PUBLIC`, it only exposes the calling backend's state.

The required patch targets a superseded statistics subsystem and cannot be assumed to apply to modern PostgreSQL. Do not forward-port it casually into a production server. If researching it, isolate the patched build, reconcile the module install path, restrict preload and reset privileges, and test nested functions, errors, rollback, parallel queries, and connection churn. Use supported observability extensions for production call and object statistics.
