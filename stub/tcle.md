## Usage

Sources:

- [Official upstream documentation](https://github.com/julmon/tcle/blob/b709024ea638df07eff74205999f645471a8d3b7/README.md)

`tcle` — Experimental table access method for transparent AES-256-CBC encryption of selected column types.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "tcle";
SELECT extversion
FROM pg_extension
WHERE extname = 'tcle';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
