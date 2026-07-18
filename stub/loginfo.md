## Usage

Sources:

- [Official extension control file](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/loginfo.control)
- [Official upstream documentation](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/README.md)

`loginfo` — Functions and views for PostgreSQL to query the postgres log

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "loginfo";
SELECT extversion
FROM pg_extension
WHERE extname = 'loginfo';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
