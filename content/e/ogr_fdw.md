---
title: "ogr_fdw"
linkTitle: "ogr_fdw"
description: "foreign-data wrapper for GIS data access"
weight: 1550
categories: ["GIS"]
width: full
---

[**ogr_fdw**](https://github.com/pramsey/pgsql-ogr-fdw) : foreign-data wrapper for GIS data access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1550** | {{< badge content="ogr_fdw" link="https://github.com/pramsey/pgsql-ogr-fdw" >}} | {{< ext "ogr_fdw" >}} | `1.1.7` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "file_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `ogr_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.7` | {{< bg "18" "ogr_fdw_18" "green" >}} {{< bg "17" "ogr_fdw_17" "green" >}} {{< bg "16" "ogr_fdw_16" "green" >}} {{< bg "15" "ogr_fdw_15" "green" >}} {{< bg "14" "ogr_fdw_14" "green" >}} {{< bg "13" "ogr_fdw_13" "green" >}} | `ogr_fdw_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.7` | {{< bg "18" "postgresql-18-ogr-fdw" "green" >}} {{< bg "17" "postgresql-17-ogr-fdw" "green" >}} {{< bg "16" "postgresql-16-ogr-fdw" "green" >}} {{< bg "15" "postgresql-15-ogr-fdw" "green" >}} {{< bg "14" "postgresql-14-ogr-fdw" "green" >}} {{< bg "13" "postgresql-13-ogr-fdw" "green" >}} | `postgresql-$v-ogr-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_13 : AVAIL 9" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_13 : AVAIL 9" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_13 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.1.7" "postgresql-18-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-17-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-16-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-15-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-14-ogr-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.7" "postgresql-13-ogr-fdw : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_18` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ogr_fdw_18-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ogr_fdw_18-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ogr_fdw_18-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ogr_fdw_18-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ogr_fdw_18-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ogr_fdw_18-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.2 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.9 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.5 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.3 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 92.2 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 89.8 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.5 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.7 KiB | [postgresql-18-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_17` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.3 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.8 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_17-1.1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.5-3PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.7 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.5-3PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ogr_fdw_17-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 89.9 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.3 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.4 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 88.9 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.4 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.9 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.3 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.6 KiB | [postgresql-17-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_16` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.5 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.2 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.8 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.5 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_16-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.1 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.7 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.8 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.9 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ogr_fdw_16-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.1 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.1 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.4 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 88.7 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 105.8 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.7 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.3 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.4 KiB | [postgresql-16-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_15` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.5 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.4-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_15-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.3-1.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_15-1.1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.2-2.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.9 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.4-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.9 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_15-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.4-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.8 KiB | [ogr_fdw_15-1.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.3-1.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.9 KiB | [ogr_fdw_15-1.1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.2-2.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.4-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.0 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ogr_fdw_15-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.8 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.8 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.2 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.4 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.1 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.4 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.7 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.4 KiB | [postgresql-15-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_14` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.4 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.4-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.3-1.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.2-2.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_14-1.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.2-1.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.7 KiB | [ogr_fdw_14-1.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.1-1.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.9 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.4-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_14-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.4-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_14-1.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.3-1.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.2-2.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_14-1.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.2-1.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.4-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.0 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ogr_fdw_14-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.9 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.6 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.1 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.1 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.9 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.5 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.6 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.3 KiB | [postgresql-14-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_13` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.0 KiB | [ogr_fdw_13-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.6 KiB | [ogr_fdw_13-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_13-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_13-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_13-1.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.4-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_13-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.3-1.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_13-1.1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.2-2.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_13-1.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.2-1.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.4 KiB | [ogr_fdw_13-1.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ogr_fdw_13-1.1.1-1.rhel8.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.3 KiB | [ogr_fdw_13-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ogr_fdw_13-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_13-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ogr_fdw_13-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_13-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ogr_fdw_13-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.1 KiB | [ogr_fdw_13-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ogr_fdw_13-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_13-1.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ogr_fdw_13-1.1.4-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_13-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.4 KiB | [ogr_fdw_13-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_13-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.8 KiB | [ogr_fdw_13-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_13-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_13-1.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.4-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_13-1.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.3-1.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_13-1.1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.2-2.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_13-1.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ogr_fdw_13-1.1.2-1.rhel9.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.8 KiB | [ogr_fdw_13-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ogr_fdw_13-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.8 KiB | [ogr_fdw_13-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ogr_fdw_13-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.6 KiB | [ogr_fdw_13-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ogr_fdw_13-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_13-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ogr_fdw_13-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.1 KiB | [ogr_fdw_13-1.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ogr_fdw_13-1.1.4-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_13` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.2 KiB | [ogr_fdw_13-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/ogr_fdw_13-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.2 KiB | [ogr_fdw_13-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/ogr_fdw_13-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_13` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.0 KiB | [ogr_fdw_13-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/ogr_fdw_13-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.0 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.0 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.4 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.4 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.9 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.3 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.9 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.7 KiB | [postgresql-13-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-13-ogr-fdw_1.1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-ogr-fdw" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-ogr-fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ogr_fdw;		# install via package name, for the active PG version

pig install ogr_fdw -v 18;   # install for PG 18
pig install ogr_fdw -v 17;   # install for PG 17
pig install ogr_fdw -v 16;   # install for PG 16
pig install ogr_fdw -v 15;   # install for PG 15
pig install ogr_fdw -v 14;   # install for PG 14
pig install ogr_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ogr_fdw;
```
