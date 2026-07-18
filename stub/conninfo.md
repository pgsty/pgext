## Usage

Sources:

- [Official extension control file](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/conninfo.control)
- [Official upstream documentation](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/README.md)

`conninfo` — Parse, validate, and inspect default values for PostgreSQL libpq conninfo strings

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "conninfo";
SELECT extversion
FROM pg_extension
WHERE extname = 'conninfo';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
