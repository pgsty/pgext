## Usage

Sources:

- [Official extension control file](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal.control)
- [Official upstream documentation](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/README.md)

`fixeddecimal` — Fixed-precision decimal type optimized for performance and storage.

The reviewed catalog snapshot records version `1.1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "fixeddecimal";
SELECT extversion
FROM pg_extension
WHERE extname = 'fixeddecimal';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
