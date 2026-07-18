## Usage

Sources:

- [Official extension control file](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/pgunicoll.control)
- [Official upstream documentation](https://github.com/madebykite/pgunicoll/blob/af4e0c71fa3fe109ecca9acb2b254af1d545c2c8/README.md)

`pgunicoll` — Generate ICU Unicode Collation Algorithm sort keys for PostgreSQL text values.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgunicoll";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgunicoll';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
