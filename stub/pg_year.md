## Usage

Sources:

- [Official upstream README](https://github.com/pierreforstmann/pg_year/blob/8e18cd6bd7872807b6e13e22741b23c6f60a5b54/README.md)
- [Official extension control file (pg_year.control)](https://github.com/pierreforstmann/pg_year/blob/8e18cd6bd7872807b6e13e22741b23c6f60a5b54/pg_year.control)
- [Official extension SQL (pg_year--0.0.1.sql)](https://github.com/pierreforstmann/pg_year/blob/8e18cd6bd7872807b6e13e22741b23c6f60a5b54/pg_year--0.0.1.sql)

`pg_year` — This extension has been validated with PostgreSQL 12, 13, 14, 15, 16, 17 and 18. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_year;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hash_year(year)` is an extension function and returns `integer`.
- `year_add(year, int)` is an extension function and returns `year`.
- `year_cmp(year, year)` is an extension function and returns `integer`.
- `year_eq(year, year)` is an extension function and returns `boolean`.
- `year_ge(year, year)` is an extension function and returns `boolean`.
- `year_gt(year, year)` is an extension function and returns `boolean`.
- `year_in(cstring)` is an extension function and returns `year`.
- `year_le(year, year)` is an extension function and returns `boolean`.
- `year_lt(year, year)` is an extension function and returns `boolean`.
- `year_minus(year, int)` is an extension function and returns `year`.
- `year_ne(year, year)` is an extension function and returns `boolean`.
- `year_out(year)` is an extension function and returns `cstring`.
- `year` is an extension-defined type.
- `btree_year_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
