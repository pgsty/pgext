## Usage

Sources:

- [Official upstream documentation](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/README.md)

`pg_redispub` — Publish channel messages from PostgreSQL to a local Redis server

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_redispub";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_redispub';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
