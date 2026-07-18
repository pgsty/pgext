## Usage

Sources:

- [Official extension control file](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/pgsynck.control)
- [Official upstream documentation](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/README.md)

`pgsynck` — Parse SQL text without execution and return per-statement syntax diagnostics

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgsynck";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgsynck';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
