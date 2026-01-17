---
title: "pglogical_origin"
linkTitle: "pglogical_origin"
description: "Dummy extension for compatibility when upgrading from Postgres 9.4"
weight: 9501
categories: ["ETL"]
width: full
---

[**pglogical**](https://github.com/2ndQuadrant/pglogical) : Dummy extension for compatibility when upgrading from Postgres 9.4


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9501** | {{< badge content="pglogical_origin" link="https://github.com/2ndQuadrant/pglogical" >}} | {{< ext "pglogical_origin" "pglogical" >}} | `2.4.6` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pglogical_origin` |
|   **See Also**    | {{< ext "pglogical_ticker" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "repmgr" >}} {{< ext "decoder_raw" >}} |
|    **Siblings**   | {{< ext "pglogical" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pglogical` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "pglogical_18" "green" >}} {{< bg "17" "pglogical_17" "green" >}} {{< bg "16" "pglogical_16" "green" >}} {{< bg "15" "pglogical_15" "green" >}} {{< bg "14" "pglogical_14" "green" >}} {{< bg "13" "pglogical_13" "green" >}} | `pglogical_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "postgresql-18-pglogical" "green" >}} {{< bg "17" "postgresql-17-pglogical" "green" >}} {{< bg "16" "postgresql-16-pglogical" "green" >}} {{< bg "15" "postgresql-15-pglogical" "green" >}} {{< bg "14" "postgresql-14-pglogical" "green" >}} {{< bg "13" "postgresql-13-pglogical" "green" >}} | `postgresql-$v-pglogical` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/2ndQuadrant/pglogical" title="Repository" icon="github" subtitle="github.com/2ndQuadrant/pglogical" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglogical_ticker-1.4.1.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pglogical;		# install via package name, for the active PG version
pig install pglogical_origin;		# install by extension name, for the current active PG version

pig install pglogical_origin -v 18;   # install for PG 18
pig install pglogical_origin -v 17;   # install for PG 17
pig install pglogical_origin -v 16;   # install for PG 16
pig install pglogical_origin -v 15;   # install for PG 15
pig install pglogical_origin -v 14;   # install for PG 14
pig install pglogical_origin -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglogical_origin;
```
