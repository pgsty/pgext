## Usage

Sources:

- [PGXN 0.4.2](https://pgxn.org/dist/pg_living_assertions/0.4.2/)

`pg_living_assertions` stores SQL checks with their verdicts, verification times and replacement history. It is a pure SQL extension requiring no preload, packaged for PostgreSQL 14–18. Distribution 0.4.2 retains SQL extension version 0.4.1.

### Register and verify

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

Checks return exactly one row with a boolean verdict and an optional detail string. The STABLE evaluator rejects writes. Checks run when requested; this is not a SQL ASSERTION constraint evaluated on each data change.

### Results and history

`living_assertions.run_all()` runs registered checks. `living_assertions.state()` distinguishes holds, broken, unknown, erroring, unchecked, retired and unregistered. `living_assertions.stale()` finds old verdicts while distinguishing checks never run.

Only trusted administrators should register or edit check SQL: future callers execute that SQL with their own privileges. Definitions are superseded with a reason rather than rewritten silently; verdict age must be considered alongside success. The registry and its recorded results are included in database dumps.
