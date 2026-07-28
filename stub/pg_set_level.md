## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_set_level/pg_set_level-0.0.1/README.md)
- [Official extension control file (pg_set_level.control)](https://api.pgxn.org/src/pg_set_level/pg_set_level-0.0.1/pg_set_level.control)
- [Official extension SQL (pg_set_level--0.0.1.sql)](https://api.pgxn.org/src/pg_set_level/pg_set_level-0.0.1/pg_set_level--0.0.1.sql)

`pg_set_level` — pg_set_level is a PostgreSQL extension which allows to customize the SET statement. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_set_level;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
