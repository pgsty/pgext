## Usage

Sources:

- [Official extension control file](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/pgprolog.control)
- [Official upstream documentation](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/README.md)
- [Official Rust package manifest](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/Cargo.toml)

`pgprolog` — PostgreSQL procedural language handler for Prolog, embedding Scryer Prolog.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgprolog";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgprolog';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
