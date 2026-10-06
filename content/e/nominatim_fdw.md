---
title: "nominatim_fdw"
linkTitle: "nominatim_fdw"
description: "Nominatim Foreign Data Wrapper for PostgreSQL"
weight: 8680
categories: ["FDW"]
languages: ["C"]
licenses: ["MIT"]
repos: ["MIXED"]
page_width: full
---

[**nominatim_fdw**](https://github.com/jimjonesbr/nominatim_fdw) : Nominatim Foreign Data Wrapper for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8680** | {{< badge content="nominatim_fdw" link="https://github.com/jimjonesbr/nominatim_fdw" >}} | {{< ext "nominatim_fdw" >}} | `2.3` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} {{< ext "tzf" >}} {{< ext "pg_geohash" >}} {{< ext "rdf_fdw" >}} |

> [!Note] SQL/control and tagged release version 2.3; PGXN distribution 2.3.0. Maintained in the local RPM and DEB recipes.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `nominatim_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.3` | {{< bg "18" "nominatim_fdw_18" "green" >}} {{< bg "17" "nominatim_fdw_17" "green" >}} {{< bg "16" "nominatim_fdw_16" "green" >}} {{< bg "15" "nominatim_fdw_15" "green" >}} {{< bg "14" "nominatim_fdw_14" "green" >}} | `nominatim_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.3` | {{< bg "18" "postgresql-18-nominatim-fdw" "green" >}} {{< bg "17" "postgresql-17-nominatim-fdw" "green" >}} {{< bg "16" "postgresql-16-nominatim-fdw" "green" >}} {{< bg "15" "postgresql-15-nominatim-fdw" "green" >}} {{< bg "14" "postgresql-14-nominatim-fdw" "green" >}} | `postgresql-$v-nominatim-fdw` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_14 : AVAIL 6" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_14 : AVAIL 6" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_18 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_16 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_15 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_14 : AVAIL 13" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_18 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_16 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_15 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_14 : AVAIL 13" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_18 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_16 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_15 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_14 : AVAIL 13" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_18 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_17 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_16 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_15 : AVAIL 13" "green" >}} | {{< bg "PIGSTY 2.3" "nominatim_fdw_14 : AVAIL 13" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.3" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_18` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 89.4 KiB | [nominatim_fdw_18-2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/nominatim_fdw_18-2.3-1PGSTY.el8.x86_64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [nominatim_fdw_18-2.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-2.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 87.5 KiB | [nominatim_fdw_18-2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/nominatim_fdw_18-2.3-1PGSTY.el8.aarch64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.8 KiB | [nominatim_fdw_18-2.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-2.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.1 KiB | [nominatim_fdw_18-2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/nominatim_fdw_18-2.3-1PGSTY.el9.x86_64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.0 KiB | [nominatim_fdw_18-2.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-2.3-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_18` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [nominatim_fdw_18-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_18` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.3 KiB | [nominatim_fdw_18-2.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-2.1-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.6 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.3-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.3-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.2 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 89.7 KiB | [nominatim_fdw_18-2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/nominatim_fdw_18-2.3-1PGSTY.el9.aarch64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.9 KiB | [nominatim_fdw_18-2.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-2.3-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_18` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.1 KiB | [nominatim_fdw_18-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_18` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.5 KiB | [nominatim_fdw_18-2.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-2.1-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.5 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.3-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.3-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.7 KiB | [nominatim_fdw_18-2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/nominatim_fdw_18-2.3-1PGSTY.el10.x86_64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [nominatim_fdw_18-2.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-2.3-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_18` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [nominatim_fdw_18-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_18` | `2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.9 KiB | [nominatim_fdw_18-2.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-2.1-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.3-2PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.7 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.3-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.3-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.3 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.6 KiB | [nominatim_fdw_18-2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/nominatim_fdw_18-2.3-1PGSTY.el10.aarch64.rpm) |
| `nominatim_fdw_18` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [nominatim_fdw_18-2.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-2.3-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_18` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [nominatim_fdw_18-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_18` | `2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.9 KiB | [nominatim_fdw_18-2.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-2.1-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_18-1.3-2PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.3-2PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.3-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_18-1.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.3-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.8 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-nominatim-fdw` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.7 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 86.5 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 88.8 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 86.7 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 85.8 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.2 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.9 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.6 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 81.9 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.5 KiB | [postgresql-18-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_17` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 89.3 KiB | [nominatim_fdw_17-2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/nominatim_fdw_17-2.3-1PGSTY.el8.x86_64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [nominatim_fdw_17-2.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-2.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 87.4 KiB | [nominatim_fdw_17-2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/nominatim_fdw_17-2.3-1PGSTY.el8.aarch64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [nominatim_fdw_17-2.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-2.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.1 KiB | [nominatim_fdw_17-2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/nominatim_fdw_17-2.3-1PGSTY.el9.x86_64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.0 KiB | [nominatim_fdw_17-2.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-2.3-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_17` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [nominatim_fdw_17-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_17` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.3 KiB | [nominatim_fdw_17-2.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-2.1-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.3-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.3-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.0 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.2 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 89.5 KiB | [nominatim_fdw_17-2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/nominatim_fdw_17-2.3-1PGSTY.el9.aarch64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.9 KiB | [nominatim_fdw_17-2.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-2.3-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_17` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.0 KiB | [nominatim_fdw_17-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_17` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.5 KiB | [nominatim_fdw_17-2.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-2.1-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.5 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.3-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.3-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.7 KiB | [nominatim_fdw_17-2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/nominatim_fdw_17-2.3-1PGSTY.el10.x86_64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [nominatim_fdw_17-2.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-2.3-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_17` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [nominatim_fdw_17-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_17` | `2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [nominatim_fdw_17-2.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-2.1-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.3-2PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.7 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.3-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.3-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.7 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.2 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.4 KiB | [nominatim_fdw_17-2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/nominatim_fdw_17-2.3-1PGSTY.el10.aarch64.rpm) |
| `nominatim_fdw_17` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [nominatim_fdw_17-2.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-2.3-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_17` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [nominatim_fdw_17-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_17` | `2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.0 KiB | [nominatim_fdw_17-2.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-2.1-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_17-1.3-2PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.3-2PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.3-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_17-1.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.3-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-nominatim-fdw` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.7 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 86.7 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 89.0 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 86.9 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 94.1 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 92.6 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.8 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.7 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 81.8 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.6 KiB | [postgresql-17-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_16` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 89.3 KiB | [nominatim_fdw_16-2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/nominatim_fdw_16-2.3-1PGSTY.el8.x86_64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [nominatim_fdw_16-2.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-2.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 87.3 KiB | [nominatim_fdw_16-2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/nominatim_fdw_16-2.3-1PGSTY.el8.aarch64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.8 KiB | [nominatim_fdw_16-2.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-2.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.1 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.1 KiB | [nominatim_fdw_16-2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/nominatim_fdw_16-2.3-1PGSTY.el9.x86_64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.0 KiB | [nominatim_fdw_16-2.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-2.3-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [nominatim_fdw_16-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_16` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.3 KiB | [nominatim_fdw_16-2.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-2.1-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.6 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.3-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.3-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.2 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 89.5 KiB | [nominatim_fdw_16-2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/nominatim_fdw_16-2.3-1PGSTY.el9.aarch64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [nominatim_fdw_16-2.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-2.3-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.1 KiB | [nominatim_fdw_16-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_16` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.6 KiB | [nominatim_fdw_16-2.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-2.1-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.3-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.3-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.7 KiB | [nominatim_fdw_16-2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/nominatim_fdw_16-2.3-1PGSTY.el10.x86_64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [nominatim_fdw_16-2.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-2.3-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_16` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [nominatim_fdw_16-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_16` | `2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [nominatim_fdw_16-2.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-2.1-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.2 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.3-2PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.7 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.3-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.3-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.3 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.5 KiB | [nominatim_fdw_16-2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/nominatim_fdw_16-2.3-1PGSTY.el10.aarch64.rpm) |
| `nominatim_fdw_16` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [nominatim_fdw_16-2.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-2.3-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_16` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [nominatim_fdw_16-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_16` | `2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.0 KiB | [nominatim_fdw_16-2.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-2.1-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [nominatim_fdw_16-1.3-2PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.3-2PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.3-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_16-1.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.3-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-nominatim-fdw` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.8 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 86.7 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 88.9 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 86.9 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 94.0 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 92.5 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.8 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.7 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 81.8 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.6 KiB | [postgresql-16-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_15` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 89.4 KiB | [nominatim_fdw_15-2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/nominatim_fdw_15-2.3-1PGSTY.el8.x86_64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [nominatim_fdw_15-2.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-2.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.7 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 87.4 KiB | [nominatim_fdw_15-2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/nominatim_fdw_15-2.3-1PGSTY.el8.aarch64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [nominatim_fdw_15-2.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-2.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.2 KiB | [nominatim_fdw_15-2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/nominatim_fdw_15-2.3-1PGSTY.el9.x86_64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.0 KiB | [nominatim_fdw_15-2.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-2.3-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [nominatim_fdw_15-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.4 KiB | [nominatim_fdw_15-2.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-2.1-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.6 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.3-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.3-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.2 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 89.5 KiB | [nominatim_fdw_15-2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/nominatim_fdw_15-2.3-1PGSTY.el9.aarch64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.9 KiB | [nominatim_fdw_15-2.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-2.3-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.0 KiB | [nominatim_fdw_15-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.6 KiB | [nominatim_fdw_15-2.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-2.1-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.3-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.3-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.8 KiB | [nominatim_fdw_15-2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/nominatim_fdw_15-2.3-1PGSTY.el10.x86_64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [nominatim_fdw_15-2.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-2.3-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_15` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [nominatim_fdw_15-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_15` | `2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [nominatim_fdw_15-2.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-2.1-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.2 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.3-2PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.7 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.3-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.3-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.3 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.5 KiB | [nominatim_fdw_15-2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/nominatim_fdw_15-2.3-1PGSTY.el10.aarch64.rpm) |
| `nominatim_fdw_15` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [nominatim_fdw_15-2.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-2.3-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_15` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [nominatim_fdw_15-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_15` | `2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.0 KiB | [nominatim_fdw_15-2.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-2.1-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_15-1.3-2PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.3-2PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.3-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_15-1.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.3-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-15-nominatim-fdw` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.7 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 86.7 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 88.8 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 86.8 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.9 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 92.5 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.8 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.7 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 81.9 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.5 KiB | [postgresql-15-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_14` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 89.3 KiB | [nominatim_fdw_14-2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/nominatim_fdw_14-2.3-1PGSTY.el8.x86_64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [nominatim_fdw_14-2.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-2.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.3-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 87.3 KiB | [nominatim_fdw_14-2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/nominatim_fdw_14-2.3-1PGSTY.el8.aarch64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [nominatim_fdw_14-2.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-2.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.3-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.2 KiB | [nominatim_fdw_14-2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/nominatim_fdw_14-2.3-1PGSTY.el9.x86_64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.0 KiB | [nominatim_fdw_14-2.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-2.3-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [nominatim_fdw_14-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.4 KiB | [nominatim_fdw_14-2.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-2.1-1PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel9.8.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.6 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.3-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.3-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.2 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 89.5 KiB | [nominatim_fdw_14-2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/nominatim_fdw_14-2.3-1PGSTY.el9.aarch64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.9 KiB | [nominatim_fdw_14-2.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-2.3-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.0 KiB | [nominatim_fdw_14-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.6 KiB | [nominatim_fdw_14-2.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-2.1-1PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel9.8.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.3-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.3-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.8 KiB | [nominatim_fdw_14-2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/nominatim_fdw_14-2.3-1PGSTY.el10.x86_64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [nominatim_fdw_14-2.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-2.3-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_14` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [nominatim_fdw_14-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_14` | `2.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.0 KiB | [nominatim_fdw_14-2.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-2.1-1PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel10.2.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.8 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.2 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.3-2PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.7 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.3-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.3-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.8 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.3 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 90.4 KiB | [nominatim_fdw_14-2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/nominatim_fdw_14-2.3-1PGSTY.el10.aarch64.rpm) |
| `nominatim_fdw_14` | `2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [nominatim_fdw_14-2.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-2.3-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_14` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [nominatim_fdw_14-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_14` | `2.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.0 KiB | [nominatim_fdw_14-2.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-2.1-1PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel10.2.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [nominatim_fdw_14-1.3-2PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.3-2PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.3-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [nominatim_fdw_14-1.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.3-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel10.0.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-14-nominatim-fdw` | `2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 88.7 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 86.6 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 88.9 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 86.9 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 93.8 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 92.4 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 82.9 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.7 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 81.9 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.6 KiB | [postgresql-14-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/nominatim_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/nominatim_fdw" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="nominatim_fdw-2.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg nominatim_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install nominatim_fdw;		# install via package name, for the active PG version

pig install nominatim_fdw -v 18;   # install for PG 18
pig install nominatim_fdw -v 17;   # install for PG 17
pig install nominatim_fdw -v 16;   # install for PG 16
pig install nominatim_fdw -v 15;   # install for PG 15
pig install nominatim_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION nominatim_fdw;
```

## Usage

Sources:

- [nominatim_fdw v2.3 README](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.3/README.md)
- [nominatim_fdw v2.3 changelog](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.3/CHANGELOG.md)
- [nominatim_fdw v2.3 control file](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.3/nominatim_fdw.control)
- [Official Nominatim API overview](https://nominatim.org/release-docs/develop/api/Overview/)
- [OpenStreetMap Nominatim usage policy](https://operations.osmfoundation.org/policies/nominatim/)

`nominatim_fdw` calls a Nominatim geocoding service from PostgreSQL. Unlike a conventional FDW, it exposes functions for search, reverse geocoding, and OSM-object lookup; it does not create queryable foreign tables.

### Configure a Server

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
  FOREIGN DATA WRAPPER nominatim_fdw
  OPTIONS (
    url 'https://nominatim.openstreetmap.org',
    connect_timeout '10',
    max_connect_retry '2',
    max_connect_redirect '1',
    accept_language 'en'
  );
```

The public OpenStreetMap endpoint has an official usage policy. For sustained or bulk workloads, use an authorized provider or operate your own Nominatim service, identify the application as required, and respect rate limits.

Version 2.1 supports HTTP Basic authentication through a per-user mapping:

```sql
CREATE USER MAPPING FOR app_user
SERVER osm
OPTIONS (user 'remote_user', password 'secret');
```

Keep endpoint and proxy credentials in `USER MAPPING`, where PostgreSQL applies its credential-visibility rules, rather than in foreign-server options.

### Core Workflow

Free-form search:

```sql
SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => 'Neubrückenstraße 63, Münster, Germany'
);
```

Reverse geocoding and object lookup:

```sql
SELECT osm_id, display_name, addressdetails
FROM nominatim_reverse(
  server_name => 'osm',
  lon => 7.6293,
  lat => 51.9648,
  addressdetails => true
);

SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959'
);
```

### Important Objects

- `nominatim_search(...)` implements free-form or structured forward search.
- `nominatim_reverse(...)` resolves longitude and latitude to the nearest suitable OSM address.
- `nominatim_lookup(...)` fetches node, way, or relation identifiers such as `N123`, `W456`, or `R789`.
- `nominatim_fdw_version()` reports the extension and principal library versions.
- `nominatim_fdw_settings` exposes dependency and build versions as rows.
- Server options include `url`, proxy configuration, timeouts, retry and redirect limits, and default `accept_language`.

All endpoint functions are `STRICT`: an explicit SQL `NULL` argument returns no rows without sending a request. In 2.0 they are correctly declared `VOLATILE`, because responses are remote and can change.

### Version 2.3 Changes and Caveats

Version 2.0 validates reverse coordinates, adds `email`, `polygon_threshold`, and `entrances`, exposes dependency settings, and fixes JSON escaping for returned detail fields. It also has user-visible compatibility changes: reverse output uses `display_name`; `addressparts` becomes `addressdetails`; address details default to true for reverse and lookup; and version output is shorter.

After installing the 2.3 extension files, upgrade existing databases:

```sql
ALTER EXTENSION nominatim_fdw UPDATE TO '2.3';
```

Version 2.1 adds Basic-auth mapping and initializes libcurl once per PostgreSQL backend instead of relying on implicit per-request initialization. The libcurl fix requires no new SQL or preload setting.

Each call performs network I/O in the database statement. Use finite timeouts, constrain who can create or alter servers, and avoid invoking a public service once per row in a large query. The upstream build requires PostgreSQL 10 or newer, libxml2 2.6.0 or newer, and libcurl 7.74 or newer.

### 2.3 Privileges and Request Limits

The SQL/control version is 2.3; the PGXN distribution version is 2.3.0. Query functions now enforce foreign-server `USAGE`; grant it to roles that need access:

```sql
GRANT USAGE ON FOREIGN SERVER osm TO app_user;
ALTER SERVER osm OPTIONS (ADD max_response_size '10485760');
```

`max_response_size` is in bytes, defaults to 0 for no additional limit, and accepts a maximum of 1 GiB. 2.3 allows only HTTP/HTTPS and rejects CR/LF in request headers; `polygon_threshold` must be finite and nonnegative. `zoom` preserves -1 as the default and clamps other values to 0–18. It also fixes response encoding, XML root validation and cleanup after cancellation. The catalog version does not imply that existing 2.2.0 packages contain these fixes.
