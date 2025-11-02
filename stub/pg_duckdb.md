
--------

## Usage

- [pg_duckdb docs](https://github.com/duckdb/pg_duckdb/tree/main/docs)

| Topic                                                                                                  | Description                                                |
|--------------------------------------------------------------------------------------------------------|------------------------------------------------------------|
| [**Functions**](https://github.com/duckdb/pg_duckdb/blob/main/docs/functions.md)                       | Complete reference for all available functions             |
| [**Syntax Guide & Gotchas**](https://github.com/duckdb/pg_duckdb/blob/main/docs/gotchas_and_syntax.md) | Quick reference for common SQL patterns and things to know |
| [**Types**](https://github.com/duckdb/pg_duckdb/blob/main/docs/types.md)                               | Supported data types and type mappings                     |
| [**Extensions**](https://github.com/duckdb/pg_duckdb/blob/main/docs/extensions.md)                     | DuckDB extension installation and usage                    |
| [**Settings**](https://github.com/duckdb/pg_duckdb/blob/main/docs/settings.md)                         | Configuration options and parameters                       |
| [**Transactions**](https://github.com/duckdb/pg_duckdb/blob/main/docs/transactions.md)                 | Transaction behavior and limitations                       |



### Quick Setup

Install pg_duckdb with pig:

```bash
pig repo add pigsty -u
pig install pg_duckdb
```

Edit `postgresql.conf`, then restart to take effect

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
duckdb.allow_community_extensions = true
wal_level = logical
```


### Accelerate Query

You can use DuckDB to query existing PostgreSQL table without modifying them:

```sql
-- pgbench -is 1000  # init some test workloads with pgbench
CREATE EXTENSION pg_duckdb;

-- default behavior, common postgres engine
SET duckdb.force_execution = false;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;

-- now the query goes to pg_duckdb
SET duckdb.force_execution = true;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;
```

The result would be 8s -> 4s on 4c VM on local laptop) :

```
postgres@el9:5432/postgres=# SET duckdb.force_execution = true;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;
SET
Time: 0.206 ms
                                              QUERY PLAN
------------------------------------------------------------------------------------------------------
 Custom Scan (DuckDBScan)  (cost=0.00..0.00 rows=0 width=0) (actual time=0.001..0.001 rows=0 loops=1)
   DuckDB Execution Plan:

 ┌─────────────────────────────────────┐
 │┌───────────────────────────────────┐│
 ││    Query Profiling Information    ││
 │└───────────────────────────────────┘│
 └─────────────────────────────────────┘
 EXPLAIN ANALYZE  SELECT count(*) AS count FROM pgduckdb.public.pgbench_accounts
 ┌────────────────────────────────────────────────┐
 │┌──────────────────────────────────────────────┐│
 ││               Total Time: 3.89s              ││
 │└──────────────────────────────────────────────┘│
 └────────────────────────────────────────────────┘
 ┌───────────────────────────┐
 │           QUERY           │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │      EXPLAIN_ANALYZE      │
 │    ────────────────────   │
 │           0 rows          │
 │          (0.00s)          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │    UNGROUPED_AGGREGATE    │
 │    ────────────────────   │
 │        Aggregates:        │
 │        count_star()       │
 │                           │
 │           1 row           │
 │          (0.00s)          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │         TABLE_SCAN        │
 │    ────────────────────   │
 │           Table:          │
 │      pgbench_accounts     │
 │                           │
 │      100,000,000 rows     │
 │          (3.88s)          │
 └───────────────────────────┘
```



### Data Lake

Let's play with a local minio instance:

```sql
SELECT duckdb.create_simple_secret(
    type := 'S3', key_id := 's3user_data', secret := 'S3User.Data',
    endpoint := 'https://sss.pigsty:9000', url_style := 'path' 
);
```

