## Usage

Sources:

- [Official extension control file](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl.control)
- [Official upstream documentation](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/README.md)

`ddl_repl` — Replicate DDL and DCL commands across PostgreSQL instances by means of a utility hook

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "ddl_repl";
SELECT extversion
FROM pg_extension
WHERE extname = 'ddl_repl';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
