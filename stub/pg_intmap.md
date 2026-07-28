## Usage

Sources:

- [Official upstream README](https://github.com/adjust/pg_intmap/blob/5ddc90c908a79c2f97878008d8f98886ace4b3af/README.md)
- [Official extension control file (pg_intmap.control)](https://github.com/adjust/pg_intmap/blob/5ddc90c908a79c2f97878008d8f98886ace4b3af/pg_intmap.control)
- [Official extension SQL (pg_intmap--0.1.sql)](https://github.com/adjust/pg_intmap/blob/5ddc90c908a79c2f97878008d8f98886ace4b3af/pg_intmap--0.1.sql)

`pg_intmap` — Compressed integer-to-integer map. Use it when application data needs this type, domain, or its operators. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_intmap;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `intarr_get_val(intarr, int4)` is an extension function and returns `int8`.
- `intarr_in(cstring)` is an extension function and returns `intarr`.
- `intarr_out(intarr)` is an extension function and returns `cstring`.
- `intmap(int8[], int8[])` is an extension function and returns `intmap`.
- `intmap_get_val(intmap, int8)` is an extension function and returns `int8`.
- `intmap_in(cstring)` is an extension function and returns `intmap`.
- `intmap_meta(intmap)` is an extension function and returns `cstring`.
- `intmap_out(intmap)` is an extension function and returns `cstring`.
- `intarr` is an extension-defined type.
- `intmap` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
