## Usage

Sources:

- [Upstream README](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/README.md)
- [Version 0.1 SQL definition](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/pg_txn_status--0.1.sql)
- [Type implementation](https://github.com/zilder/pg_txn_status/blob/c1d49147b72a8c3230349d99a819b72a66881cde/pg_txn_status.c)

`pg_txn_status` version `0.1` provides a one-byte `txn_status` type for storing one of six application-defined transaction-state labels. It does not inspect PostgreSQL transactions or WAL.

### Core Workflow

```sql
CREATE EXTENSION pg_txn_status;

CREATE TABLE job_state (
  id bigint PRIMARY KEY,
  status txn_status NOT NULL
);

INSERT INTO job_state VALUES (1, 'begin'), (2, 'complete');
SELECT * FROM job_state WHERE status = 'complete'::txn_status;
```

Accepted, case-sensitive values are `begin`, `prepare`, `commit`, `rollback`, `complete`, and `incomplete`. Any other input raises an error.

### Objects and Caveats

The extension defines input/output and binary send/receive functions plus equality `=`. It does not define inequality, ordering, hashing, a B-tree class, or a transaction-management API.

Use a native PostgreSQL enum or a reference table when schema evolution, ordering, indexes, or maintained compatibility matter more than the one-byte representation. The upstream project is abandoned and uses old C extension interfaces; validate compilation, dump/restore, replication, and on-disk compatibility on the exact PostgreSQL release before storing durable data.
