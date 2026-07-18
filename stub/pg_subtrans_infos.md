## Usage

Sources:

- [Pinned upstream README](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/README.md)
- [Version 1.0 SQL API](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/pg_subtrans_infos--1.0.sql)
- [Pinned transaction-status implementation](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/pg_subtrans_infos.c)
- [Pinned extension control file](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/pg_subtrans_infos.control)

pg_subtrans_infos version 1.0 reports internal status and parentage for a recent transaction ID. Its single function returns xid, status, immediate parent, top parent, subtransaction depth, and an optional commit timestamp. It is a low-level diagnostic helper, not a durable transaction-history catalog.

### Inspect recent lock transactions

```sql
CREATE EXTENSION pg_subtrans_infos;

SELECT *
FROM pg_subtrans_infos(txid_current());

SELECT l.pid, l.transactionid, s.*
FROM pg_locks AS l
CROSS JOIN LATERAL
  pg_subtrans_infos(l.transactionid::text::bigint) AS s
WHERE l.transactionid IS NOT NULL
ORDER BY l.pid, l.transactionid;
```

Status is committed, aborted, or in progress. commit_timestamp is populated only when track_commit_timestamp is enabled and the corresponding commit-timestamp record is still retained. An ID older than retained pg_xact state produces NULL details, while an ID in the future raises an error.

### Retention and compatibility boundaries

The function combines the supplied 32-bit transaction ID with the current epoch. It is suitable only for recent IDs whose pg_xact and pg_subtrans entries have not wrapped or been truncated. Parent information can disappear before other transaction evidence, and the README notes that top_parent_xid can be empty after a transaction is no longer in progress. Never use the result for forensic completeness, business logic, or recovery decisions.

The C module copies PostgreSQL backend logic and calls private subtransaction, commit-timestamp, transaction-status, shared-memory, and lock interfaces. Those APIs and invariants can change in every major release. The README reports tests on PostgreSQL 10 through 12.2, and the pinned code adds branches for 13; it provides no support for 14 or newer.

The catalog marks this extension abandoned. Compile it only against a supported historical server, use it briefly for diagnostics, and remove it afterward. On current PostgreSQL releases, prefer supported monitoring views and logs or port the code with upstream backend tests rather than loading the old binary.
