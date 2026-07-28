## Usage

Sources:

- [Official upstream README](https://github.com/tanglebones/pg_tuid/blob/eca0bbd95cd17dda004d1adbe2546048de5650b9/README.md)
- [Official extension control file (tuid.control)](https://github.com/tanglebones/pg_tuid/blob/eca0bbd95cd17dda004d1adbe2546048de5650b9/pg_c/tuid.control)
- [Official extension SQL (tuid--0.3.0.sql)](https://github.com/tanglebones/pg_tuid/blob/eca0bbd95cd17dda004d1adbe2546048de5650b9/pg_c/tuid--0.3.0.sql)

`tuid` — tuid datatype. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tuid;

select uuidv7();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `stuid_generate()` is an extension function and returns `bytea`.
- `tuid_generate()` is an extension function and returns `uuid`.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
