## Usage

Sources:

- [Official EDB Stat Monitor documentation](https://www.enterprisedb.com/docs/pg_extensions/edb_stat_monitor/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/edb_stat_monitor/configuring/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/edb_stat_monitor/using/)

`edb_stat_monitor` records bucketed query execution statistics for PostgreSQL and EDB Postgres, extending the pg_stat_monitor model with EDB-supported analysis features.

### Enablement

Install EDB Stat Monitor 2.1.0, preload it, restart, and create the extension:

```ini
shared_preload_libraries = 'edb_stat_monitor'
```

```sql
CREATE EXTENSION edb_stat_monitor;
```

The extension must be created in each database whose statistics you want to query.

### Query Collected Statistics

```sql
SELECT application_name,
       userid::regrole AS user_name,
       datname,
       calls,
       query
FROM edb_stat_monitor
ORDER BY calls DESC;
```

Use the documented bucket, plan, histogram, client, and error fields to separate workload intervals instead of aggregating unlike executions blindly.

### Operational Boundaries

Collection uses shared memory and adds per-statement accounting. Size retention, buckets, query text, plan capture, and histogram settings before enabling it broadly, and restrict access because normalized text, client addresses, usernames, and plans can expose sensitive workload details. Extension upgrades do not automatically rewrite application queries that depend on view columns; review release notes first.

