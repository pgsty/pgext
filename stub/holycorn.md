## Usage

Sources:

- [Official extension control file](https://github.com/franckverrot/holycorn/blob/6cdfa71caffa28ba3ca720f253c4d5e112d5fa1d/holycorn.control)
- [Official upstream documentation](https://github.com/franckverrot/holycorn/blob/6cdfa71caffa28ba3ca720f253c4d5e112d5fa1d/README.md)

`holycorn` — Ruby foreign data wrapper provider

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "holycorn";
SELECT extversion
FROM pg_extension
WHERE extname = 'holycorn';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
