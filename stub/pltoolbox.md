## Usage

Sources:

- [Official extension control file](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/pltoolbox.control)
- [Official upstream documentation](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/README.md)

`pltoolbox` — provide C utility functions for PostgreSQL stored procedures

The reviewed catalog snapshot records version `1.0.3`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pltoolbox";
SELECT extversion
FROM pg_extension
WHERE extname = 'pltoolbox';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
