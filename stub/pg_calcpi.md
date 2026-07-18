## Usage

Sources:

- [Version 1.0 PL/pgSQL implementation](https://github.com/itarient/pg_calcpi/blob/72012875fb47239b69b0347cde5de272b477ef52/pg_calcpi--1.0.sql)
- [Extension control file](https://github.com/itarient/pg_calcpi/blob/72012875fb47239b69b0347cde5de272b477ef52/pg_calcpi.control)
- [Upstream test query](https://github.com/itarient/pg_calcpi/blob/72012875fb47239b69b0347cde5de272b477ef52/test/find_subseq8_before_subseq8.sql)

`pg_calcpi` implements a PL/pgSQL spigot algorithm for digits of pi. `calcpi(n)` returns a set of the composite type `digit`, with a one-based sequence number and one decimal digit per row.

```sql
CREATE EXTENSION pg_calcpi;
SELECT seq, val
FROM calcpi(20)
ORDER BY seq;
```

Each invocation creates and repeatedly updates a session-local working table, so runtime grows sharply with the requested digit count. This is educational code, not a numeric library or efficient workload for a shared database server.

The function first executes an unqualified `DROP TABLE IF EXISTS calcpi_internal_data` before creating its temporary table. If a caller can resolve a persistent table with that name and owns it, the call can drop that table. Use only in an isolated schema/database with a controlled `search_path`, and do not grant execution to untrusted roles.
