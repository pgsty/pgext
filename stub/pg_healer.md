## Usage

Sources:

- [Official extension control file](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/pg_healer.control)
- [Official upstream README](https://github.com/turnstep/pg_healer/blob/d313cccc2944a3c4a54757c5fa2693e9ae21d637/README.md)

`pg_healer` — Proof-of-concept extension for automatically repairing PostgreSQL data corruption

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plperlu`.

```sql
CREATE EXTENSION "pg_healer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_healer';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
