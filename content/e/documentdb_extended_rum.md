---
title: "documentdb_extended_rum"
linkTitle: "documentdb_extended_rum"
description: "DocumentDB Extended RUM index access method"
weight: 9030
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : DocumentDB Extended RUM index access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9030** | {{< badge content="documentdb_extended_rum" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb_extended_rum" "documentdb" >}} | `0.110` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "rum" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} {{< ext "mongo_fdw" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.110` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `documentdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.110` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.110` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.109" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.109" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.109" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.109" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.109" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.109" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 2" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 2" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.109" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.109" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.108" "postgresql-18-documentdb : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.108" "postgresql-17-documentdb : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.108" "postgresql-16-documentdb : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.108" "postgresql-15-documentdb : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.108" "postgresql-18-documentdb : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.108" "postgresql-17-documentdb : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.108" "postgresql-16-documentdb : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.108" "postgresql-15-documentdb : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.110-0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install documentdb;		# install via package name, for the active PG version
pig install documentdb_extended_rum;		# install by extension name, for the current active PG version

pig install documentdb_extended_rum -v 18;   # install for PG 18
pig install documentdb_extended_rum -v 17;   # install for PG 17
pig install documentdb_extended_rum -v 16;   # install for PG 16
pig install documentdb_extended_rum -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_documentdb_extended_rum';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb_extended_rum;
```



## Usage

> [documentdb_extended_rum: DocumentDB Extended RUM index access method](https://github.com/documentdb/documentdb)

The `documentdb_extended_rum` extension provides an enhanced RUM (Recursive Union Merge) index access method for DocumentDB on PostgreSQL. It improves query performance for document-based workloads.

### Overview

This extension extends the RUM index type to better support BSON document indexing within DocumentDB. It provides optimized index access methods for:

- Full-text search on document fields
- Compound index operations on BSON data
- Efficient range queries and sorting on indexed document properties

### Prerequisites

Requires `documentdb_core` to be installed.

### Enabling

```sql
CREATE EXTENSION documentdb_extended_rum;
```

The extended RUM indexes are automatically utilized by the DocumentDB query planner when appropriate for document query patterns.
