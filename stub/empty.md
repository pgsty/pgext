## Usage

Sources:

- [Official extension control file](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty.control)
- [Official upstream documentation](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/README.md)

`empty` — Demo C extension with sample functions, a matrix type, an FDW, logical-decoding callbacks, and shared-memory hooks.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "empty";
SELECT extversion
FROM pg_extension
WHERE extname = 'empty';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
