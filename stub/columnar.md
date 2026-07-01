


## Usage

Sources:

- [Hydra v1.1.2 README](https://github.com/hydradatabase/columnar/blob/v1.1.2/README.md)
- [Hydra Columnar README](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/README.md)
- [Columnar storage README](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/src/backend/columnar/README.md)
- [columnar.control](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/src/backend/columnar/columnar.control)
- [CHANGELOG 1.1.2](https://github.com/hydradatabase/columnar/blob/v1.1.2/CHANGELOG.md)

> [!WARNING]
> `columnar` is the PostgreSQL extension name, while Pigsty packages it as `hydra`. Pigsty metadata marks this extension as obsolete and no longer maintained; local packages are retained for PostgreSQL 14-16 only. Prefer newer actively maintained analytics extensions for new deployments.

Hydra Columnar is a PostgreSQL table access method for column-oriented storage. It stores selected tables in a columnar format to reduce I/O for analytical scans, compression-heavy datasets, projections over a subset of columns, and aggregate queries. It originated from Citus Columnar and is exposed through `CREATE EXTENSION columnar`.

### Create Columnar Tables

```sql
CREATE EXTENSION IF NOT EXISTS columnar;

CREATE TABLE events_columnar (
  event_id     bigint,
  account_id   bigint,
  event_time   timestamptz,
  event_type   text,
  amount       numeric
) USING columnar;

INSERT INTO events_columnar
SELECT
  g,
  g % 10000,
  now() - (g || ' seconds')::interval,
  CASE WHEN g % 10 = 0 THEN 'purchase' ELSE 'view' END,
  (g % 1000)::numeric / 10
FROM generate_series(1, 1000000) AS g;

SELECT event_type, count(*), sum(amount)
FROM events_columnar
WHERE event_time >= now() - interval '1 day'
GROUP BY event_type;
```

Use `USING columnar` when creating a table or materialized view. Reads and bulk inserts use normal PostgreSQL SQL, but the storage format is optimized for large analytical scans rather than high-churn OLTP tables.

### Table Options

```sql
SELECT columnar.alter_columnar_table_set(
  'events_columnar'::regclass,
  compression           => 'zstd',
  compression_level     => 3,
  stripe_row_limit      => 150000,
  chunk_group_row_limit => 10000
);

SELECT * FROM columnar.options;

SELECT columnar.alter_columnar_table_reset(
  'events_columnar'::regclass,
  compression => true,
  stripe_row_limit => true
);
```

Available table options include `compression`, `compression_level`, `stripe_row_limit`, and `chunk_group_row_limit`. Compression choices depend on build support, but documented values include `none`, `pglz`, `zstd`, `lz4`, and `lz4hc`. Option changes apply to newly inserted data; existing stripes are not automatically rewritten.

You can also set defaults for newly created columnar tables with GUCs:

```sql
SET columnar.compression = 'zstd';
SET columnar.compression_level = 3;
SET columnar.stripe_row_limit = 150000;
SET columnar.chunk_group_row_limit = 10000;
```

These GUCs affect newly created tables, not new stripes on an already existing table.

### Convert Existing Tables

```sql
CREATE TABLE events_heap (
  event_id bigint,
  payload  jsonb
);

INSERT INTO events_heap
SELECT g, jsonb_build_object('kind', 'view', 'seq', g)
FROM generate_series(1, 10000) AS g;

SELECT columnar.alter_table_set_access_method('events_heap', 'columnar');
SELECT columnar.alter_table_set_access_method('events_heap', 'heap');
```

`columnar.alter_table_set_access_method(table, method)` rewrites a heap table as columnar storage or a columnar table back to heap storage. Review foreign keys, identity columns, row-level security, triggers, indexes, constraints, inheritance, and storage options before conversion; the helper rejects or skips features it cannot safely recreate.

### Partitioning

```sql
CREATE TABLE measurements (
  ts timestamptz,
  device_id bigint,
  value double precision
) PARTITION BY RANGE (ts);

CREATE TABLE measurements_2024 PARTITION OF measurements
  FOR VALUES FROM ('2024-01-01') TO ('2025-01-01')
  USING columnar;

CREATE TABLE measurements_hot PARTITION OF measurements
  FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
```

Partitioned tables can mix row and columnar partitions. Operations that target only row partitions can use row-table behavior, while operations that touch columnar partitions must respect columnar limitations. This is useful for keeping recent mutable data in heap partitions and older analytical history in columnar partitions.

### Maintenance and Introspection

```sql
VACUUM VERBOSE events_columnar;
VACUUM FULL events_columnar;

SELECT * FROM columnar.stats('events_columnar'::regclass);
SELECT columnar.vacuum('events_columnar'::regclass);
SELECT columnar.vacuum_full('public', 0.1, 25);
```

`VACUUM VERBOSE` reports columnar storage statistics such as file size, compression rate, row count, stripe count, and chunk count. `columnar.stats()` exposes stripe-level metadata. `columnar.vacuum()` and `columnar.vacuum_full()` compact columnar storage incrementally; ordinary `VACUUM` is metadata-oriented and cheaper than a full rewrite.

### Caveats

- This extension is obsolete in Pigsty metadata and conflicts with `citus`/`citus_columnar` style columnar storage. Avoid installing conflicting columnar table access methods in the same PostgreSQL major unless you have tested the exact combination.
- Pigsty packages `hydra`/`columnar` for PostgreSQL 14-16; PostgreSQL 17 and 18 are marked unsupported locally.
- Hydra 1.1.x added update/delete and upsert improvements, but the project itself still describes columnar storage as unsuitable for frequent large updates, small transactions, and OLTP-style single-row workloads.
- Unsupported or limited areas include logical decoding, unlogged columnar tables, serializable isolation, some scan types, and many non-btree/non-hash indexes. Check constraints and index-backed constraints carefully before relying on them.
- The `columnar` schema contains internal metadata tables such as `columnar.options`, `columnar.stripe`, `columnar.chunk_group`, and `columnar.chunk`. Query public views/functions for inspection, but do not mutate metadata tables directly.
