## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/plv8/)
- [Official project or provider page](https://plv8.github.io/)
- [Official primary documentation](https://pgxn.org/dist/plv8/doc/plv8.html)

`plcoffee` — CoffeeScript procedural language for PostgreSQL, compiled and executed by the PL/v8 runtime

The reviewed catalog snapshot records version `3.1.10`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "plcoffee";
SELECT extversion
FROM pg_extension
WHERE extname = 'plcoffee';
```

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
