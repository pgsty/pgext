## Usage

Sources:

- [Official extension control file (extension_drop.control)](https://api.pgxn.org/src/extension_drop/extension_drop-0.1.1/extension_drop.control)
- [Official extension SQL (extension_drop.sql)](https://api.pgxn.org/src/extension_drop/extension_drop-0.1.1/sql/extension_drop.sql)

`extension_drop` — Run custom commands when an extension is dropped. Use it when administering or automating the database behavior described above. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION extension_drop;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `__extension_drop.create_function(function_name text , args text , options text , body text , comment text , grants text DEFAULT NULL)` is an extension function and returns `void`.
- `__extension_drop.exec(sql text)` is an extension function and returns `void`.
- `__extension_drop.safe_dump(relation regclass , filter text DEFAULT '')` is an extension function and returns `void`.
- `__extension_drop.messages` is a table installed or managed by the extension.
- `extension_drop__commands` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.1`.
- Install the confirmed extension dependencies first: `cat_tools`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
