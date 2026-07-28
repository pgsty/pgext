## Usage

Sources:

- [Official upstream README](https://github.com/zombodb/postgresconf/blob/f72a1f01f408b19076d20e182568908ee1f6f873/README.md)
- [Official extension control file (postgresconf.control)](https://github.com/zombodb/postgresconf/blob/f72a1f01f408b19076d20e182568908ee1f6f873/postgresconf.control)
- [Official implementation source](https://github.com/zombodb/postgresconf/blob/f72a1f01f408b19076d20e182568908ee1f6f873/src/lib.rs)

`postgresconf` — First off, you need the following software to build this extension:. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION postgresconf;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `array_of_names()` is an extension function.
- `array_of_names_with_null()` is an extension function.
- `hello_postgresconf()` is an extension function.
- `my_generate_series` is an extension function.
- `rust_tuple` is an extension function.
- `set_of_animals()` is an extension function.
- `sum_array` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
