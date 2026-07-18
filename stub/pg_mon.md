## Usage

Sources:

- [Upstream README](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/README.md)
- [Extension SQL](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/pg_mon--1.0.sql)
- [Extension control file](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/pg_mon.control)
- [Upstream CI matrix](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/.github/workflows/tests.yaml)

`pg_mon` records query execution-time and row-count histograms plus scan and join information in shared memory. It must be preloaded:

```conf
shared_preload_libraries = 'pg_mon'
```

Restart PostgreSQL, then create the extension and query its view. Restrict the reset function if ordinary users must not erase monitoring state:

```sql
CREATE EXTENSION pg_mon;
REVOKE EXECUTE ON FUNCTION pg_mon_reset() FROM PUBLIC;

SELECT queryid, total_time, actual_rows,
       seq_scans, index_scans, hash_join_count,
       hist_time_ubounds, hist_time_freq
FROM pg_mon
ORDER BY total_time DESC;
```

### Operation

Version `1.0` has an upstream CI matrix for PostgreSQL 11 through 17. `pg_mon.plan_info_immediate` makes plan information visible after planning but can add locking overhead; `pg_mon.plan_info_disable` suppresses plan details. Histogram boundaries are fixed, and the collected state is memory-resident rather than a durable history. The install SQL grants `SELECT` on the view to `PUBLIC`, so review visibility requirements on shared systems.
