## Usage

Sources:

- [pg_xid README at the reviewed commit](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/README.md)
- [pg_xid 1.0 install SQL at the reviewed commit](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid--1.0.sql)
- [pg_xid C entry points at the reviewed commit](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid.c)

`pg_xid` 1.0 generates Mongo ObjectID-compatible 12-byte identifiers. `xid` returns the raw value as `bytea`; `xid_encoded` returns its 20-character encoded `text` form. The layout combines a seconds timestamp, machine identifier, process ID, and per-process counter.

```sql
CREATE EXTENSION pg_xid;

SELECT encode(xid(), 'hex');
SELECT xid_encoded();
```

The timestamp prefix makes the identifiers roughly ordered by creation time without a central generator.

### Caveats

- These identifiers are not secrets or cryptographic random tokens. Their layout exposes creation time and process/machine-derived components.
- The 24-bit counter allows 16,777,216 values per second per host/process before wrapping, according to upstream; generator limits still need workload testing.
- Upstream provides no release series or current PostgreSQL compatibility matrix. Build and concurrency-test the C extension on every target major and architecture.
- Store the raw and encoded forms consistently. Text ordering, binary ordering, and application serialization should be verified before using them as distributed sort keys.
