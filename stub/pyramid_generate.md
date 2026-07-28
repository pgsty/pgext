## Usage

Sources:

- [Official upstream README](https://github.com/suhaskamath2712/pg_pyramid/blob/19156aa292ba7f73f8e3137f8ff424b8876a5ec9/README.md)
- [Official extension control file (pyramid_generate.control)](https://github.com/suhaskamath2712/pg_pyramid/blob/19156aa292ba7f73f8e3137f8ff424b8876a5ec9/pyramid_generate/pyramid_generate.control)
- [Official extension SQL (pyramid_generate--1.0.sql)](https://github.com/suhaskamath2712/pg_pyramid/blob/19156aa292ba7f73f8e3137f8ff424b8876a5ec9/pyramid_generate/pyramid_generate--1.0.sql)

`pyramid_generate` — Generate uniformly distributed random float8 vectors by dimension and size. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pyramid_generate;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pyramid_generate(dimension int4, size int8)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
