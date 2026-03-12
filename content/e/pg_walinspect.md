---
title: "pg_walinspect"
linkTitle: "pg_walinspect"
description: "functions to inspect contents of PostgreSQL Write-Ahead Log"
weight: 6940
categories: ["STAT"]
width: full
---

[**pg_walinspect**](https://www.postgresql.org/docs/current/pgwalinspect.html) : functions to inspect contents of PostgreSQL Write-Ahead Log


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6940** | {{< badge content="pg_walinspect" link="https://www.postgresql.org/docs/current/pgwalinspect.html" >}} | {{< ext "pg_walinspect" >}} | `1.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "basic_archive" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "wal2json" >}} {{< ext "basebackup_to_shell" >}} {{< ext "decoderbufs" >}} {{< ext "test_decoding" >}} {{< ext "pg_profile" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_walinspect;
```



## Usage

> [pg_walinspect: low-level WAL inspection](https://www.postgresql.org/docs/current/pgwalinspect.html)

pg_walinspect provides SQL functions to inspect write-ahead log (WAL) contents on a running server. Similar to `pg_waldump` but accessible via SQL.

### Functions

**`pg_get_wal_record_info(in_lsn pg_lsn)`** -- single WAL record details:

```sql
SELECT * FROM pg_get_wal_record_info('0/E419E28');

 start_lsn        | 0/E419E28
 end_lsn          | 0/E419E68
 prev_lsn         | 0/E419D78
 xid              | 0
 resource_manager | Heap2
 record_type      | VACUUM
 record_length    | 58
 main_data_length | 2
 fpi_length       | 0
 description      | nunused: 5, unused: [1, 2, 3, 4, 5]
 block_ref        | blkref #0: rel 1663/16385/1249 fork main blk 364
```

**`pg_get_wal_records_info(start_lsn, end_lsn)`** -- all records in an LSN range:

```sql
SELECT * FROM pg_get_wal_records_info('0/1E913618', '0/1E913740');
```

**`pg_get_wal_block_info(start_lsn, end_lsn, show_data)`** -- block references from WAL records:

```sql
SELECT * FROM pg_get_wal_block_info('0/1230278', '0/12302B8');
```

Returns per block reference: `start_lsn`, `end_lsn`, `block_id`, `reltablespace`, `reldatabase`, `relfilenode`, `relforknumber`, `relblocknumber`, `xid`, `fork_flags`, `block_data`, `block_fpi_data`, etc.

**`pg_get_wal_stats(start_lsn, end_lsn, per_record)`** -- aggregate WAL statistics:

```sql
SELECT * FROM pg_get_wal_stats('0/1E847D00', '0/1E84F500')
WHERE count > 0;
```

### Tips

- Use `FFFFFFFF/FFFFFFFF` as `end_lsn` to read up to the current server LSN
- If `in_lsn` is not at a record boundary, the next valid record is returned
- All functions use the server's current timeline ID

### Access

Restricted to superusers and members of `pg_read_server_files`.
