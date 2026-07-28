## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/varint/varint-0.1.0/README)
- [Official extension control file (varint.control)](https://api.pgxn.org/src/varint/varint-0.1.0/varint.control)
- [Official extension SQL (varint.sql)](https://api.pgxn.org/src/varint/varint-0.1.0/sql/varint.sql)

`varint` — Intro to PostgreSQL-varint varint is a data type for PostgreSQL that encodes integers using variable width encoding in order to save space. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION varint;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `int2(varint64)` is an extension function and returns `int2`.
- `int2(varuint64)` is an extension function and returns `int2`.
- `int4(varint64)` is an extension function and returns `int4`.
- `int4(varuint64)` is an extension function and returns `int4`.
- `int8(varint64)` is an extension function and returns `int8`.
- `int8(varuint64)` is an extension function and returns `int8`.
- `varint64(int2)` is an extension function and returns `varint64`.
- `varint64(int4)` is an extension function and returns `varint64`.
- `varint64(int8)` is an extension function and returns `varint64`.
- `varint64_cmp(varint64, varint64)` is an extension function and returns `int4`.
- `varint64_eq(varint64, varint64)` is an extension function.
- `varint64_ge(varint64, varint64)` is an extension function.
- `varint64_gt(varint64, varint64)` is an extension function.
- `varint64_in(cstring)` is an extension function and returns `varint64`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
