## Usage

Sources:

- [Official extension control file](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/pg_plsql_graphs.control)
- [Official upstream documentation](https://github.com/KLIEBHAN/pg_plsql_graphs/blob/22b2eaef35a90234aa364aee8bf21f600c639a0e/README.md)

`pg_plsql_graphs` — Capture PL/pgSQL control-flow and dependence graphs for DOT export

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_plsql_graphs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_plsql_graphs';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
