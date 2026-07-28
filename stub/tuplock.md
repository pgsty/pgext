## Usage

Sources:

- [Official extension control file (tuplock.control)](https://api.pgxn.org/src/tuplock/tuplock-1.2.2/tuplock.control)
- [Official extension SQL (tuplock.sql)](https://api.pgxn.org/src/tuplock/tuplock-1.2.2/sql/tuplock.sql)

`tuplock` — lock tuples (rows) with a boolean attribute. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tuplock;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tuplock()` is an extension function and returns `TRIGGER`.
- `test` is a table installed or managed by the extension.
- `test2` is a table installed or managed by the extension.
- `test_tuplock` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.2.2`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
