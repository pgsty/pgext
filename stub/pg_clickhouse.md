

## Usage

> [pg_clickhouse: ClickHouse integration for PostgreSQL](https://github.com/ClickHouse/pg_clickhouse)

`pg_clickhouse` enables analytics queries on ClickHouse directly from PostgreSQL without SQL rewriting.
It supports PostgreSQL 13+ and ClickHouse v23+.

### Create the Extension

```sql
CREATE EXTENSION pg_clickhouse;
```

Or into a specific schema:

```sql
CREATE SCHEMA env;
CREATE EXTENSION pg_clickhouse SCHEMA env;
```

### Query Pushdown

The extension automatically pushes down analytical queries to ClickHouse for execution,
providing significant performance improvements. For example, TPC-H benchmarks show:

- Query 1: 268ms (vs 4,693ms on standard PostgreSQL)
- Query 6: 53ms (vs 764ms on standard PostgreSQL)

When query pushdown is active, ClickHouse handles execution directly, avoiding
data transfer overhead for complex analytical workloads.
