---
title: "citus_columnar"
linkTitle: "citus_columnar"
description: "Citus columnar storage engine"
weight: 2401
categories: ["OLAP"]
width: full
---

Citus columnar storage engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2401** | {{< badge content="citus_columnar" link="https://github.com/citusdata/citus" >}} | {{< ext "citus_columnar" "citus" >}} | `13.2.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "columnar" >}} {{< ext "pg_parquet" >}} {{< ext "timescaledb" >}} {{< ext "pg_analytics" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "orioledb" >}} |
|    **Siblings**   | {{< ext "citus" >}} |

> [!Note] conflict with hydra columnar


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/citus" >}} | `13.2.0` | {{< bg "18" "citus_18*" "red" >}} {{< bg "17" "citus_17*" "green" >}} {{< bg "16" "citus_16*" "green" >}} {{< bg "15" "citus_15*" "green" >}} {{< bg "14" "citus_14*" "green" >}} {{< bg "13" "citus_13*" "red" >}} | `citus_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/citus" >}} | `13.2.0` | {{< bg "18" "postgresql-18-citus" "red" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "green" >}} {{< bg "14" "postgresql-14-citus" "green" >}} {{< bg "13" "postgresql-13-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 23" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 29" "green" >}} | {{< bg "PGDG 9.5.4" "citus_13 : AVAIL 29" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 22" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 23" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 26" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 16" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 15" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 23" "green" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 16" "green" >}} | {{< bg "PGDG 11.3.0" "citus_13 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "citus_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "citus_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "citus_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 13.2.0" "citus_15 : AVAIL 5" "green" >}} |      {{< bg "MISS" "citus_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "citus_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-citus : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-citus : MISS 0" "red" >}}      | {{< bg "PIGSTY 13.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-citus : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-13.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get citus_columnar; # get citus_columnar source code
pig build dep citus_columnar; # install build dependencies
pig build pkg citus_columnar; # build extension rpm or deb
pig build ext citus_columnar; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install citus_columnar; # install by extension name, for the current active PG version
pig ext install citus; # install via package alias, for the active PG version
pig ext install citus_columnar -v 17;   # install for PG 17
pig ext install citus_columnar -v 16;   # install for PG 16
pig ext install citus_columnar -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION citus_columnar CASCADE SCHEMA pg_catalog;
```

