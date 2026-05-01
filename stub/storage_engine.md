## Usage

Sources: [README](https://github.com/saulojb/storage_engine/blob/main/README.md), [release 1.3.4](https://github.com/saulojb/storage_engine/releases/tag/v1.3.4), [PGXN 1.3.4](https://pgxn.org/dist/storage_engine/1.3.4/), [PGXN changelog](https://pgxn.org/dist/storage_engine/1.3.4/CHANGELOG.html), [META.json](https://github.com/saulojb/storage_engine/blob/main/META.json)

`storage_engine` provides two PostgreSQL table access methods in the `engine` schema:

- `colcompress` for column-oriented compressed storage with vectorized execution, min/max pruning, and parallel scans.
- `rowcompress` for row-batch compression with parallel scans.

```sql
CREATE EXTENSION storage_engine;
```

### Quick Start

Create tables using either access method:

```sql
CREATE TABLE events (
  ts timestamptz NOT NULL,
  user_id bigint,
  event_type text,
  value float8
) USING colcompress;

CREATE TABLE logs (
  id bigserial,
  logged_at timestamptz NOT NULL,
  message text
) USING rowcompress;
```

### Main Tuning Knobs

Session-level GUCs documented upstream include:

- `storage_engine.enable_parallel_execution`
- `storage_engine.min_parallel_processes`
- `storage_engine.enable_vectorization`
- `storage_engine.enable_custom_scan`
- `storage_engine.enable_column_cache`
- `storage_engine.enable_columnar_index_scan`
- `storage_engine.enable_dml`
- `storage_engine.stripe_row_limit`
- `storage_engine.chunk_group_row_limit`
- `storage_engine.compression_level`

The README says these GUCs become visible once the library is loaded; add `storage_engine` to `shared_preload_libraries` if you want them available immediately in every session.

### Types and Operators

`engine.uint8` stores unsigned 64-bit values for `colcompress` workloads that need the full `0` through `2^64 - 1` range. Upstream documents comparison operators (`=`, `<>`, `<`, `<=`, `>`, `>=`), B-tree and hash opclasses, casts to and from `bigint`, `numeric`, and `text`, plus `engine.min`, `engine.max`, and `engine.sum` aggregates.

### Useful Management Functions

For `colcompress` tables:

```sql
SELECT engine.alter_colcompress_table_set(
  'events'::regclass,
  orderby => 'ts ASC, user_id ASC',
  compression => 'zstd',
  compression_level => 9
);

SELECT engine.colcompress_merge('events');
CALL engine.colcompress_repack('events');
CALL engine.colcompress_repack('events', 0.7);
```

In 1.3.4, `engine.colcompress_repack(table_name regclass, min_fill_ratio float8 DEFAULT 0.9)` is a procedure for online per-stripe defragmentation of `colcompress` tables. It repacks stripes whose live-row ratio falls below the threshold. Use `engine.colcompress_merge()` when you need the old full-table rewrite that globally sorts by the `orderby` key.

For `rowcompress` tables:

```sql
SELECT engine.alter_rowcompress_table_set(
  'logs'::regclass,
  batch_size => 10000,
  compression => 'zstd',
  compression_level => 5
);

SELECT engine.rowcompress_repack('logs');
```

### When to Use Which AM

- Use `colcompress` for analytical scans, aggregates, and range predicates where projection, vectorization, and stripe/chunk pruning pay off.
- Use `rowcompress` for append-heavy logs or wide rows that are usually fetched together, where compression matters more than column projection.
- For point lookups on `colcompress`, upstream recommends enabling `storage_engine.enable_columnar_index_scan` or per-table `index_scan`.

### Caveats

- Upgrade existing installations with `ALTER EXTENSION storage_engine UPDATE TO '1.3.4';`.
- The old `FUNCTION engine.colcompress_repack(regclass)` alias was removed in 1.3.4. Existing callers should switch to `CALL engine.colcompress_repack('table')` for stripe defragmentation or `SELECT engine.colcompress_merge('table')` for the old full rewrite behavior.
- `colcompress` and `rowcompress` do not support foreign keys or `AFTER ROW` triggers.
- `VACUUM FULL` and `CREATE UNLOGGED TABLE ... USING colcompress` are not supported; upstream recommends the extension's repack functions instead.
- On `colcompress`, combining `orderby` with B-tree indexes can disable the sort-on-write path, and B-tree indexes on ordered columns can defeat stripe pruning for range queries. Use `engine.colcompress_merge()` after loading data when global ordering matters, and prefer `index_scan => false` for analytical tables.
