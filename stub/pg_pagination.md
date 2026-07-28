## Usage

Sources:

- [Official upstream README](https://github.com/arunsahoo-xt/pg_extensions/blob/63fe9163fe6864fbfd020c1e72507613780120c7/pg_pagination/README.md)
- [Official extension control file (pg_pagination.control)](https://github.com/arunsahoo-xt/pg_extensions/blob/63fe9163fe6864fbfd020c1e72507613780120c7/pg_pagination/pg_pagination.control)
- [Official implementation source](https://github.com/arunsahoo-xt/pg_extensions/blob/63fe9163fe6864fbfd020c1e72507613780120c7/pg_pagination/src/lib.rs)

`pg_pagination` — A **high-performance PostgreSQL extension** written in **Rust** using pgrx. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_pagination;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `paginate_by_cursor` is an extension function.
- `paginate_by_offset` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
