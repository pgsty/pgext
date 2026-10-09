## Usage

Sources:

- [README.md](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/README.md)
- [sql/ybgist--1.2.sql](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/sql/ybgist--1.2.sql)
- [docs/INSTALL-RELEASE.md](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/docs/INSTALL-RELEASE.md)
- [ybgist.control](https://github.com/makeafide/yb-pggist/blob/0ecfd02df2eca84a1cbfbdc977789a12331680ca/ybgist.control)

`ybgist` supplies spatial operator classes and guard functions for the native access method in a specific YugabyteDB fork. It is not installable on stock PostgreSQL or stock YugabyteDB.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ybgist;
SELECT ybgist_tune_costs();
```

### Operational Boundaries

Deploy the matching core fork, patched PostGIS and extension as one compatible stack; its C++ runtime and ABI requirements are part of the installation contract. Create PostGIS first, then ybgist as a superuser. The fixed schema is public. Bulk-load data before building indexes in separate statements. `ybgist_record_index` records build settings, `ybgist_verify_index` checks backfill completeness and `ybgist_tune_costs` adjusts predicate costs.

The source control version is 1.2; published r2 bundles predate current features. Geometry, geography and multicolumn scalar/spatial indexes are documented. Index-only scans and nearest-neighbor ordering are not supported. Changing mapping or coverage settings requires following the upstream drift/reindex procedure. No project license declaration was found in the inspected source tree.
