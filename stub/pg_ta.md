## Usage

Sources:

- [Official extension control file](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/pg_ta.control)
- [Official upstream documentation](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/README.md)

`pg_ta` — TA-Lib technical-analysis functions exposed as PostgreSQL functions

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_ta";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ta';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
