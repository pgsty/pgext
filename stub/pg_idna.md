## Usage

Sources:

- [Official upstream README](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/README.md)

`pg_idna` — WHATWG URL IDNA helpers for PostgreSQL

The reviewed catalog snapshot records version `0.1.1-docs`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_idna";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_idna';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
