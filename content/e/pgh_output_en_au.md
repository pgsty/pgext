---
title: "pgh_output_en_au"
linkTitle: "pgh_output_en_au"
description: "Australian English output pack for PgHydro"
weight: 1604
categories: ["GIS"]
width: full
---

[**pghydro**](https://github.com/pghydro/pghydro) : Australian English output pack for PgHydro


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1604** | {{< badge content="pgh_output_en_au" link="https://github.com/pghydro/pghydro" >}} | {{< ext "pgh_output_en_au" "pghydro" >}} | `6.6` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgh_output_en_au` |
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "postgis" >}} {{< ext "pghydro" >}} |
|    **Siblings**   | {{< ext "pghydro" >}} {{< ext "pgh_raster" >}} {{< ext "pgh_hgm" >}} {{< ext "pgh_output" >}} {{< ext "pgh_output_pt_br" >}} {{< ext "pgh_consistency" >}} |


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
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |


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
pig install pgh_output_en_au;		# install by extension name, for the current active PG version

pig install pgh_output_en_au -v 18;   # install for PG 18
pig install pgh_output_en_au -v 17;   # install for PG 17
pig install pgh_output_en_au -v 16;   # install for PG 16
pig install pgh_output_en_au -v 15;   # install for PG 15
pig install pgh_output_en_au -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgh_output_en_au CASCADE; -- requires plpgsql, postgis, pghydro
```




## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_output_en_au SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output_en-au--6.6.sql)

`pgh_output_en_au` provides an Australian-English PgHydro output schema. It creates localized export tables for drainage lines, drainage areas, and drainage points, plus a refresh function that populates them from the base PgHydro network.

### Enable

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output_en_au;
```

### Output Tables

The extension creates the `pgh_output_en_au` schema with tables including:

| Table | Purpose |
|-------|---------|
| `geoft_bho_drainage_line` | Drainage-line export table |
| `geoft_bho_drainage_area` | Drainage-area export table |
| `geoft_bho_drainage_point` | Drainage-point export table |

### Refresh Output

```sql
SELECT pgh_output_en_au.pghfn_updateexporttables();
```

Run the refresh after recalculating or editing the source PgHydro objects.

### Notes

Use this extension when the exported schema should expose Australian-English naming conventions while still being derived from the PgHydro drainage network.
