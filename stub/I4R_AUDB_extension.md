## Usage

Sources:

- [Official upstream README](https://github.com/asxvi/audb/blob/6018074c6e7e7e8416ccc7a04f1de9365d2f44dc/README.md)
- [Official extension control file (I4R_AUDB_extension.control)](https://github.com/asxvi/audb/blob/6018074c6e7e7e8416ccc7a04f1de9365d2f44dc/c_extension/i4r_audb_extension/I4R_AUDB_extension.control)
- [Official extension SQL (i4r_audb_extension--1.1.sql)](https://github.com/asxvi/audb/blob/6018074c6e7e7e8416ccc7a04f1de9365d2f44dc/c_extension/i4r_audb_extension/i4r_audb_extension--1.1.sql)

`I4R_AUDB_extension` — AUDB operations for PostgreSQL int4range values. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "I4R_AUDB_extension";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `agg_avg_range_finalfunc(internal)` is an extension function and returns `int4range`.
- `agg_avg_range_transfunc(state internal, range int4range, mult int4range)` is an extension function and returns `internal`.
- `agg_count_transfunc(state int4range, input int4range)` is an extension function and returns `int4range`.
- `agg_max_range_transfunc(state int4range, input int4range)` is an extension function and returns `int4range`.
- `agg_max_set_transfunc(state int4range[], input int4range[])` is an extension function and returns `int4range[]`.
- `agg_min_max_set_finalfunc(int4range[])` is an extension function and returns `int4range[]`.
- `agg_min_range_transfunc(state int4range, input int4range)` is an extension function and returns `int4range`.
- `agg_min_set_transfunc(state int4range[], input int4range[])` is an extension function and returns `int4range[]`.
- `agg_sum_range_transfunc(int4range, int4range)` is an extension function and returns `int4range`.
- `agg_sum_set_finalfunc(internal)` is an extension function and returns `int4range[]`.
- `agg_sum_set_finalfunc_metrics(internal)` is an extension function and returns `sum_set_metrics`.
- `agg_sum_set_transfunc(internal, set int4range[], resizeTrigger integer, reduceToSize integer)` is an extension function and returns `internal`.
- `agg_sum_set_transfunc_metrics(internal, int4range[], integer, integer, bool)` is an extension function and returns `internal`.
- `array_length(set int4range[])` is an extension function and returns `int4`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
