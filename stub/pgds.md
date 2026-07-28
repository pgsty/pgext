## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pgds/pgds-0.0.3/README.md)
- [Official extension control file (pgds.control)](https://api.pgxn.org/src/pgds/pgds-0.0.3/pgds.control)
- [Official extension SQL (pgds--0.0.3.sql)](https://api.pgxn.org/src/pgds/pgds-0.0.3/pgds--0.0.3.sql)

`pgds` — PostgreSQL extension to gather dynamic statistics. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgds;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `find_tables(p_oid oid)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
