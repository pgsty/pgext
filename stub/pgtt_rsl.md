## Usage

Sources:

- [Official extension control file](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/pgtt_rsl.control)
- [Official upstream documentation](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/README.md)

`pgtt_rsl` — Provide Oracle/DB2-style global temporary tables using unlogged tables, row-level security, views, and a cleanup worker

The reviewed catalog snapshot records version `2.0.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgtt_rsl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgtt_rsl';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
