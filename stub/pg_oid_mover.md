## Usage

Sources:

- [Official source](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/README.md)
- [Official source](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/pg_oid_mover.control)
- [Official source](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/pg_oid_mover--1.0.sql)
- [Official source](https://github.com/e-buzyurkin/pg_oid_mover/blob/f5a02b2162fa356dff49312980d1b60ee29d2496/pg_oid_mover.c)

pg_oid_mover is a test extension for disposable clusters. It advances PostgreSQL's cluster-wide OID allocator to exercise high-OID behavior. The change is not rolled back with the SQL transaction and must not be used on production clusters.

### Test interface

```sql
CREATE EXTENSION pg_oid_mover;
SELECT oid_mover.advance_to(3000000000::bigint);
```

Run only as a superuser after selecting a target appropriate for the disposable test. advance_to(bigint) returns the previous next-OID value; it accepts only forward movement in the valid OID range. The allocator is shared by every database. Taking and rolling back a transaction does not undo the advance.

### Boundaries

The nonrelocatable extension installs into oid_mover and revokes PUBLIC execution. It rejects recovery, read-only transactions and parallel execution. No preload is required. Rebuild against the exact server ABI; the source does not establish a general supported-major matrix or a license.
