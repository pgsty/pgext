## Usage

Sources:

- [Official upstream README](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/README.md)
- [Official extension control file (manip.control)](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/manip/manip.control)
- [Official extension SQL (manip--0.1.sql)](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/manip/manip--0.1.sql)

`manip` — Various manipulation functions. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION manip;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_prepared_plan(stmt_name text)` is an extension function and returns `text`.
- `scan_table` is an extension procedure.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
