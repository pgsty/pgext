## Usage

Sources:

- [Official extension control file](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/array_textsort.control)
- [Official upstream documentation](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/README.md)

`array_textsort` — C functions for sorting and de-duplicating one-dimensional text arrays.

The reviewed catalog snapshot records version `1.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "array_textsort";
SELECT extversion
FROM pg_extension
WHERE extname = 'array_textsort';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
