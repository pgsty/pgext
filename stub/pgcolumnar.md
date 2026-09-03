## Usage

Sources:

- [Official documentation](https://github.com/commandprompt/pgcolumnar/blob/0e4884c18a678bf0a990e6a7dfcdd47248d9111d/README.md)
- [Extension control file](https://github.com/commandprompt/pgcolumnar/blob/0e4884c18a678bf0a990e6a7dfcdd47248d9111d/pgcolumnar.control)
- [Official repository](https://github.com/commandprompt/pgcolumnar)

`pgcolumnar` Native column-oriented table access method with compression, vectorized scans, and Parquet workflows.

### Enablement

Merge `pgcolumnar` into the existing preload list, restart PostgreSQL, and then create `pgcolumnar` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pgcolumnar'
```

```sql
CREATE EXTENSION pgcolumnar;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE EXTENSION pgcolumnar;

CREATE TABLE events (id bigint, ts timestamptz, kind int, payload text)
  USING pgcolumnar;

INSERT INTO events
  SELECT g, now(), g % 8, 'p' || g
  FROM generate_series(1, 1000000) g;

SELECT count(*), avg(kind) FROM events WHERE kind = 3;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pgcolumnar.projection` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pgcolumnar.bloom` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pgcolumnar.export_arrow` | FUNCTION | Callable function from the reviewed install surface. |
| `pgcolumnar.export_parquet` | FUNCTION | Callable function from the reviewed install surface. |
| `pgcolumnar.options` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pgcolumnar.storage` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pgcolumnar.vacuum_sorted` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 15, 16, 17, 18, 19; do not infer unlisted majors.
- Preloading `pgcolumnar` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
- The extension fixes or creates schema objects under `pgcolumnar`; include them in privilege and backup review.
- Catalog lifecycle is preview; test upgrades, dump/restore, and server compatibility before production use.
