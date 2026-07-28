## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/test_factory/test_factory-0.4.2/README.md)
- [Official extension control file (test_factory_pgtap.control)](https://api.pgxn.org/src/test_factory/test_factory-0.4.2/test_factory_pgtap.control)
- [Official extension SQL (test_factory_pgtap.sql)](https://api.pgxn.org/src/test_factory/test_factory-0.4.2/sql/test_factory_pgtap.sql)

`test_factory_pgtap` — A system for managing unit test data in Postgres. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION test_factory_pgtap;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tf.tap(table_name text , set_name text DEFAULT 'base')` is an extension function and returns `SETOF text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `pgtap`, `test_factory`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
