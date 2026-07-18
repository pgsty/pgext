## Usage

Sources:

- [Official extension control file](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper.control)

`pg_llm_helper` — Captures PostgreSQL errors in shared memory and exposes history plus an optional pgai-backed LLM troubleshooting function.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`, `vector`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_llm_helper";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_llm_helper';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
