## Usage

Sources:

- [Official upstream README](https://github.com/docteurklein/fakeset/blob/65675f9111d39da4eda4b780405cf7e8e8ebdf93/README.md)
- [Official extension control file (fakeset.control)](https://github.com/docteurklein/fakeset/blob/65675f9111d39da4eda4b780405cf7e8e8ebdf93/fakeset.control)
- [Official implementation source](https://github.com/docteurklein/fakeset/blob/65675f9111d39da4eda4b780405cf7e8e8ebdf93/src/lib.rs)

`fakeset` — A postgres extension to generate fake data. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION fakeset;

select lorem(3, 10) from generate_series(1, 2000);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lorem` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
