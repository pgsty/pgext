## Usage

Sources:

- [Official extension control file](https://github.com/funbringer/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/pg_linegazer.control)
- [Official upstream documentation](https://github.com/funbringer/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/README.md)

`pg_linegazer` — Transparent code coverage for PL/pgSQL

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "pg_linegazer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_linegazer';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
