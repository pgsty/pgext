## Usage

Sources:

- [Upstream source 1](https://github.com/justusepostgres/pg_lamina/blob/9ba7509f39b362bc487cc3c5843e03f73769b24a/README.md)
- [Upstream source 2](https://github.com/justusepostgres/pg_lamina/blob/9ba7509f39b362bc487cc3c5843e03f73769b24a/pg_lamina.control)

`pg_lamina` 0.2.0 converts an ordinary heap into PostgreSQL-hosted Parquet segments queried through DataFusion.

### Core Workflow

Successful conversion drops the source heap in the same transaction. `analytics.convert_table_sorted` optionally sorts by named columns first. `analytics.query` accepts DataFusion SQL over the registered analytical tables, and `analytics.storage_compression` chooses compression for future conversions.

```sql
CREATE EXTENSION pg_lamina;
CREATE TABLE lamina_source (region text, amount bigint);
INSERT INTO lamina_source VALUES ('west', 125), ('west', 200);
BEGIN;
SELECT analytics.convert_table('lamina_source', 'events');
SELECT * FROM analytics.query(
  'SELECT region, sum(amount) AS total FROM events GROUP BY region'
);
COMMIT;
```

### Operational Boundaries

This PostgreSQL 17 ARM64 research prototype is not a continuously synchronized copy or a table access method. Conversion takes an exclusive lock; source indexes, constraints, triggers, privileges and relation identity are not transferred. Preserve the source independently before cutover, review dependent objects and budget temporary space. The SQL API and storage contract are experimental.
