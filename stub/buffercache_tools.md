## Usage

Sources:

- [Official extension control file](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/buffercache_tools.control)
- [Official upstream README](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/README.md)

`buffercache_tools` — Inspect and manipulate PostgreSQL buffer cache entries for relations, forks, databases, and tablespaces.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "buffercache_tools";
SELECT extversion
FROM pg_extension
WHERE extname = 'buffercache_tools';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
