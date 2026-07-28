## Usage

Sources:

- [Official upstream README](https://github.com/evancarroll/pg-srf-repeat-benchmark/blob/866ae8b1d657bb03cf71d7e6f52a5406afd64aca/README.md)
- [Official extension control file (repeat.control)](https://github.com/evancarroll/pg-srf-repeat-benchmark/blob/866ae8b1d657bb03cf71d7e6f52a5406afd64aca/repeat.control)
- [Official extension SQL (repeat--0.0.1.sql)](https://github.com/evancarroll/pg-srf-repeat-benchmark/blob/866ae8b1d657bb03cf71d7e6f52a5406afd64aca/repeat--0.0.1.sql)

`repeat` — PostgreSQL Set-Returning-Function (SRF) C-Extension Benchmarks ====. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION repeat;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `repeat_materialize(object int4, times int4)` is an extension function and returns `TABLE`.
- `repeat_materialize_preferred(object int4, times int4)` is an extension function and returns `TABLE`.
- `repeat_valuepercall(object int4, times int4)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
