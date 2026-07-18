## Usage

Sources:

- [Official extension control file](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/bktree.control)
- [Official upstream documentation](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/README.md)

`bktree` — SP-GiST BK-tree index extension for 64-bit perceptual hashes and Hamming-distance searches.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "bktree";
SELECT extversion
FROM pg_extension
WHERE extname = 'bktree';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
