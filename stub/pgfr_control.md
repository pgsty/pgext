## Usage

Sources:

- [Official primary documentation](https://database.dev/dventimi/pgfr_control)

`pgfr_control` — Vacuum diagnostics, autovacuum scale-factor recommendations, and table bloat analysis for pgfr_record.

The reviewed catalog snapshot records version `2.28.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pgfr_record`.

```sql
CREATE EXTENSION "pgfr_control";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfr_control';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
