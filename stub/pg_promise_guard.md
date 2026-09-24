## Usage

Sources:

- [PGXN 0.2.0 README](https://pgxn.org/dist/pg_promise_guard/0.2.0/README.html)
- [pg_promise_guard control file](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.0/pg_promise_guard.control)
- [pg_promise_guard 0.2.0 SQL definitions](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.0/pg_promise_guard--0.2.0.sql)
- [PostgreSQL license](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.0/LICENSE)

`pg_promise_guard` reports schema guarantees that look present but are wholly or partly unenforced. It reads PostgreSQL catalogs only, making it suitable for a periodic integrity check after migrations, bulk loads, or operational overrides.

### Core Workflow

```sql
CREATE EXTENSION pg_living_assertions;
CREATE EXTENSION pg_promise_guard;

SELECT * FROM promise_guard.promise_breaks;
SELECT promise_guard.promises_kept();
SELECT promise_guard.promises_kept('app');
SELECT * FROM promise_guard.check_promises('app');
```

`promise_guard.promise_breaks` is the database-wide view. `promise_guard.check_promises(text)` optionally limits results to one schema, while `promise_guard.promises_kept(text)` returns false only when at least one current breach is present.

### Findings

- An invalid unique index is a `breach` because uniqueness is not being enforced.
- A disabled user trigger is a `breach` because its declared behavior is inactive.
- Row-level security enabled but not forced is a `breach` because the table owner can bypass policies.
- A `NOT VALID` check or foreign-key constraint is a `gap`: new rows are checked, but existing rows were not validated.

The result type `promise_break` includes the finding kind, object, relation, severity, and explanation. Extension-owned objects are skipped so the check focuses on application schema state.

### Boundaries

The check reads catalogs without scanning user tables or taking application locks. It reports current state, not who or when changed it. It does not verify physical index integrity, permissions, default privileges, `search_path` shadowing, or historical enforcement; use tools such as `amcheck` and a separate security audit for those questions.

PGXN metadata claims PostgreSQL 13 or newer, but upstream has only tested 18.6 and 19 beta. Validate PostgreSQL 13 through 17 locally before relying on the result. The extension is pure SQL, needs no preload or shared library, and requires `pg_living_assertions`.

### Persistent Scan Status

Version 0.2.0 requires `pg_living_assertions` and adds `promise_guard.watch`, recording scan results and their age through living assertions. Only breaches make the verdict false; gaps remain distinct. The catalog scan itself stays stateless. This release also fixes internal name resolution so calls work with an independent caller search path.

```sql
SELECT promise_guard.watch('public');
SELECT * FROM living_assertions.status;
```
