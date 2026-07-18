## Usage

Sources:

- [Official extension control file](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/hyperspace.control)

`hyperspace` — Four-dimensional point, box, and circle types with B-tree and SP-GiST/KD-tree operator classes.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "hyperspace";
SELECT extversion
FROM pg_extension
WHERE extname = 'hyperspace';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
