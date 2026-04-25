---
title: "pgh_output"
linkTitle: "pgh_output"
description: "Output and reporting objects for PgHydro"
weight: 1603
categories: ["GIS"]
width: full
---

[**pghydro**](https://github.com/pghydro/pghydro) : Output and reporting objects for PgHydro


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1603** | {{< badge content="pgh_output" link="https://github.com/pghydro/pghydro" >}} | {{< ext "pgh_output" "pghydro" >}} | `6.6` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgh_output` |
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "postgis" >}} {{< ext "pghydro" >}} |
|    **Siblings**   | {{< ext "pghydro" >}} {{< ext "pgh_raster" >}} {{< ext "pgh_hgm" >}} {{< ext "pgh_output_en_au" >}} {{< ext "pgh_output_pt_br" >}} {{< ext "pgh_consistency" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `6.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pghydro` | `plpgsql`, `postgis`, `pghydro` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `6.6` | {{< bg "18" "pghydro_18" "green" >}} {{< bg "17" "pghydro_17" "green" >}} {{< bg "16" "pghydro_16" "green" >}} {{< bg "15" "pghydro_15" "green" >}} {{< bg "14" "pghydro_14" "green" >}} | `pghydro_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `6.6` | {{< bg "18" "postgresql-18-pghydro" "green" >}} {{< bg "17" "postgresql-17-pghydro" "green" >}} {{< bg "16" "postgresql-16-pghydro" "green" >}} {{< bg "15" "postgresql-15-pghydro" "green" >}} {{< bg "14" "postgresql-14-pghydro" "green" >}} | `postgresql-$v-pghydro` | `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pghydro : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pghydro : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pghydro : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pghydro/pghydro" title="Repository" icon="github" subtitle="github.com/pghydro/pghydro" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pghydro-6.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pghydro;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pghydro;		# install via package name, for the active PG version
pig install pgh_output;		# install by extension name, for the current active PG version

pig install pgh_output -v 18;   # install for PG 18
pig install pgh_output -v 17;   # install for PG 17
pig install pgh_output -v 16;   # install for PG 16
pig install pgh_output -v 15;   # install for PG 15
pig install pgh_output -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgh_output CASCADE; -- requires plpgsql, postgis, pghydro
```
