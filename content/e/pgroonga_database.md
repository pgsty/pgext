---
title: "pgroonga_database"
linkTitle: "pgroonga_database"
description: "PGroonga database management module"
weight: 2111
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pgroonga**](https://github.com/pgroonga/pgroonga) : PGroonga database management module


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2111** | {{< badge content="pgroonga_database" link="https://github.com/pgroonga/pgroonga" >}} | {{< ext "pgroonga_database" "pgroonga" >}} | `4.0.9` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_fts" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_rrf" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |
|    **Siblings**   | {{< ext "pgroonga" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgroonga` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.9` | {{< bg "18" "pgroonga_18" "green" >}} {{< bg "17" "pgroonga_17" "green" >}} {{< bg "16" "pgroonga_16" "green" >}} {{< bg "15" "pgroonga_15" "green" >}} {{< bg "14" "pgroonga_14" "green" >}} | `pgroonga_$v` | `groonga-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0.9` | {{< bg "18" "postgresql-18-pgroonga" "green" >}} {{< bg "17" "postgresql-17-pgroonga" "green" >}} {{< bg "16" "postgresql-16-pgroonga" "green" >}} {{< bg "15" "postgresql-15-pgroonga" "green" >}} {{< bg "14" "postgresql-14-pgroonga" "green" >}} | `postgresql-$v-pgroonga` | `libgroonga0` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "pgroonga_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-18-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-17-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-16-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-15-pgroonga : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0.9" "postgresql-14-pgroonga : AVAIL 1" "green" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgroonga/pgroonga" title="Repository" icon="github" subtitle="github.com/pgroonga/pgroonga" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgroonga-4.0.9.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgroonga;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgroonga;		# install via package name, for the active PG version
pig install pgroonga_database;		# install by extension name, for the current active PG version

pig install pgroonga_database -v 18;   # install for PG 18
pig install pgroonga_database -v 17;   # install for PG 17
pig install pgroonga_database -v 16;   # install for PG 16
pig install pgroonga_database -v 15;   # install for PG 15
pig install pgroonga_database -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgroonga_database;
```

## Usage

Sources:

- [Version 4.0.9 SQL](https://github.com/pgroonga/pgroonga/blob/4.0.9/data/pgroonga_database.sql)
- [Version 4.0.9 control](https://github.com/pgroonga/pgroonga/blob/4.0.9/pgroonga_database.control)
- [Version 4.0.9 implementation](https://github.com/pgroonga/pgroonga/blob/4.0.9/src/pgroonga-database.c)
- [Official recovery procedure](https://pgroonga.github.io/reference/functions/pgroonga-database-remove.html)

`pgroonga_database` 4.0.9 is a recovery-only helper for a damaged internal PGroonga database. It provides one SQL function that removes PGroonga files from database and applicable tablespace directories. It does not provide the search access method.

### Recovery Workflow

Normal index corruption may be repairable with REINDEX alone. Use this module only when the internal Groonga database itself is damaged and a planned rebuild is necessary. Arrange a recovery window and disconnect every session using PGroonga before proceeding; remaining sessions may crash when their files are removed.

In a fresh administrative connection that has not opened any PGroonga index:

```sql
CREATE EXTENSION pgroonga_database;
SELECT pgroonga_database_remove();
```

Disconnect that connection immediately afterward. In another fresh connection, run REINDEX for **every** PGroonga index to recreate the internal database from PostgreSQL table data. Resume application traffic only after rebuilding and checking all affected indexes.

### Return Value and Boundaries

`pgroonga_database_remove()` returns true when it reaches the end of its cleanup loop. If a tablespace ownership check fails, the loop stops and the function can still return true; this result alone does not prove that every location was cleaned. Other failures can raise errors. It removes the internal files directly; it neither exports them nor rebuilds indexes. Do not use other PGroonga features in the cleanup connection. This is not a routine vacuum, an uninstall command, or an action that can be made safe merely by wrapping the call in a SQL transaction.

The control file does not mark the extension trusted or relocatable. The C implementation checks tablespace ownership when traversing locations. Use an administrator who owns the required locations and verify that the cleanup and full reindex completed. The module has no preload requirement; enable it only for the recovery task.
