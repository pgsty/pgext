## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_approx.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_approx/alohadb_approx.control)
- [Official extension SQL (alohadb_approx--1.0--1.1.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_approx/alohadb_approx--1.0--1.1.sql)

`alohadb_approx` — Approximate query processing: HLL, Count-Min Sketch, Top-K. Use it when SQL needs these specialized functions or aggregates. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_approx;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `approx_count_distinct_finalfn(internal)` is an extension function and returns `int8`.
- `approx_count_distinct_transfn(internal, anyelement)` is an extension function and returns `internal`.
- `approx_percentile_finalfn(internal)` is an extension function and returns `float8`.
- `approx_percentile_transfn(internal, float8, float8, float8)` is an extension function and returns `internal`.
- `bloom_add(bf bloom_filter, item text)` is an extension function and returns `bloom_filter`.
- `bloom_agg_finalfn(internal)` is an extension function and returns `bloom_filter`.
- `bloom_agg_transfn(internal, text, int, float8)` is an extension function and returns `internal`.
- `bloom_contains(bf bloom_filter, item text)` is an extension function and returns `boolean`.
- `bloom_create(expected_items int, fpr float8 DEFAULT 0.01)` is an extension function and returns `bloom_filter`.
- `bloom_in(cstring)` is an extension function and returns `bloom_filter`.
- `bloom_merge(bf1 bloom_filter, bf2 bloom_filter)` is an extension function and returns `bloom_filter`.
- `bloom_out(bloom_filter)` is an extension function and returns `cstring`.
- `bloom_stats(bf bloom_filter)` is an extension function and returns `TABLE`.
- `cms_add(cms, text)` is an extension function and returns `cms`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
