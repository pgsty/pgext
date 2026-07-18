## Usage

Sources:

- [Official extension control file](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres.control)
- [Official upstream documentation](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/README.md)

`pargres` — Prototype parallel query execution module for shared-nothing PostgreSQL clusters

The reviewed catalog snapshot records version `0.1`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pargres";
SELECT extversion
FROM pg_extension
WHERE extname = 'pargres';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
