## Usage

Sources:

- [Last-known official upstream repository; currently unavailable](https://github.com/dmtkfs/pg-tamper-log)

Current-source status: the official repository and owner return HTTP 404. Checks of GitHub repository, fork, code, and commit search, raw and codeload paths, PGXN, and Software Heritage found no current official source or archive.

The frozen catalog historically described `pg_tamperlog` version `1.1` as a pure-SQL tamper-evident audit log with hash chaining and dependencies `pg_tamperlog_rust` and `pgcrypto`. Those details are historical clues only and could not be revalidated after deletion.

```sql
-- Run only if matching, independently verified artifacts are recovered.
CREATE EXTENSION "pg_tamperlog";
```

Do not obtain replacement binaries or SQL from an unverified fork. Before any use, recover a trustworthy signed or content-addressed source, inspect every installed object and privilege, confirm the hash-chain threat model, and independently test deletion, update, concurrency, key management, backup, restore, and verification behavior.
