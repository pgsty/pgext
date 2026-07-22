## Usage

Sources:

- [Official README](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/README.md)
- [Version 0.0.1 SQL API](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/kv_fdw--0.0.1.sql)
- [Extension control file](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/kv_fdw.control)

`kv_fdw` embeds RocksDB behind a PostgreSQL foreign table, avoiding a separate remote database server. It supports basic reads and writes, but the upstream project describes version `0.0.1` as PostgreSQL 13-only and not product-ready.

### Core Workflow

Install a compatible RocksDB library, add `kv_fdw` to `shared_preload_libraries`, and restart PostgreSQL before creating the extension and server.

```sql
CREATE EXTENSION kv_fdw;
CREATE SERVER kv_server FOREIGN DATA WRAPPER kv_fdw;
CREATE FOREIGN TABLE student (
    id integer,
    name text
) SERVER kv_server;

INSERT INTO student VALUES (20757123, 'Rafferty');
SELECT * FROM student WHERE id = 20757123;
```

The first foreign-table attribute is the key. To model a composite key, upstream suggests using a composite PostgreSQL type as that first attribute. The wrapper does not take server options in the documented workflow.

### Limitations and Recovery

`kv_fdw` has no rollback or abort support. Once a foreign table is created, the documented implementation cannot add or drop columns, and it provides no secondary indexes. These constraints make ordinary PostgreSQL transaction expectations unsafe: rehearse failures and crash recovery, avoid mixing critical relational changes with RocksDB writes, and keep an independent backup strategy.

Dropping a foreign table does not imply that its on-disk RocksDB data has been removed. Follow the upstream removal procedure and inspect `data_directory` deliberately. Pin RocksDB `6.11.4`, PostgreSQL 13, and the extension commit used for the reviewed build; ABI or storage-format drift can prevent startup or data access.
