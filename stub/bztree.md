## Usage

Sources:

- [Upstream README](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/README.md)
- [Extension control file](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree.control)
- [SQL installation script](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree--1.0.sql)
- [PostgreSQL access-method implementation](https://github.com/postgrespro/bztree/blob/a4b238720b742a336d637702ed5985b17ea31af1/bztree_fdw.cc)

`bztree` version `1.0` is an experimental, shared-memory BzTree index access method. It provides equality operator classes for selected integer, text, binary, network, and identifier types.

### Disposable demonstration

Add `bztree` to `shared_preload_libraries`, restart PostgreSQL, and create the index while the demonstration table is still empty. Rows inserted after the index exists go through the implemented insertion callback:

```sql
CREATE EXTENSION bztree;
CREATE TABLE bztree_demo (event_key integer NOT NULL, payload text);
CREATE INDEX bztree_demo_key_bz ON bztree_demo USING bztree (event_key);
INSERT INTO bztree_demo VALUES (42, 'after index creation');
EXPLAIN SELECT * FROM bztree_demo WHERE event_key = 42;
```

Do not build this index over existing rows. In the reviewed source, `bztree_build()` scans the heap but its build callback only serializes each key and never inserts it into the BzTree; it nevertheless reports all scanned rows as index tuples. A planner-selected scan of such an index can therefore omit existing rows and return incorrect results. NULL keys are also rejected.

This access method supports equality lookup but declares no ordering, uniqueness, backward scan, included columns, or parallel scan. The source contains unfinished memory-reclamation/deletion work and does not document durable recovery or current PostgreSQL compatibility. Treat it only as source-level research on disposable data; do not rely on it for correct query results or substitute it for a production B-tree.
