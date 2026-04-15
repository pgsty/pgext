## Usage

> Sources: [extension README](https://github.com/boringSQL/regresql/blob/master/pg_ext/README.md), [control file](https://github.com/boringSQL/regresql/blob/master/pg_ext/pg_regresql.control), [portable stats article](https://boringsql.com/posts/portable-stats/)

`pg_regresql` is a PostgreSQL extension that makes the planner trust catalog statistics from `pg_class` instead of recomputing relation size from physical file blocks. It is the extension part of the RegreSQL project, intended for realistic plan regression testing with injected production statistics.

### Problem

The upstream extension README explains that PostgreSQL normally does not fully trust `pg_class.relpages` and `pg_class.reltuples` when estimating relation size. Instead, planner code reads the current physical file size and rescales statistics from that.

That behavior is useful for stale-statistics safety, but it breaks test setups where catalog statistics were intentionally restored from another environment and the local table files are much smaller.

### What It Overrides

`pg_regresql` hooks into `get_relation_info_hook` after `estimate_rel_size()` and replaces planner estimates with catalog values.

| Planner field | Default source | `pg_regresql` source |
| --- | --- | --- |
| `rel->pages` | `smgrnblocks()` via table access method | `pg_class.relpages` |
| `rel->tuples` | density scaled by physical pages | `pg_class.reltuples` |
| `rel->allvisfrac` | `relallvisible / physical pages` | `pg_class.relallvisible / relpages` |
| `IndexOptInfo->pages` | `RelationGetNumberOfBlocks()` | `pg_class.relpages` for the index |
| `IndexOptInfo->tuples` | copied from `rel->tuples` | `pg_class.reltuples` for the index |

### Installation

The upstream README documents three installation paths:

```bash
sudo pgxn install pg_regresql
```

```bash
make PG_SOURCE=/path/to/postgresql
make install PG_SOURCE=/path/to/postgresql
```

```bash
make USE_PGXS=1
make install USE_PGXS=1
```

The control file ships as `pg_regresql.control` with `default_version = '2.0'` and `module_pathname = '$libdir/pg_regresql'`.

### Activation

The extension becomes active when the shared library is loaded:

```sql
LOAD 'pg_regresql';

EXPLAIN SELECT ...;
```

For a whole test instance, the README recommends:

```conf
session_preload_libraries = 'pg_regresql'
```

This is the important runtime configuration: package installation alone is not the point; the planner hook only takes effect after the library is loaded for the session or instance.

### Typical Workflow

The main use case is plan regression testing with restored production statistics. After injecting catalog statistics into a CI or test database, `pg_regresql` makes the planner use those restored values instead of the tiny local heap size.

The README gives this example:

```sql
SELECT pg_restore_relation_stats(
    'schemaname', 'public',
    'relname', 'test_orders',
    'relpages', 123513::integer,
    'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);

LOAD 'pg_regresql';

EXPLAIN SELECT * FROM test_orders WHERE created_at > '2024-06-01';
```

That pattern is useful for:

- reproducing production plans locally
- testing schema migrations against realistic plan estimates
- simulating table growth and index choices
- improving partition-planning experiments

### Compatibility

- PostgreSQL 14 and newer in this repository packaging
- upstream README notes the hook itself exists since PostgreSQL 8.3
- intended to coexist with extensions such as `pg_hint_plan` and `hypopg`, though upstream marks that as not yet fully tested
