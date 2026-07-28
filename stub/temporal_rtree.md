## Usage

Sources:

- [Official upstream README](https://github.com/yash7312/dbis_assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporal_rtree/README)
- [Official extension control file (temporal_rtree.control)](https://github.com/yash7312/dbis_assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporal_rtree/temporal_rtree.control)
- [Official extension SQL (temporal_rtree--1.0.sql)](https://github.com/yash7312/dbis_assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporal_rtree/temporal_rtree--1.0.sql)

`temporal_rtree` — A PostgreSQL extension implementing a new index access method (temporal_rtree) optimized for temporal range indexing. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION temporal_rtree;

CREATE INDEX idx_temporal ON temporal_data
  USING temporal_rtree (temporalbox(attr, valid_period) temporal_cube_ops);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `temporal_rtree_handler(internal)` is an extension function and returns `index_am_handler`.
- `temporal_rtree_hook_reset()` is an extension function and returns `void`.
- `temporal_rtree_hook_stats()` is an extension function and returns `TABLE`.
- `FAMILY` is an extension-defined operator.
- `temporal_cube_ops` is an extension-defined operator class.
- `temporal_rtree` is an extension-defined access method.
- `temporal_tsrange_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
