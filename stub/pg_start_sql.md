## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_start_sql/pg_start_sql-0.0.2/README.md)
- [Official extension control file (pg_start_sql.control)](https://api.pgxn.org/src/pg_start_sql/pg_start_sql-0.0.2/pg_start_sql.control)

`pg_start_sql` — PostgreSQL extension to execute SQL statements at instance start. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_start_sql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
