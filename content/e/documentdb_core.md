---
title: "documentdb_core"
linkTitle: "documentdb_core"
description: "Core API surface for DocumentDB for PostgreSQL"
weight: 9010
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/microsoft/documentdb) : Core API surface for DocumentDB for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9010** | {{< badge content="documentdb_core" link="https://github.com/microsoft/documentdb" >}} | {{< ext "documentdb_core" "documentdb" >}} | `0.106` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_cron" >}} {{< ext "postgis" >}} {{< ext "vector" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] FerretDB Fork, work with FerretDB 2.5


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.106` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `documentdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.106` | {{< bg "18" "documentdb_18*" "red" >}} {{< bg "17" "documentdb_17*" "green" >}} {{< bg "16" "documentdb_16*" "green" >}} {{< bg "15" "documentdb_15*" "green" >}} {{< bg "14" "documentdb_14*" "red" >}} {{< bg "13" "documentdb_13*" "red" >}} | `documentdb_$v*` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.106` | {{< bg "18" "postgresql-18-documentdb" "red" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} {{< bg "13" "postgresql-13-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/documentdb" title="Repository" icon="github" subtitle="github.com/microsoft/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.107.0-ferretdb-2.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install documentdb;		# install via package name, for the active PG version
pig install documentdb_core;		# install by extension name, for the current active PG version

pig install documentdb_core -v 17;   # install for PG 17
pig install documentdb_core -v 16;   # install for PG 16
pig install documentdb_core -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_documentdb, pg_documentdb_core';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb_core;
```
