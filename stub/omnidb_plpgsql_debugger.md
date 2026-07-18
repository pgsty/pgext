## Usage

Sources:

- [Official extension control file](https://github.com/OmniDB/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/omnidb_plpgsql_debugger.control)
- [Official upstream documentation](https://github.com/OmniDB/plpgsql_debugger/blob/223f9150306e00ef7706119a71f0c52385125a99/README.md)

`omnidb_plpgsql_debugger` — PostgreSQL extension for enabling PL/pgSQL debugger in OmniDB

The reviewed catalog snapshot records version `1.0.0`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "omnidb_plpgsql_debugger";
SELECT extversion
FROM pg_extension
WHERE extname = 'omnidb_plpgsql_debugger';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
