## Usage

Sources:

- [Official upstream README](https://github.com/vexdb-thu/vexdb-lite/blob/315dd536ca1446f1e563594687254647a06790c7/README.md)
- [Official extension control file (vexdb_lite.control)](https://github.com/vexdb-thu/vexdb-lite/blob/315dd536ca1446f1e563594687254647a06790c7/vexdb_pg/vexdb_lite.control)
- [Official extension SQL (vexdb_lite--1.0.sql)](https://github.com/vexdb-thu/vexdb-lite/blob/315dd536ca1446f1e563594687254647a06790c7/vexdb_pg/sql/vexdb_lite--1.0.sql)

`vexdb_lite` — A cross-platform vector database, which can be integrated into existing databases as a plugin. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION vexdb_lite;

CREATE TABLE items (
    id  BIGSERIAL PRIMARY KEY,
    vec floatvector(128)
);

INSERT INTO items (vec) VALUES
    ('[0.10, 0.20, 0.30]'),
    ('[0.40, 0.50, 0.60]');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `array_to_floatvector(double precision[], integer, boolean)` is an extension function and returns `floatvector`.
- `array_to_floatvector(integer[], integer, boolean)` is an extension function and returns `floatvector`.
- `array_to_floatvector(numeric[], integer, boolean)` is an extension function and returns `floatvector`.
- `array_to_floatvector(real[], integer, boolean)` is an extension function and returns `floatvector`.
- `cosine_distance(floatvector, floatvector)` is an extension function and returns `float8`.
- `floatvector(floatvector, integer, boolean)` is an extension function and returns `floatvector`.
- `floatvector_add(floatvector, floatvector)` is an extension function and returns `floatvector`.
- `floatvector_cmp(floatvector, floatvector)` is an extension function and returns `int4`.
- `floatvector_eq(floatvector, floatvector)` is an extension function.
- `floatvector_ge(floatvector, floatvector)` is an extension function.
- `floatvector_gt(floatvector, floatvector)` is an extension function.
- `floatvector_in(cstring, oid, integer)` is an extension function and returns `floatvector`.
- `floatvector_l2_squared_distance(floatvector, floatvector)` is an extension function and returns `float8`.
- `floatvector_le(floatvector, floatvector)` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
