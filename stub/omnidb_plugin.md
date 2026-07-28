## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/omnidb_plugin/omnidb_plugin-0.0.1/README.md)
- [Official extension control file (omnidb_plugin.control)](https://api.pgxn.org/src/omnidb_plugin/omnidb_plugin-0.0.1/omnidb_plugin.control)
- [Official extension SQL (omnidb_plugin--0.0.1.sql)](https://api.pgxn.org/src/omnidb_plugin/omnidb_plugin-0.0.1/omnidb_plugin--0.0.1.sql)

`omnidb_plugin` — nano /etc/postgresql/X.Y/main/postgresql.conf shared_preload_libraries = '/opt/omnidb-plugin/omnidb_plugin_XY'. Use it when database code must run in or interoperate with this procedural language. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION omnidb_plugin;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `omnidb.omnidb_enable_debugger(character varying)` is an extension function and returns `void`.
- `omnidb.contexts` is a table installed or managed by the extension.
- `omnidb.statistics` is a table installed or managed by the extension.
- `omnidb.variables` is a table installed or managed by the extension.
- `omnidb` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
