## Usage

Sources:

- [Official upstream README](https://github.com/daviszhen/pg_hybrid/blob/7e5d30c151d263f5101c3f61cef969a67646dadc/README.md)
- [Official extension control file (pg_hybrid.control)](https://github.com/daviszhen/pg_hybrid/blob/7e5d30c151d263f5101c3f61cef969a67646dadc/pg_hybrid.control)
- [Official extension SQL (pg_hybrid--1.0.sql)](https://github.com/daviszhen/pg_hybrid/blob/7e5d30c151d263f5101c3f61cef969a67646dadc/pg_hybrid--1.0.sql)

`pg_hybrid` — PostgreSQL 16 extension providing IVFFlat index access method for vector similarity search. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_hybrid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `array_to_hvector(double precision[], integer, boolean)` is an extension function and returns `hvector`.
- `array_to_hvector(integer[], integer, boolean)` is an extension function and returns `hvector`.
- `array_to_hvector(numeric[], integer, boolean)` is an extension function and returns `hvector`.
- `array_to_hvector(real[], integer, boolean)` is an extension function and returns `hvector`.
- `hvector(hvector, integer, boolean)` is an extension function and returns `hvector`.
- `hvector_accum(double precision[], hvector)` is an extension function and returns `double`.
- `hvector_add(hvector, hvector)` is an extension function and returns `hvector`.
- `hvector_avg(double precision[])` is an extension function and returns `hvector`.
- `hvector_binary_quantize(hvector)` is an extension function and returns `bit`.
- `hvector_cmp(hvector, hvector)` is an extension function and returns `int4`.
- `hvector_combine(double precision[], double precision[])` is an extension function and returns `double`.
- `hvector_concat(hvector, hvector)` is an extension function and returns `hvector`.
- `hvector_cosine_distance(hvector, hvector)` is an extension function and returns `float8`.
- `hvector_dims(hvector)` is an extension function and returns `integer`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
