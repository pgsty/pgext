---
title: "pgrowlocks"
linkTitle: "pgrowlocks"
description: "show row-level locking information"
weight: 6910
categories: ["STAT"]
width: full
---

[**pgrowlocks**](https://www.postgresql.org/docs/current/pgrowlocks.html) : show row-level locking information


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6910** | {{< badge content="pgrowlocks" link="https://www.postgresql.org/docs/current/pgrowlocks.html" >}} | {{< ext "pgrowlocks" >}} | `1.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgrowlocks;
```



## Usage

> [pgrowlocks: display row-level locking information](https://www.postgresql.org/docs/current/pgrowlocks.html)

pgrowlocks shows which rows in a table are currently locked, by which transactions, and the lock modes.

### Function

```sql
SELECT * FROM pgrowlocks('my_table');

 locked_row | locker | multi | xids  |     modes      |  pids
------------+--------+-------+-------+----------------+--------
 (0,1)      |    609 | f     | {609} | {"For Share"}  | {3161}
 (0,2)      |    609 | f     | {609} | {"For Share"}  | {3161}
 (0,3)      |    607 | f     | {607} | {"For Update"} | {3107}
```

### Return Columns

| Column | Type | Description |
|--------|------|-------------|
| `locked_row` | tid | Tuple ID of the locked row |
| `locker` | xid | Transaction ID (or multixact ID) |
| `multi` | boolean | True if locker is a multitransaction |
| `xids` | xid[] | Transaction IDs of all lockers |
| `modes` | text[] | Lock modes: `For Key Share`, `For Share`, `For No Key Update`, `For Update`, etc. |
| `pids` | integer[] | Process IDs of locking backends |

### View Locked Row Contents

```sql
SELECT * FROM accounts AS a, pgrowlocks('accounts') AS p
WHERE p.locked_row = a.ctid;
```

### Access

Restricted to superusers, roles with `pg_stat_scan_tables`, and users with `SELECT` on the target table.

### Caveats

- Takes `AccessShareLock` on the target table
- Not guaranteed to produce a self-consistent snapshot
- Can be slow on large tables
