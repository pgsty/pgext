## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/postgrespro/lsm/blob/bbd87030ae68fd855d91b43ca9367968bdc2fa69/README.md)
- [Extension control file](https://github.com/postgrespro/lsm/blob/bbd87030ae68fd855d91b43ca9367968bdc2fa69/lsm.control)
- [FDW implementation](https://github.com/postgrespro/lsm/blob/bbd87030ae68fd855d91b43ca9367968bdc2fa69/lsm_fdw.cpp)

`lsm` embeds RocksDB in a PostgreSQL backend and exposes it through the `lsm_fdw` foreign-data wrapper. Build compatibility depends on both PostgreSQL and RocksDB; the reviewed upstream reports only Ubuntu 18.04, PostgreSQL 11, and RocksDB 6.2.4 testing. Add `lsm` to `shared_preload_libraries` and restart before use.

```sql
CREATE EXTENSION lsm;
CREATE SERVER lsm_server FOREIGN DATA WRAPPER lsm_fdw;
CREATE FOREIGN TABLE student (
  id integer,
  name text
) SERVER lsm_server;
```

The first foreign-table attribute acts as the primary key, and composite primary keys are unsupported. PostgreSQL data types are not stored natively, and upstream says ACID behavior relies on the storage engine.

RocksDB files are external state from PostgreSQL's perspective. Prove backup, restore, crash recovery, replication, transaction abort, concurrent access, and cleanup behavior before storing important data. The upstream removal example names the wrong extension, so use the installed catalog name and preserve the RocksDB directory until recovery requirements are settled.
