## Usage

Sources:

- [README](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/README.md)
- [Control file](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/pgcolumnar.control)
- [SQL](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/pgcolumnar--1.0-alpha4.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/CHANGELOG.md)
- [docs/limitations.md](https://api.pgxn.org/src/pgcolumnar/pgcolumnar-1.0.0-alpha.4/docs/limitations.md)

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

### Alpha 4 Upgrade

The control version is `1.0-alpha4`, distributed on PGXN as 1.0.0-alpha.4. Install matching files and run `ALTER EXTENSION pgcolumnar UPDATE` in each database; replacing the library alone is insufficient. Alpha 4 adds Hilbert clustering and further correctness fixes. PostgreSQL 19 evidence is against beta2. Keep original data reloadable: the project still provides no general promise of compatibility across future on-disk format changes. Some older limitation-page release labels remain stale; the control file and current changelog identify this release.
