## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.8/README.md)
- [pg_promise_guard.control](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.8/pg_promise_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_promise_guard/pg_promise_guard-0.2.8/CHANGELOG.md)

`pg_promise_guard` 0.2.8 scans system catalogs for declared guarantees that are no longer fully enforced. It requires `pg_living_assertions`, uses pure SQL and needs no preload.

### Core Workflow

```sql
CREATE EXTENSION pg_living_assertions;
CREATE EXTENSION pg_promise_guard;

SELECT * FROM promise_guard.promise_breaks;
SELECT promise_guard.promises_kept();
SELECT promise_guard.promises_kept('app');
SELECT * FROM promise_guard.check_promises('app');
```

```sql
SELECT promise_guard.watch('public');
SELECT * FROM living_assertions.status;
```

### Operational Boundaries

`promise_guard.check_promises()` returns breaches and gaps; `promise_guard.promises_kept()` is false only for breaches, and `promise_guard.watch()` records the verdict in living assertions. Only READ COMMITTED is accepted, and nonexistent schemas raise errors. Disabled triggers, unenforced constraints and inactive policies/rules can be breaches. Replica-only triggers, unvalidated history and a ready-but-invalid unique index can be gaps; failed indexes that accept duplicates remain breaches. Only objects of superuser-owned extensions are excluded. Old watches from before 0.2.7 should be retired and recreated to adopt catalog-first name resolution. The scan does not prove physical integrity or audit all privileges. PGXN claims PostgreSQL 13+, while documented tests cover 18.6 and 19 beta. Current Apache-2.0 licensing starts at 0.2.3.
