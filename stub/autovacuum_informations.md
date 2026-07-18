## Usage

Sources:

- [Official extension control file](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/autovacuum_informations.control)
- [Official upstream documentation](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/README.md)

`autovacuum_informations` — PostgreSQL extension that adds functions for reporting the autovacuum launcher PID and autovacuum worker information.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "autovacuum_informations";
SELECT extversion
FROM pg_extension
WHERE extname = 'autovacuum_informations';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
