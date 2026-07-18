## Usage

Sources:

- [Official upstream documentation](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/README.md)

`ora_migrator` — Tools to migrate Oracle databases to PostgreSQL

The reviewed catalog snapshot records version `1.1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `db_migrator`, `oracle_fdw`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ora_migrator";
SELECT extversion
FROM pg_extension
WHERE extname = 'ora_migrator';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
