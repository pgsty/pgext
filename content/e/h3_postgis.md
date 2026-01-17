---
title: "h3_postgis"
linkTitle: "h3_postgis"
description: "H3 PostGIS integration"
weight: 1531
categories: ["GIS"]
width: full
---

[**pg_h3**](https://github.com/zachasme/h3-pg) : H3 PostGIS integration


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1531** | {{< badge content="h3_postgis" link="https://github.com/zachasme/h3-pg" >}} | {{< ext "h3_postgis" "pg_h3" >}} | `4.2.3` | {{< category "GIS" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "h3" >}} {{< ext "postgis" >}} {{< ext "postgis_raster" >}} |
|   **See Also**    | {{< ext "mobilitydb" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} |
|    **Siblings**   | {{< ext "h3" >}} |

> [!Note] pgdg missing el8.x86.pg17 and el8.x86.pg18


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_h3` | `h3`, `postgis`, `postgis_raster` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.3` | {{< bg "18" "h3-pg_18" "green" >}} {{< bg "17" "h3-pg_17" "green" >}} {{< bg "16" "h3-pg_16" "green" >}} {{< bg "15" "h3-pg_15" "green" >}} {{< bg "14" "h3-pg_14" "green" >}} {{< bg "13" "h3-pg_13" "green" >}} | `h3-pg_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.3` | {{< bg "18" "postgresql-18-h3" "green" >}} {{< bg "17" "postgresql-17-h3" "green" >}} {{< bg "16" "postgresql-16-h3" "green" >}} {{< bg "15" "postgresql-15-h3" "green" >}} {{< bg "14" "postgresql-14-h3" "green" >}} {{< bg "13" "postgresql-13-h3" "green" >}} | `postgresql-$v-h3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "h3-pg_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "h3-pg_17 : MISS 0" "red" >}}      | {{< bg "PGDG 4.1.3" "h3-pg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.2.3" "h3-pg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.2.3" "h3-pg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.2.3" "h3-pg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.3" "h3-pg_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.2.3" "h3-pg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.2.3" "h3-pg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.2" "h3-pg_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.2.3" "postgresql-18-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-17-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-16-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-15-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-14-h3 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "postgresql-13-h3 : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zachasme/h3-pg" title="Repository" icon="github" subtitle="github.com/zachasme/h3-pg" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_h3;		# install via package name, for the active PG version
pig install h3_postgis;		# install by extension name, for the current active PG version

pig install h3_postgis -v 18;   # install for PG 18
pig install h3_postgis -v 17;   # install for PG 17
pig install h3_postgis -v 16;   # install for PG 16
pig install h3_postgis -v 15;   # install for PG 15
pig install h3_postgis -v 14;   # install for PG 14
pig install h3_postgis -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION h3_postgis CASCADE; -- requires h3, postgis, postgis_raster
```
