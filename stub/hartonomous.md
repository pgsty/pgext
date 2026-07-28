## Usage

Sources:

- [Official upstream README](https://github.com/saltypatron/hartonomous-001/blob/1ea12b4dfaa0a24b46ffa8c41290d9e1767fec4b/README.md)
- [Official extension control file (hartonomous.control)](https://github.com/saltypatron/hartonomous-001/blob/1ea12b4dfaa0a24b46ffa8c41290d9e1767fec4b/ext/hartonomous_pg/hartonomous.control)
- [Official extension SQL (hartonomous--1.0.sql)](https://github.com/saltypatron/hartonomous-001/blob/1ea12b4dfaa0a24b46ffa8c41290d9e1767fec4b/ext/hartonomous_pg/sql/hartonomous--1.0.sql)

`hartonomous` — Hartonomous substrate — schemas, types, tables, BLAKE3, S^3 geometry, traversal, UCD/UCA atoms. Use it when an application needs this specific database capability. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION hartonomous;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `antipode(point4d)` is an extension function and returns `point4d`.
- `array_to_linestring4d(double precision[])` is an extension function and returns `linestring4d`.
- `array_to_point4d(double precision[])` is an extension function and returns `point4d`.
- `bbox(linestring4d)` is an extension function and returns `box4d`.
- `bbox(point4d)` is an extension function and returns `box4d`.
- `bbox_4d_combine(box4d, box4d)` is an extension function and returns `box4d`.
- `bbox_4d_sfunc(box4d, point4d)` is an extension function and returns `box4d`.
- `bbox_expand(box4d, point4d)` is an extension function and returns `box4d`.
- `bbox_union(box4d, box4d)` is an extension function and returns `box4d`.
- `blake3_hash(bytea)` is an extension function and returns `bytea`.
- `blake3_hash_text(text)` is an extension function and returns `bytea`.
- `box4d_contained_by_box(box4d, box4d)` is an extension function and returns `boolean`.
- `box4d_contains_box(box4d, box4d)` is an extension function and returns `boolean`.
- `box4d_contains_point(box4d, point4d)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `postgis`, `btree_gist`, `pg_trgm`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream explicitly says the project is not production-ready.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
