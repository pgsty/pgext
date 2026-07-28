## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_sample_ext/pg_sample_ext-1.0.2/README.md)
- [Official extension control file (pg_sample_ext.control)](https://api.pgxn.org/src/pg_sample_ext/pg_sample_ext-1.0.2/pg_sample_ext.control)
- [Official extension SQL (pg_sample_ext--1.0.0.sql)](https://api.pgxn.org/src/pg_sample_ext/pg_sample_ext-1.0.2/pg_sample_ext--1.0.0.sql)

`pg_sample_ext` — pg_sample_ext is a PostgreSQL extension that provides a sample function to demonstrate extending PostgreSQL's functionality. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_sample_ext;

SELECT square(5);  -- Returns 25
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `square(num integer)` is an extension function and returns `integer`.
- `person_type` is an extension-defined type.
- `status_type` is an extension-defined type.
- `positive_integer` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.2`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
