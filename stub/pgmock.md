## Usage

Sources:

- [Project README and examples](https://github.com/avito-tech/pgmock/blob/f786e6f835771c65b2721fd7f275a673d627193d/README.md)
- [Extension control file](https://github.com/avito-tech/pgmock/blob/f786e6f835771c65b2721fd7f275a673d627193d/pgmock.control)
- [Version 0.2 SQL implementation](https://github.com/avito-tech/pgmock/blob/f786e6f835771c65b2721fd7f275a673d627193d/pgmock--0.2.sql)
- [Regression tests](https://github.com/avito-tech/pgmock/tree/f786e6f835771c65b2721fd7f275a673d627193d/sql)

`pgmock` 0.2 is a pure-SQL unit-testing extension originally documented for PostgreSQL 9.4 and later. Its `mock()` function recreates selected functions and table dependencies in `pg_temp`, then exposes `pg_temp.setup()` and `pg_temp.teardown()` to reset test state.

### Isolate a function test

Install it in a dedicated schema to avoid name conflicts, and wrap each suite in a transaction:

```sql
CREATE SCHEMA pgmock;
CREATE EXTENSION pgmock WITH SCHEMA pgmock;

BEGIN;

SELECT pgmock.mock($$"${'public.universal_answer'::regproc}"$$);
SELECT pg_temp.setup();
SELECT pg_temp.universal_answer();
SELECT pg_temp.teardown();

ROLLBACK;
```

The JSON request can identify `regproc` and `regclass` dependencies and selectively reproduce constraints, defaults, constant functions, and triggers. Only requested behavior is reconstructed; a temporary table is not automatically a complete semantic copy of its source.

### Test-only security and compatibility

`mock()` reads catalogs and generates SQL that recreates objects under a temporary search path. Accept definitions only from trusted test fixtures, use an isolated nonproduction database and role, and retain the outer rollback. Never run unreviewed mock requests against a production session.

The 0.2 implementation references old catalog fields including `pg_attrdef.adsrc`, which was removed from modern PostgreSQL. The repository has no current compatibility matrix or release activity, so the original PostgreSQL 9.4+ statement is not evidence of support on current majors. Test installation and every requested object shape on the target major. Pay particular attention to overloaded functions, identity and generated columns, partitioned tables, policies, security-definer functions, triggers, and search-path resolution.
