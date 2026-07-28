## Usage

Sources:

- [Official upstream README](https://github.com/willibrandon/pgl_validate/blob/fff68897716dc1d089719b0366039c6c79df3481/README.md)
- [Official extension control file (pgl_validate.control)](https://github.com/willibrandon/pgl_validate/blob/fff68897716dc1d089719b0366039c6c79df3481/pgl_validate.control)
- [Official implementation source](https://github.com/willibrandon/pgl_validate/blob/fff68897716dc1d089719b0366039c6c79df3481/src/lib.rs)

`pgl_validate` — pgl_validate is a PostgreSQL extension for validating table contents across pglogical and PostgreSQL logical-replication topologies. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgl_validate;

SELECT *
FROM pgl_validate.compare_table('public.accounts'::regclass);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
