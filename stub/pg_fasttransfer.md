## Usage

Sources:

- [Official project or provider page](https://fasttransfer.arpe.io/)

`pg_fasttransfer` — Postgresql Extension for FastTransfer

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pgcrypto`.
The curated compatibility set is `15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_fasttransfer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fasttransfer';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
