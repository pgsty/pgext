---
title: "citus_columnar"
linkTitle: "citus_columnar"
description: "Citus columnar storage engine"
weight: 2401
categories: ["OLAP"]
width: full
---

[**citus**](https://github.com/citusdata/citus) : Citus columnar storage engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2401** | {{< badge content="citus_columnar" link="https://github.com/citusdata/citus" >}} | {{< ext "citus_columnar" "citus" >}} | `14.0.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "columnar" >}} {{< ext "pg_parquet" >}} {{< ext "timescaledb" >}} {{< ext "pg_analytics" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "orioledb" >}} |
|    **Siblings**   | {{< ext "citus" >}} |

> [!Note] conflict with hydra columnar, no pg18


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "citus_18" "green" >}} {{< bg "17" "citus_17" "green" >}} {{< bg "16" "citus_16" "green" >}} {{< bg "15" "citus_15" "red" >}} {{< bg "14" "citus_14" "red" >}} | `citus_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.0.0` | {{< bg "18" "postgresql-18-citus" "green" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "red" >}} {{< bg "14" "postgresql-14-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 29" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 21" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 26" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "citus_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 14.0.0" "citus_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.0.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-14.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg citus;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install citus;		# install via package name, for the active PG version
pig install citus_columnar;		# install by extension name, for the current active PG version

pig install citus_columnar -v 18;   # install for PG 18
pig install citus_columnar -v 17;   # install for PG 17
pig install citus_columnar -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citus_columnar;
```
