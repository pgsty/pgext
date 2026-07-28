## Usage

Sources:

- [Official upstream README](https://github.com/michelp/pgs2/blob/2a193a0f76578eb78eada9041a5da8d85e449f3b/README.md)
- [Official extension control file (pgs2.control)](https://github.com/michelp/pgs2/blob/2a193a0f76578eb78eada9041a5da8d85e449f3b/pgs2.control)
- [Official extension SQL (pgs2--0.0.1.sql)](https://github.com/michelp/pgs2/blob/2a193a0f76578eb78eada9041a5da8d85e449f3b/pgs2--0.0.1.sql)

`pgs2` — Postgres extension for S2 spherical geometry. Use it for the corresponding spatial data or geospatial workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgs2;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `btS2Cellsortsupport(internal)` is an extension function and returns `void`.
- `S2Cap(center S2Point, radium float8 = 0.0)` is an extension function and returns `S2Cap`.
- `S2Cap_eq(A S2Cap, B S2Cap)` is an extension function.
- `S2Cap_in(cstring)` is an extension function and returns `S2Cap`.
- `S2Cap_out(S2Cap)` is an extension function and returns `cstring`.
- `S2Cell_as_S2LatLng(c S2Cell)` is an extension function and returns `S2LatLng`.
- `S2Cell_as_S2Point(c S2Cell)` is an extension function and returns `S2Point`.
- `S2Cell_cmp(A S2Cell, B S2Cell)` is an extension function and returns `int`.
- `S2Cell_distance(A S2Cell, B S2Cell)` is an extension function and returns `float8`.
- `S2Cell_eq(A S2Cell, B S2Cell)` is an extension function.
- `S2Cell_ge(A S2Cell, B S2Cell)` is an extension function.
- `S2Cell_gt(A S2Cell, B S2Cell)` is an extension function.
- `S2Cell_in(cstring)` is an extension function and returns `S2Cell`.
- `S2Cell_le(A S2Cell, B S2Cell)` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
