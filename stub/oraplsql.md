## Usage

Sources:

- [Official upstream README](https://github.com/opentenbase/opentenbase/blob/b612d77cbfd4d762f20c54c35f7caf09d57ef098/README.md)
- [Official extension control file (oraplsql.control)](https://github.com/opentenbase/opentenbase/blob/b612d77cbfd4d762f20c54c35f7caf09d57ef098/src/pl/oraplsql/src/oraplsql.control)
- [Official extension SQL (oraplsql--1.0.sql)](https://github.com/opentenbase/opentenbase/blob/b612d77cbfd4d762f20c54c35f7caf09d57ef098/src/pl/oraplsql/src/oraplsql--1.0.sql)

`oraplsql` — Oracle-compatible procedural SQL language for OpenTenBase. Use it when porting or emulating the corresponding database API. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION oraplsql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
