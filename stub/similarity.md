## Usage

Sources:

- [Official extension control file](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity.control)
- [Official upstream documentation](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/README.md)

`similarity` — C extension that computes a normalized similarity score between two strings using an fstrcmp-style approximate matching algorithm.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "similarity";
SELECT extversion
FROM pg_extension
WHERE extname = 'similarity';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
