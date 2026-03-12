---
title: "pgpool_recovery"
linkTitle: "pgpool_recovery"
description: "recovery functions for pgpool-II for V4.3"
weight: 5910
categories: ["ADMIN"]
width: full
---

[**pgpool**](https://pgpool.net/) : recovery functions for pgpool-II for V4.3


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5910** | {{< badge content="pgpool_recovery" link="https://pgpool.net/" >}} | {{< ext "pgpool_recovery" "pgpool" >}} | `4.7.1` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgautofailover" >}} {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "repmgr" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |
|    **Siblings**   | {{< ext "pgpool_adm" >}} {{< ext "pgpool_regclass" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgpool` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7.1` | {{< bg "18" "pgpool-II-pg18-extensions" "green" >}} {{< bg "17" "pgpool-II-pg17-extensions" "green" >}} {{< bg "16" "pgpool-II-pg16-extensions" "green" >}} {{< bg "15" "pgpool-II-pg15-extensions" "green" >}} {{< bg "14" "pgpool-II-pg14-extensions" "green" >}} | `pgpool-II-pg$v-extensions` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7.1` | {{< bg "18" "postgresql-18-pgpool2" "green" >}} {{< bg "17" "postgresql-17-pgpool2" "green" >}} {{< bg "16" "postgresql-16-pgpool2" "green" >}} {{< bg "15" "postgresql-15-pgpool2" "green" >}} {{< bg "14" "postgresql-14-pgpool2" "green" >}} | `postgresql-$v-pgpool2` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg18-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg17-extensions : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg16-extensions : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg15-extensions : AVAIL 16" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg14-extensions : AVAIL 19" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg18-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg17-extensions : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg16-extensions : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg15-extensions : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg14-extensions : AVAIL 15" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg18-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg17-extensions : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg16-extensions : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg15-extensions : AVAIL 16" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg14-extensions : AVAIL 18" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg18-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg17-extensions : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg16-extensions : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg15-extensions : AVAIL 16" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg14-extensions : AVAIL 16" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg18-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg17-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg16-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg15-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg14-extensions : AVAIL 8" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg18-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg17-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg16-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg15-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7.1" "pgpool-II-pg14-extensions : AVAIL 8" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.7.1" "postgresql-18-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-17-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-16-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-15-pgpool2 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7.1" "postgresql-14-pgpool2 : AVAIL 2" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://pgpool.net/" title="Repository" icon="link" subtitle="pgpool.net/" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgpool;		# install via package name, for the active PG version
pig install pgpool_recovery;		# install by extension name, for the current active PG version

pig install pgpool_recovery -v 18;   # install for PG 18
pig install pgpool_recovery -v 17;   # install for PG 17
pig install pgpool_recovery -v 16;   # install for PG 16
pig install pgpool_recovery -v 15;   # install for PG 15
pig install pgpool_recovery -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgpool_recovery;
```




## Usage

> [pgpool_recovery: recovery functions for pgpool-II](https://pgpool.net/)

The `pgpool_recovery` extension provides recovery-related functions used by Pgpool-II for online recovery of backend PostgreSQL nodes.

### Functions

```sql
-- Trigger online recovery of a backend node
-- Executes the recovery script on the primary node
SELECT pgpool_recovery(
    'recovery_1st_stage_script',   -- script name in $PGDATA
    'target_hostname',             -- hostname of node to recover
    'target_pgdata',               -- data directory on target
    'target_port'                  -- port number on target
);

-- Second stage recovery (optional, for streaming replication)
SELECT pgpool_remote_start(
    'target_hostname',             -- hostname of recovered node
    'target_pgdata'                -- data directory on target
);

-- Check if the target node is ready
SELECT pgpool_pgctl('status', 'target_pgdata');
```

### How It Works

1. Pgpool-II calls `pgpool_recovery()` on the primary node, which executes a user-defined shell script to perform base backup and setup
2. The recovery script copies data to the target node and configures replication
3. `pgpool_remote_start()` starts the recovered PostgreSQL instance
4. Pgpool-II attaches the recovered node back into the pool

### Recovery Scripts

The recovery script (e.g., `recovery_1st_stage`) must be placed in the primary node's `$PGDATA` directory. It typically performs:

- `pg_basebackup` to copy data to the target
- Configuration of `primary_conninfo` for streaming replication
- Creation of `standby.signal` on the target

The extension must be installed on all PostgreSQL backend nodes managed by Pgpool-II.
