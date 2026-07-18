## Usage

Sources:

- [Official extension control file](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/pgintegrity.control)
- [Official upstream documentation](https://github.com/liang7878/pgintegrity/blob/7f812066f20c5d3052708612ee823bdd1e2ec663/README.md)

`pgintegrity` — Trigger-based row integrity watermark checker for PostgreSQL tables

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `dblink`.

```sql
CREATE EXTENSION "pgintegrity";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgintegrity';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
