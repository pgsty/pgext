## Usage

Sources:

- [Version 1.0 extension SQL](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx--1.0.sql)
- [Index access-method implementation](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx.c)
- [Example configuration](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx.conf)

`lsm_idx` is an incomplete proof-of-concept PostgreSQL index access method, despite its LSM-oriented name. Version 1.0 registers an access method and B-tree-like operator classes, but its C handler leaves all scan callbacks unimplemented. It is source-study material and must not be used for application indexes.

### Disposable Inspection Only

In an isolated server built for the same PostgreSQL headers as the source, the extension's catalog surface can be inspected without placing the index on important data:

```sql
CREATE EXTENSION lsm_idx;

SELECT amname, amtype, amhandler::regproc
FROM pg_am
WHERE amname = 'lsm_idx';

SELECT opcname, opcintype::regtype
FROM pg_opclass
WHERE opcmethod = (SELECT oid FROM pg_am WHERE amname = 'lsm_idx')
ORDER BY 1, 2;
```

The SQL defines default classes for common integer, floating-point, boolean, character, binary, date/time, interval, network-related, numeric, OID, text/name, money, and UUID types, mostly by referring to PostgreSQL's internal B-tree support functions.

The apparent creation syntax is:

```sql
CREATE INDEX experimental_lsm_idx
ON experimental_table USING lsm_idx (key_column);
```

Do not execute that example against valuable data. The handler delegates build and insert operations to private B-tree internals after temporarily changing the relation's access-method OID, writes extra counters into the B-tree metapage, and emits diagnostic output with `printf`.

### Missing Functionality and Risks

The access-method routine sets `ambeginscan`, `amrescan`, `amgettuple`, `amgetbitmap`, and `amendscan` to null. It also disables unique indexes, array searches, and parallel scans. There is no implemented LSM merge, level management, compaction, recovery protocol, cost model, or user-visible statistics interface.

The repository contains an example `shared_preload_libraries` entry and `lsm_idx.top_index_size`, but the C source defines no startup hook or GUC with that name. Those lines are not evidence of a working preload requirement or tunable option.

Version 1.0 is a single-commit 2020 prototype with no README, tests, upgrade path, or PostgreSQL compatibility matrix. The control file is relocatable and declares no dependency. Use PostgreSQL's supported B-tree access method or a maintained alternative for production; any attempt to continue this prototype requires a complete access-method implementation plus crash, WAL, vacuum, concurrency, planner, upgrade, and corruption testing.
