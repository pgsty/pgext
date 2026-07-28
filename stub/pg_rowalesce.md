## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_rowalesce/pg_rowalesce-0.1.12/README.md)
- [Official extension control file (pg_rowalesce.control)](https://api.pgxn.org/src/pg_rowalesce/pg_rowalesce-0.1.12/pg_rowalesce.control)
- [Official extension SQL (pg_rowalesce--0.1.0.sql)](https://api.pgxn.org/src/pg_rowalesce/pg_rowalesce-0.1.12/sql/pg_rowalesce--0.1.0.sql)

`pg_rowalesce` — The pg_rowalesce PostgreSQL extension its defining feature is the rowalesce() function. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_rowalesce;

select rowalesce('{"my_attr_1": 3, "my_attr_2": "b"}'::jsonb, null::my.type)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `insert_row(inout anyelement)` is an extension function and returns `anyelement`.
- `pg_rowalesce_meta_pgxn()` is an extension function and returns `jsonb`.
- `pg_rowalesce_readme()` is an extension function and returns `text`.
- `table_defaults(pg_class$ regclass, include_columns$ hstore = null)` is an extension function and returns `hstore`.
- `test__pg_rowalesce` is an extension procedure.
- `myrow` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.12`.
- Install the confirmed extension dependencies first: `hstore`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
