## Usage

Sources:

- [Official README](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/README.md)
- [Version 0.0.1 extension SQL](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/pg_frozen--0.0.1.sql)
- [C implementation](https://github.com/lawrencejones/pg_frozen/blob/68004b4f21b01041487e2a080658565dc4cdc937/pg_frozen.c)

`pg_frozen` exposes a low-level diagnostic that inspects a heap tuple header and reports whether its inserting transaction ID has been frozen. It is useful for narrow VACUUM experiments, not as a durable tuple property or a substitute for PostgreSQL's vacuum statistics.

### Core Workflow

```sql
CREATE EXTENSION pg_frozen;

CREATE TABLE messages (id bigint GENERATED ALWAYS AS IDENTITY, msg text);
INSERT INTO messages (msg) VALUES ('hello');

SELECT id, ctid, xmin, frozen(tableoid, ctid)
FROM messages;

VACUUM (FREEZE) messages;

SELECT id, ctid, xmin, frozen(tableoid, ctid)
FROM messages;
```

A visible tuple normally reports `0` before freezing and `1` after its header is marked frozen.

### Object Index

- `frozen(tableoid oid, tid tid) RETURNS integer` fetches the physical tuple identified by relation OID and `ctid`; it returns the header's frozen flag, or `3` when the tuple cannot be fetched.

### Operational Notes

Version `0.0.1` is relocatable and declares no dependency or preload requirement. The function is `STRICT` and `PARALLEL SAFE`, but it reads physical heap state through PostgreSQL internals.

Always derive `tableoid` and `ctid` from the same current row in one statement. A `ctid` is not a stable identifier: UPDATE, VACUUM, pruning, and concurrent activity can replace or remove the referenced tuple. The extension uses `SnapshotAny`, so its result is an implementation-level observation that can include tuples outside normal MVCC visibility. Keep it out of application logic and test it against the exact PostgreSQL major in use.
