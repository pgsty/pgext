## Usage

Sources:

- [Upstream README](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/README.md)
- [Extension control file](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree.control)
- [SQL installation script](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree--1.0.sql)
- [PostgreSQL access-method implementation](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree_fdw.cc)

`bztree` version `1.0` is an experimental, shared-memory BzTree index access method. It provides equality operator classes for selected integer, text, binary, network, and identifier types.

### Example

Add `bztree` to `shared_preload_libraries`, restart PostgreSQL, then run:

```sql
CREATE EXTENSION bztree;
CREATE INDEX events_key_bz ON events USING bztree (event_key);
EXPLAIN SELECT * FROM events WHERE event_key = 42;
```

This access method supports equality lookup but declares no ordering, uniqueness, backward scan, included columns, or parallel scan. The reviewed source contains unfinished memory-reclamation/deletion work and does not document durable recovery or current PostgreSQL compatibility. Do not substitute it for a production B-tree without crash/restart, vacuum, concurrency, and corruption testing on disposable data.
