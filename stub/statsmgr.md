## Usage

Sources:

- [Official extension control file](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/statsmgr.control)
- [Official upstream documentation](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/README.statsmgr.md)
- [Official upstream source](https://codeberg.org/Data-Bene/StatsMgr)

`statsmgr` — C extension that snapshots PostgreSQL cumulative WAL, SLRU, checkpoint, archiver, I/O, and background-writer statistics in dynamic shared memory.

The reviewed catalog snapshot records version `0.1-alpha`, kind `standard`, and implementation language `C`.
The curated compatibility set is `17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "statsmgr";
SELECT extversion
FROM pg_extension
WHERE extname = 'statsmgr';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
