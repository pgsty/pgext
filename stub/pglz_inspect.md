## Usage

Sources:

- [pgrx-examples/pglz_inspect/pglz_inspect.control](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/pglz_inspect.control)
- [pgrx-examples/pglz_inspect/README.md](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/README.md)
- [pgrx-examples/pglz_inspect/Cargo.toml](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/Cargo.toml)
- [pgrx-examples/pglz_inspect/src/lib.rs](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/src/lib.rs)
- [LICENSE](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/LICENSE)

`pglz_inspect` is a usable PGRX example extension for measuring PGLZ compression on individual values or sampled columns. Its unpublished example manifest uses version 0.0.0.

### Core Workflow

```sql
CREATE EXTENSION pglz_inspect;
SELECT * FROM pglz_size(convert_to(repeat('abc', 100), 'UTF8'));
CREATE TEMP TABLE compression_sample AS SELECT repeat('abc', 100) AS payload;
SELECT * FROM pglz_analyze_column('compression_sample'::regclass, 'payload', 1000, 'default');
SELECT pglz_recommend('compression_sample'::regclass, 'payload', 1000);
```

### Operational Boundaries

`pglz_size` returns raw/compressed byte counts, ratio and acceptance; `pglz_analyze_column` reports sampled statistics, `pglz_ratio_histogram` groups ratios, and `pglz_recommend` returns advice. Column sampling uses an ORDER BY random() query, which can scan and sort substantial data; estimates depend on samples and planner row estimates. Functions run with the caller's table permissions and do not change column storage. The control is not superuser-only and does not mark the extension trusted; no preload is specified. The README documents PostgreSQL 13–18; the manifest additionally defines a PG19 build feature, not a verified support claim.
