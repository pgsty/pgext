## Usage

Sources:

- [Official README](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/README.md)
- [Extension control](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/ch_table_engines.control)
- [Versioned SQL](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/sql/ch_table_engines--0.1.sql)
- [Preload boundary test](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/sql/nopreload.sql)
- [Settings](https://github.com/igor-suhorukov/pg_ch_table_engines/blob/99c5a49d4aa412a384134004e60d1eacc22798c8/src/module.c)

`ch_table_engines` 0.1 exposes ClickHouse MergeTree-family table access methods on unmodified PostgreSQL 19, with PostgreSQL catalogs controlling snapshot visibility. It needs the separately supplied `libchte_engine.so`, built from ClickHouse 26.8.8.8-lts. This early project is independent of ClickHouse, Inc.

### Core Workflow

Permanent tables require preloading and a PostgreSQL restart. Configure the engine library through `ch.engine_library` if its default location does not match the installation:

```ini
shared_preload_libraries = 'ch_table_engines'
```

As a superuser, install the non-relocatable extension, which owns schema `ch`:

```sql
CREATE EXTENSION ch_table_engines;
CREATE TABLE events (account_id int, event_day date, amount bigint)
    USING mergetree WITH (order_by = 'account_id, event_day', storage = 'pages');
INSERT INTO events VALUES (1, DATE '2026-10-22', 100);
SELECT account_id, sum(amount) FROM events GROUP BY account_id;
SELECT * FROM ch.parts('events'::regclass);
```

TEMP and UNLOGGED tables can work without preload, but permanent tables need the WAL resource manager registered at server startup. Normal persistent deployments should use the preloaded configuration, which also enables merge workers.

### Engines, Options and Maintenance

The installed methods are `mergetree`, `replacingmergetree`, `summingmergetree`, `aggregatingmergetree`, `collapsingmergetree`, `versionedcollapsingmergetree` and `coalescingmergetree`. At this revision, automatic and session merges apply only to MergeTree tables; do not assume that the other methods already perform their expected merge reductions.

- `storage` chooses buffered relation pages or directories of files; `order_by`, `primary_key`, `index_granularity`, `compression` and `low_cardinality` describe part layout. Layout-defining options change only on empty tables. The `primary_key` option is a storage key, not a PostgreSQL uniqueness constraint.
- `ch.parts`, `ch.definition`, `ch.table_size` and `ch.engine_info` inspect tables and the engine.
- `ch.fold` rewrites parts to omit rows deleted for every snapshot; `ch.optimize(rel, final)` merges parts. VACUUM releases obsolete part storage after catalog visibility permits it.
- `ch.max_merge_workers` defaults to 2 and `ch.merge_naptime` to 10 seconds. Workers do not merge on standbys.

### Boundaries

Unique indexes, exclusion constraints, concurrent index builds and BRIN are unsupported. This also rules out primary-key constraints and unique arbiters for ON CONFLICT. Row-lock modes currently all behave exclusively. Test the required type, transaction, backup and recovery behavior before adoption; the index and merge restrictions are material storage limitations.
