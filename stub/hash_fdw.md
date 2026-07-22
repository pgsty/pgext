## Usage

Sources:

- [Official README](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/README)
- [Official extension control file](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw.control)
- [Official extension SQL](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw--1.0.sql)
- [Official shared-memory implementation](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_if.c)
- [Official FDW implementation](https://github.com/sangli00/hash_fdw/blob/9c685317f6a08381b426706c235ac515505cac51/hash_fdw.c)

`hash_fdw` 1.0 is a 2016 in-memory FDW experiment that stores foreign-table rows in PostgreSQL shared-memory hash tables. It can scan and insert keyed rows without an external server, but it is not durable storage and does not provide normal PostgreSQL table semantics.

### Preload and Basic Workflow

The library allocates shared memory in `_PG_init`, so add it to `shared_preload_libraries`, size the store, and restart before creating the extension.

```conf
shared_preload_libraries = 'hash_fdw'
hash.shmem_size = 1024
```

```sql
CREATE EXTENSION hash_fdw;

CREATE SERVER hash_server
FOREIGN DATA WRAPPER hash_fdw;

CREATE FOREIGN TABLE session_cache (
    id bigint,
    name text,
    age integer
)
SERVER hash_server
OPTIONS (key 'id', hash_idx '1');

INSERT INTO session_cache VALUES (1, 'Ada', 37);
SELECT * FROM session_cache;
```

`key` names the foreign-table column converted to the hash key. `hash_idx` selects one shared hash table; the public helper functions hard-limit usable indexes to 1–8 in this source.

### Installed Surface

- `hash_fdw_handler()` and `hash_fdw_validator(text[], oid)` implement the FDW.
- Foreign scans iterate an entire selected hash table; inserts add or replace entries addressed by the configured key.
- `hash_insert_hashdata(int, cstring, record)` and `hash_get_hashdata(int, cstring)` expose low-level C-backed helpers, but their `record` contract is awkward and tied to internal tuple representation.
- GUCs include `hash.shmem_size` and `hash_fdw.hash_idx`; shared-memory layout is fixed at server startup, so do not change sizing/count assumptions dynamically.

### Data and Safety Boundaries

Rows are not written to heap storage, WAL, or an external service. A restart, crash, failover, or binary incompatibility can lose or invalidate all entries. The code does not implement MVCC visibility, transactional rollback, constraints, indexes, backup, replication, or concurrency guarantees equivalent to a regular table.

The shared hash tables are cluster-wide, while foreign-table definitions are database-local. Reusing the same `hash_idx` and key across databases or incompatible row layouts can collide or interpret tuple bytes with the wrong descriptor. The delete callback is incomplete and no update path is registered. Use unique isolated indexes, disposable data, and a dedicated test instance only.

The README states it was built for PostgreSQL 9.5.3. Server shared-memory and FDW APIs have changed substantially; do not load this unreviewed binary on a modern server.
