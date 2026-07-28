## Usage

Sources:

- [Official upstream README](https://github.com/marthinl/ssk/blob/be8781c8a4d3cb943fde0e6f039207f087fd5c1e/README.md)
- [Official extension control file (ssk.control)](https://github.com/marthinl/ssk/blob/be8781c8a4d3cb943fde0e6f039207f087fd5c1e/ssk.control)
- [Official extension SQL (ssk--1.0.sql)](https://github.com/marthinl/ssk/blob/be8781c8a4d3cb943fde0e6f039207f087fd5c1e/sql/ssk--1.0.sql)

`ssk` — **Note:** PostgreSQL serves as host for the reference implementation for SSK, showcasing its universal applicability in a production database environment. While rooted in PostgreSQL, SSK's concepts extend to any relational database. Use it when an application needs this specific database capability. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION ssk;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `cardinality(ssk)` is an extension function and returns `bigint`.
- `length(ssk)` is an extension function and returns `bigint`.
- `size(ssk)` is an extension function and returns `bigint`.
- `ssk()` is an extension function and returns `ssk`.
- `ssk(bigint)` is an extension function and returns `ssk`.
- `ssk(bigint[])` is an extension function and returns `ssk`.
- `ssk(integer)` is an extension function and returns `ssk`.
- `ssk(integer[])` is an extension function and returns `ssk`.
- `ssk_add(ssk, bigint)` is an extension function and returns `ssk`.
- `ssk_add(ssk, integer)` is an extension function and returns `ssk`.
- `ssk_add_comm(bigint, ssk)` is an extension function and returns `ssk`.
- `ssk_agg_finalfunc(bigint)` is an extension function and returns `ssk`.
- `ssk_agg_sfunc(bigint, bigint)` is an extension function and returns `bigint`.
- `ssk_cmp(ssk, ssk)` is an extension function and returns `int`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
