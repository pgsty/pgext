## Usage

Sources:

- [Official upstream documentation](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/README.md)
- [Official extension SQL](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/pg_dirty_hands--1.0.sql)
- [Official implementation](https://github.com/dsarafan/pg_dirty_hands/blob/3d5bb01c39e987bc394fc0f093c37d673744ba7e/src/pg_dirty_hands.c)

`pg_dirty_hands` exposes low-level functions that rewrite a single heap tuple's visibility metadata. It copies PostgreSQL 9.6-era vacuum/freezing internals and deliberately bypasses normal MVCC safeguards. Use it only for expert forensic or recovery experiments on a disposable copy; an incorrect relation or stale CTID can corrupt data, indexes, backups, and replicas.

### Lock Down Installation

The C code performs no independent superuser check, while extension functions can be executable by `PUBLIC`. Revoke access immediately after installation and grant it only to a dedicated emergency role:

```sql
CREATE EXTENSION pg_dirty_hands;

REVOKE EXECUTE ON FUNCTION freeze_tuple(regclass, tid, boolean) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION freeze_tuple_unlogged(regclass, tid) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION freeze_tuple(regclass, tid, boolean)
  TO emergency_dba;
```

Do not grant `freeze_tuple_unlogged(regclass, tid)` for normal use.

### Important Objects

- `freeze_tuple(regclass, tid, boolean)` freezes one tuple identified by its physical CTID. With the default `false`, it follows the source's limited eligibility checks; with `true`, it forcibly marks the insert transaction frozen and invalidates delete/update metadata.
- `freeze_tuple_unlogged(regclass, tid)` changes tuple headers without marking the buffer dirty and without WAL. Despite its name, the implementation does not verify that the relation is unlogged.

### Extreme-Risk Boundary

Upstream reports testing only PostgreSQL 9.6.3 through 9.6.6. The implementation is coupled to heap-page layout, buffer handling, WAL records, and transaction-header bits from that generation. Never assume binary or semantic compatibility with another PostgreSQL release. Stop writes, preserve a physical copy, verify the exact tuple and page with forensic tools, and rehearse recovery before any call. Never use the `true` force path or the unlogged function on a production primary. A boolean success result is not proof that logical visibility, indexes, replicas, or future vacuum operations remain valid.
