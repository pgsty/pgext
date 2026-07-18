## Usage

Sources:

- [Official extension control file](https://github.com/blkchain/pg_blkchain/blob/master/pg_blkchain.control)
- [Official upstream documentation](https://github.com/blkchain/pg_blkchain/blob/master/README.md)

`pg_blkchain` — C extension providing Bitcoin blockchain parsing and helper functions for PostgreSQL.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_blkchain";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_blkchain';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
