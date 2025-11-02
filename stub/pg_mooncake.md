

--------

## Usage

[`pg_mooncake`](https://github.com/Mooncake-Labs/pg_mooncake) 0.2.0 (unpublished yet) is rewritten in Rust and designed as a sub-extension of `pg_duckdb`.

- pg_mooncake docs: https://docs.mooncake.dev/



### Quick Setup

Install pg_duckdb and pg_mooncake with pig:

```bash
pig repo add pigsty -u
pig install pg_duckdb pg_mooncake
```

Edit postgresql.conf, then restart to take effect

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
duckdb.allow_community_extensions = true
wal_level = logical
```



### Hello Worlds

- [Tutorial](https://docs.mooncake.dev/pg/get-started/Hello-world)

```sql
-- create the extension alone with pg_duckdb
CREATE EXTENSION pg_mooncake CASCADE;

-- Next, create a regular Postgres table trades:
CREATE TABLE trades(
  id bigint PRIMARY KEY,
  symbol text,
  time timestamp,
  price real
);

-- Then, create a columnstore mirror trades_iceberg that stays in sync with trades:
CALL mooncake.create_table('trades_iceberg', 'trades');

-- Now, insert some data into trades:
INSERT INTO trades VALUES
    (1,  'AMD', '2024-06-05 10:00:00', 119),
    (2, 'AMZN', '2024-06-05 10:05:00', 207),
    (3, 'AAPL', '2024-06-05 10:10:00', 203),
    (4, 'AMZN', '2024-06-05 10:15:00', 210);

-- Finally, query it with duckdb
EXPLAIN
    SELECT avg(price) FROM trades_iceberg WHERE symbol = 'AMZN';
```

You will see the DuckDBScan in the execution plan:

```bash
                         QUERY PLAN
------------------------------------------------------------
 Custom Scan (DuckDBScan)  (cost=0.00..0.00 rows=0 width=0)
   DuckDB Execution Plan:

 ┌───────────────────────────┐
 │    UNGROUPED_AGGREGATE    │
 │    ────────────────────   │
 │    Aggregates: avg(#0)    │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │         PROJECTION        │
 │    ────────────────────   │
 │   CAST(price AS DOUBLE)   │
 │                           │
 │          ~0 rows          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │       MOONCAKE_SCAN       │
 │    ────────────────────   │
 │   Table: trades_iceberg   │
 │     Projections: price    │
 │                           │
 │          Filters:         │
 │       symbol='AMZN'       │
 │                           │
 │          ~0 rows          │
 └───────────────────────────┘
```
