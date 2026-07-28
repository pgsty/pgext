## Usage

Sources:

- [Official upstream README](https://github.com/spa5k/slugify-postgres/blob/54c6a9eef1d30b3434e32e30404d65ae5f91a440/README.md)
- [Official extension control file (slugify.control)](https://github.com/spa5k/slugify-postgres/blob/54c6a9eef1d30b3434e32e30404d65ae5f91a440/slugify.control)
- [Official implementation source](https://github.com/spa5k/slugify-postgres/blob/54c6a9eef1d30b3434e32e30404d65ae5f91a440/src/lib.rs)

`slugify` — PostgreSQL Extension to generate various variant of Slugs from a string. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION slugify;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `slug` is an extension function.
- `slug_rand` is an extension function.
- `slug_rand_c` is an extension function.
- `slug_rand_sep` is an extension function.
- `slug_rand_sep_c` is an extension function.
- `slug_sep` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
