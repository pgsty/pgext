## Usage

Sources:

- [Official extension control file](https://github.com/MarkAntipin/pg_sparse_vector/blob/e03511c809bd0e083ff977f4ce182fb90d37a29f/sparse_vector.control)

`sparse_vector` — C data type and functions for sparse float vectors, normalization, dot product, and cosine similarity.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "sparse_vector";
SELECT extversion
FROM pg_extension
WHERE extname = 'sparse_vector';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
