## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/count_nulls/count_nulls-0.9.7/README.md)
- [Official extension control file (count_nulls.control)](https://api.pgxn.org/src/count_nulls/count_nulls-0.9.7/count_nulls.control)
- [Official extension SQL (count_nulls--0.9.0--0.9.2.sql)](https://api.pgxn.org/src/count_nulls/count_nulls-0.9.7/sql/count_nulls--0.9.0--0.9.2.sql)

`count_nulls` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION count_nulls;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `not_null_count(argument json)` is an extension function and returns `int`.
- `not_null_count(argument jsonb)` is an extension function and returns `int`.
- `not_null_count(VARIADIC argument anyarray)` is an extension function and returns `int`.
- `not_null_count_trigger()` is an extension function and returns `trigger`.
- `null_count(argument json)` is an extension function and returns `int`.
- `null_count(argument jsonb)` is an extension function and returns `int`.
- `null_count(VARIADIC argument anyarray)` is an extension function and returns `int`.
- `null_count_trigger()` is an extension function and returns `trigger`.

### Requirements and Caveats

- The reviewed control file declares default version `0.9.6`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
