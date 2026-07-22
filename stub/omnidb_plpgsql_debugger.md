## Usage

Sources:

- [Official README](https://github.com/omnidb/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/README.md)
- [Extension SQL for 1.0.0](https://github.com/omnidb/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/omnidb_plpgsql_debugger--1.0.0.sql)
- [OmniDB debugger documentation](https://omnidb.readthedocs.io/en/latest/en/13_writing_and_debugging_plpgsql_functions.html)

`omnidb_plpgsql_debugger` is the PostgreSQL server component for debugging PL/pgSQL functions from the OmniDB GUI. It is not a standalone interactive debugger: the server module and OmniDB client work together, and debugging opens an additional local database connection.

### Core Workflow

Preload the module and restart PostgreSQL before creating the extension in each database to be debugged.

```conf
shared_preload_libraries = 'omnidb_plpgsql_debugger'
```

```sql
CREATE EXTENSION omnidb_plpgsql_debugger;

GRANT ALL ON SCHEMA omnidb TO debugger_user;
GRANT ALL ON ALL TABLES IN SCHEMA omnidb TO debugger_user;
```

Connect OmniDB as the debugging role, open a PL/pgSQL function in the function editor, set breakpoints, and start the debugger from the GUI.

### Objects and Connection Requirements

- `omnidb.omnidb_enable_debugger(character varying)` initializes the server-side debugging session.
- `omnidb.contexts`, `omnidb.variables`, and `omnidb.statistics` hold debugger state used by the client.

Creating the extension and granting debugger access are superuser tasks. The extra debugger connection must be allowed by `pg_hba.conf` for the same database and role. Upstream documents either a local trust rule or password authentication backed by `.pgpass`; prefer authenticated local access and restrict the rule as narrowly as possible.

Preloading requires a server restart. Granting all privileges in the `omnidb` schema exposes mutable debugger state, so grant it only to trusted developers and only in databases where interactive debugging is acceptable.
